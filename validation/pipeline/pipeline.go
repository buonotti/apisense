package pipeline

import (
	"fmt"
	"sync"
	"time"

	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/util"
	"github.com/buonotti/apisense/validation"
	"github.com/buonotti/apisense/validation/definitions"
	"github.com/buonotti/apisense/validation/validators"
	"github.com/speps/go-hashids/v2"
)

// Pipeline represents the validation pipeline
type Pipeline struct {
	Validators []validators.Validator `json:"validators"` // Validators are the validators that will be applied to the items in the pipeline
}

// ValidatedEndpoint is the collection of results for each endpoint (definition)
// that are generated for each different call to the endpoint (produced by the
// multiple variable values)
type ValidatedEndpoint struct {
	EndpointName    string           `json:"endpointName"`    // EndpointName is he name of the endpoint
	TestCaseResults []TestCaseResult `json:"testCaseResults"` // TestCaseResults are the collection of TestCaseResult that describe the result of validating a single api call
}

// TestCaseResult is the result of validating a single api call
type TestCaseResult struct {
	Url              string            `json:"url"`              // Url is the url of the api call (with query parameters)
	ValidatorResults []ValidatorResult `json:"validatorResults"` // ValidatorResults is the collection of ValidatorResult that describe the result of each validator
}

// ValidatorResult is the output of a single validator
type ValidatorResult struct {
	Name    string                     `json:"name"`    // Name is the name of the validator
	Status  validators.ValidatorStatus `json:"status"`  // Status is the status of the validator (success/fail/skipped)
	Message string                     `json:"message"` // Message is the error message of the validator
}

// NewPipelineWithValidators creates a new validation pipeline with the given validators already added
func NewPipelineWithValidators(validators ...validators.Validator) (Pipeline, error) {
	pipeline, err := NewPipeline()
	if err != nil {
		return Pipeline{}, err
	}

	pipeline.Validators = validators
	return pipeline, nil
}

// NewPipeline creates a new validation pipeline without any validators
func NewPipeline() (Pipeline, error) {
	pipeline := Pipeline{
		Validators: make([]validators.Validator, 0),
	}

	return pipeline, nil
}

// AddValidator adds a validator to the end of the validation pipeline
func (p *Pipeline) AddValidator(validator validators.Validator) {
	p.Validators = append(p.Validators, validator)
}

// RemoveValidator removes a validator from the validation pipeline identified by its name
func (p *Pipeline) RemoveValidator(name string) {
	for i, v := range p.Validators {
		if v.Name() == name {
			p.Validators = append(p.Validators[:i], p.Validators[i+1:]...)
		}
	}
}

// Validate validates all the test cases in the pipeline and returns a Report
func (p *Pipeline) Validate() (Report, error) {
	allDefinitions, err := definitions.Endpoints()
	if err != nil {
		return Report{}, err
	}

	enabledDefinitions := util.Where(allDefinitions, func(d definitions.Endpoint) bool { return d.IsEnabled })
	results := make([]ValidatedEndpoint, len(enabledDefinitions))
	if len(enabledDefinitions) != 0 {
		testCases := make([]validation.EndpointTestCases, len(enabledDefinitions))
		for i, definition := range enabledDefinitions {
			testCase, err := validation.Preprocess(definition)
			if err != nil {
				return Report{}, err
			}
			testCases[i] = testCase
		}

		var wg sync.WaitGroup

		for i, endpointTestCase := range testCases {
			wg.Add(1)
			go func(results *[]ValidatedEndpoint, slot int, testCases validation.EndpointTestCases) {
				validatorResults := p.testCaseResults(testCases)
				(*results)[slot] = ValidatedEndpoint{
					EndpointName:    testCases.Definition.Name,
					TestCaseResults: validatorResults,
				}
				wg.Done()
			}(&results, i, endpointTestCase)
		}

		wg.Wait()
	}

	currentTime := time.Now().UTC()
	hashIDData := hashids.NewData()
	hashIDData.Salt = "apisense"
	hashIDData.MinLength = 5
	h, _ := hashids.NewWithData(hashIDData)
	id, _ := h.Encode([]int{int(currentTime.Unix())})

	// return the report with the current timestamp
	return Report{
		Id:        id,
		Time:      util.ApisenseTime(currentTime),
		Endpoints: results,
	}, nil
}

// testCaseResults validates a collection of items and returns the results
func (p *Pipeline) testCaseResults(testCases validation.EndpointTestCases) []TestCaseResult {
	testCaseResults := make([]TestCaseResult, len(testCases.HttpRequests))

	var wg sync.WaitGroup

	for i, request := range testCases.HttpRequests {
		wg.Add(1)
		go func(results *[]TestCaseResult, slot int, request validation.EndpointRequest, endpoint definitions.Endpoint) {
			response, err := validation.Fetch(request, endpoint)
			if err != nil {
				testCaseResults[i] = TestCaseResult{
					Url: response.Url,
					ValidatorResults: []ValidatorResult{
						{
							Name:    "fetcher",
							Status:  validators.ValidatorStatusFail,
							Message: err.Error(),
						},
					},
				}
			} else {
				testCaseResults[i] = TestCaseResult{
					Url:              response.Url,
					ValidatorResults: p.validateTestCase(response, endpoint),
				}
			}

			wg.Done()
		}(&testCaseResults, i, request, testCases.Definition)
	}

	wg.Wait()

	return testCaseResults
}

// validateTestCase validates a single item and returns the result of the validators
func (p *Pipeline) validateTestCase(response validation.EndpointResponse, definition definitions.Endpoint) []ValidatorResult {
	validatorResults := make([]ValidatorResult, 1)

	okCode := 200
	if definition.OkCode > 0 {
		okCode = definition.OkCode
	}

	if response.StatusCode != okCode {
		validatorResults[0] = ValidatorResult{
			Name:    "status",
			Status:  validators.ValidatorStatusFail,
			Message: fmt.Sprintf("status validation failed. Expected %d, got %d", definition.OkCode, response.StatusCode),
		}
		return validatorResults
	}

	validatorResults[0] = ValidatorResult{
		Name:    "status",
		Status:  validators.ValidatorStatusSuccess,
		Message: "",
	}

	for _, validator := range p.Validators {
		validatorResult := ValidatorResult{
			Name:    validator.Name(),
			Status:  validators.ValidatorStatusSuccess,
			Message: "",
		}

		if util.Contains(definition.ExcludedValidators, validator.Name()) {
			log.DaemonLogger().Warn("Validator is excluded", "endpoint", definition.Name, "validator", validator.Name())
			validatorResult.Status = validators.ValidatorStatusSkipped
			validatorResults = append(validatorResults, validatorResult)
			continue
		}

		err := validator.Validate(validators.ValidationItem{
			Response:   response,
			Definition: definition,
		})
		if err != nil {
			validatorResult.Message = err.Error()
			validatorResult.Status = validators.ValidatorStatusFail
			validatorResults = append(validatorResults, validatorResult)
			if validator.IsFatal() {
				break
			} else {
				continue
			}
		}

		validatorResults = append(validatorResults, validatorResult)
	}

	return validatorResults
}

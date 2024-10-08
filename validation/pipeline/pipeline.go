package pipeline

import (
	"sync"
	"time"

	"github.com/speps/go-hashids/v2"

	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/util"
	"github.com/buonotti/apisense/validation/definitions"
	"github.com/buonotti/apisense/validation/fetcher"
	"github.com/buonotti/apisense/validation/response"
	"github.com/buonotti/apisense/validation/validators"
)

// TestCase represents an item in the validation pipeline
type TestCase struct {
	SchemaEntries      []response.SchemaEntry `json:"schemaEntries"`      // SchemaEntries are the schema definitions of every field in the items Data
	Data               map[string]any         `json:"data"`               // Data is the raw response data mapped in a map
	Url                string                 `json:"url"`                // Url is the request url of the item
	EndpointName       string                 `json:"endpointName"`       // EndpointName is the name of the endpoint in the definition file
	Code               int                    `json:"code"`               // Code is the response code of the request
	ExcludedValidators []string               `json:"excludedValidators"` // ExcludedValidators is a list of validators that should be excluded from the validation
}

// Pipeline represents the validation pipeline
type Pipeline struct {
	TestCases  map[string][]fetcher.TestCase `json:"testCases"`  // TestCases are the collection of TestCase for each endpoint (definition file)
	Validators []validators.Validator        `json:"validators"` // Validators are the validators that will be applied to the items in the pipeline
	fetcher    fetcher.Fetcher               // fetcher is the fetcher that will be used to fetch the items in the pipeline
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
	allDefinitions, err := definitions.Endpoints()
	if err != nil {
		return Pipeline{}, err
	}

	pipeline := Pipeline{
		TestCases:  make(map[string][]fetcher.TestCase),
		Validators: make([]validators.Validator, 0),
		fetcher:    fetcher.New(),
	}

	enabledDefinitions := util.Where(allDefinitions, func(d definitions.Endpoint) bool { return d.IsEnabled })

	for _, definition := range enabledDefinitions {
		items, err := pipeline.fetcher.Fetch(definition)
		if err != nil {
			return Pipeline{}, err
		}

		pipeline.TestCases[definition.Name] = items
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

// Reload re-populates the Pipeline.TestCases collection
func (p *Pipeline) Reload() error {
	log.DaemonLogger().Info("Reloading pipeline")
	defs, err := definitions.Endpoints()
	if err != nil {
		return err
	}

	if len(defs) == 0 {
		log.DaemonLogger().Warn("No endpoint definitions found.")
		return nil
	}

	for _, definition := range defs {
		items, err := p.fetcher.Fetch(definition)
		if err != nil {
			return err
		}
		p.TestCases[definition.Name] = items
	}

	return nil
}

// Validate validates all the test cases in the pipeline and returns a Report
func (p *Pipeline) Validate() Report {
	results := make([]ValidatedEndpoint, len(p.TestCases))

	var wg sync.WaitGroup

	i := 0
	for endpoint, testCases := range p.TestCases {
		wg.Add(1)
		go func(results *[]ValidatedEndpoint, slot int, testCases []fetcher.TestCase, endpoint string) {
			validatorResults := p.testCaseResults(testCases)
			(*results)[slot] = ValidatedEndpoint{
				EndpointName:    endpoint,
				TestCaseResults: validatorResults,
			}
			wg.Done()
		}(&results, i, testCases, endpoint)

		i += 1
	}

	wg.Wait()

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
	}
}

// testCaseResults validates a collection of items and returns the results
func (p *Pipeline) testCaseResults(items []fetcher.TestCase) []TestCaseResult {
	testCaseResults := make([]TestCaseResult, 0)

	// testCaseResults each single item and append to the results
	for _, item := range items {
		validatorResults := p.validateTestCase(item)
		testCaseResults = append(testCaseResults, TestCaseResult{
			Url:              item.Url,
			ValidatorResults: validatorResults,
		})
	}

	return testCaseResults
}

// validateTestCase validates a single item and returns the result of the validators
func (p *Pipeline) validateTestCase(item fetcher.TestCase) []ValidatorResult {
	validatorResults := make([]ValidatorResult, 0)

	for _, validator := range p.Validators {
		validatorResult := ValidatorResult{
			Name:    validator.Name(),
			Status:  validators.ValidatorStatusSuccess,
			Message: "",
		}

		if util.Contains(item.ExcludedValidators, validator.Name()) {
			log.DaemonLogger().Warn("Validator is excluded", "endpoint", item.EndpointName, "validator", validator.Name())
			validatorResult.Status = validators.ValidatorStatusSkipped
			validatorResults = append(validatorResults, validatorResult)
			continue
		}

		err := validator.Validate(item)
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

package validation

import (
	"time"

	"github.com/speps/go-hashids/v2"

	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/util"
	"github.com/buonotti/apisense/validation/response"
)

// PipelineTestCase represents an item in the validation pipeline
type PipelineTestCase struct {
	SchemaEntries      []response.SchemaEntry `json:"schemaEntries"`      // SchemaEntries are the schema definitions of every field in the items Data
	Data               map[string]any         `json:"data"`               // Data is the raw response data mapped in a map
	Url                string                 `json:"url"`                // Url is the request url of the item
	EndpointName       string                 `json:"endpointName"`       // EndpointName is the name of the endpoint in the definition file
	Code               int                    `json:"code"`               // Code is the response code of the request
	ExcludedValidators []string               `json:"excludedValidators"` // ExcludedValidators is a list of validators that should be excluded from the validation
}

// Pipeline represents the validation pipeline
type Pipeline struct {
	TestCases  map[string][]PipelineTestCase `json:"testCases"`  // TestCases are the collection of PipelineTestCase for each endpoint (definition file)
	Validators []Validator                   `json:"validators"` // Validators are the validators that will be applied to the items in the pipeline
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
	Name    string `json:"name"`    // Name is the name of the validator
	Status  string `json:"status"`  // Status is the status of the validator (success/fail/skipped)
	Message string `json:"message"` // Message is the error message of the validator
}

// NewPipelineWithValidators creates a new validation pipeline with the given validators already added
func NewPipelineWithValidators(validators ...Validator) (Pipeline, error) {
	pipeline, err := NewPipeline()
	if err != nil {
		return Pipeline{}, err
	}

	pipeline.Validators = validators
	return pipeline, nil
}

// NewPipeline creates a new validation pipeline without any validators
func NewPipeline() (Pipeline, error) {
	definitions, err := EndpointDefinitions()
	if err != nil {
		return Pipeline{}, err
	}

	pipeline := Pipeline{
		TestCases:  make(map[string][]PipelineTestCase),
		Validators: make([]Validator, 0),
	}

	enabledDefinitions := util.Where(definitions, func(d EndpointDefinition) bool { return d.IsEnabled })

	for _, definition := range enabledDefinitions {
		items, err := loadTestCases(definition)
		if err != nil {
			return Pipeline{}, err
		}

		pipeline.TestCases[definition.Name] = items
	}

	return pipeline, nil
}

// AddValidator adds a validator to the end of the validation pipeline
func (p *Pipeline) AddValidator(validator Validator) {
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
	definitions, err := EndpointDefinitions()
	if err != nil {
		return err
	}

	if len(definitions) == 0 {
		log.DaemonLogger.Warnf("No endpoint definitions found.")
		return nil
	}

	for _, definition := range definitions {
		items, err := loadTestCases(definition)
		if err != nil {
			return err
		}
		p.TestCases[definition.Name] = items
	}

	return nil
}

// Validate validates all the test cases in the pipeline and returns a Report
func (p *Pipeline) Validate() Report {
	results := make([]ValidatedEndpoint, 0)

	// for each endpoint testCaseResults all the testCases
	for endpoint, testCases := range p.TestCases {
		validatorResults := p.testCaseResults(testCases)
		results = append(results, ValidatedEndpoint{
			EndpointName:    endpoint,
			TestCaseResults: validatorResults,
		})
	}

	currentTime := time.Now()
	hashIDData := hashids.NewData()
	hashIDData.Salt = "apisense"
	hashIDData.MinLength = 5
	h, _ := hashids.NewWithData(hashIDData)
	id, _ := h.Encode([]int{int(currentTime.Unix())})

	// return the report with the current timestamp
	return Report{
		Id:        id,
		Time:      ReportTime(currentTime),
		Endpoints: results,
	}
}

// testCaseResults validates a collection of items and returns the results
func (p *Pipeline) testCaseResults(items []PipelineTestCase) []TestCaseResult {
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
func (p *Pipeline) validateTestCase(item PipelineTestCase) []ValidatorResult {
	validatorResults := make([]ValidatorResult, 0)

	for _, validator := range p.Validators {
		validatorResult := ValidatorResult{
			Name:    validator.Name(),
			Status:  "success",
			Message: "",
		}

		if util.Contains(item.ExcludedValidators, validator.Name()) {
			log.DaemonLogger.Warnf("validator %s is excluded for %s", validator.Name(), item.Url)
			validatorResult.Status = "skipped"
			validatorResults = append(validatorResults, validatorResult)
			continue
		}

		err := validator.Validate(item)

		if err != nil {
			validatorResult.Message = err.Error()
			validatorResult.Status = "fail"
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

// loadTestCases parses the definition files and populates the Pipeline.TestCases collection
func loadTestCases(definition EndpointDefinition) ([]PipelineTestCase, error) {
	log.DaemonLogger.Infof("loading pipeline test-cases for %s", definition.Name)
	var testCases []PipelineTestCase
	requests, err := parseRequests(definition)
	if err != nil {
		return nil, err
	}

	responses, err := send(requests)
	if err != nil {
		return nil, err
	}

	for _, resp := range responses {
		testCases = append(testCases, PipelineTestCase{
			SchemaEntries:      definition.ResultSchema,
			Data:               resp.RawData,
			Url:                resp.Url,
			Code:               resp.StatusCode,
			ExcludedValidators: definition.ExcludedValidators,
			EndpointName:       definition.Name,
		})
	}

	return testCases, nil
}

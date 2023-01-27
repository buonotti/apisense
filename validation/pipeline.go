package validation

import (
	"time"

	"github.com/speps/go-hashids/v2"

	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/util"
)

// Validator is an interface that all validators in the pipeline must implement
type Validator interface {
	Name() string                     // Name returns the name of the validator
	Validate(item PipelineItem) error // Validate validates the given item and return nil on success or an error on failure
	Fatal() bool                      // Fatal returns true if the validator is fatal and the pipeline should stop on failure
}

// PipelineItem represents an item in the validation pipeline
type PipelineItem struct {
	SchemaEntries      []SchemaEntry  `json:"schemaEntries"`      // SchemaEntries are the schema definitions of every field in the items Data
	Data               map[string]any `json:"data"`               // Data is the raw response data mapped in a map
	Url                string         `json:"url"`                // Url is the request url of the item
	Code               int            `json:"code"`               // Code is the response code of the request
	ExcludedValidators []string       `json:"excludedValidators"` // ExcludedValidators is a list of validators that should be excluded from the validation
}

// Pipeline represents the validation pipeline
type Pipeline struct {
	EndpointItems map[string][]PipelineItem `json:"endpointItems"` // EndpointItems are the collection of PipelineItem for each endpoint (definition file)
	Validators    []Validator               `json:"validators"`    // Validators are the validators that will be applied to the items in the pipeline
}

// ValidatedEndpoint is the collection of results for each endpoint (definition)
// that are generated for each different call to the endpoint (produced by the
// multiple variable values)
type ValidatedEndpoint struct {
	EndpointName string   `json:"endpointName"` // EndpointName is he name of the endpoint
	Results      []Result `json:"results"`      // Results are the collection of Result that describe the result of validating a single api call
}

// Result is the result of validating a single api call
type Result struct {
	Url              string            `json:"url"`              // Url is the url of the api call (with query parameters)
	ValidatorsOutput []ValidatorOutput `json:"validatorsOutput"` // ValidatorsOutput is the collection of ValidatorOutput that describe the result of each validator
}

// ValidatorOutput is the output of a single validator
type ValidatorOutput struct {
	Validator string `json:"validator"` // Validator is the name of the validator
	Status    string `json:"status"`    // Status is the status of the validator (success/fail/skipped)
	Error     string `json:"error"`     // Error is the error message of the validator
}

// NewPipelineV creates a new validation pipeline with the given validators already added
func NewPipelineV(validators ...Validator) (Pipeline, error) {
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
		EndpointItems: make(map[string][]PipelineItem),
		Validators:    make([]Validator, 0),
	}

	for _, definition := range definitions {
		items, err := loadItems(definition)
		if err != nil {
			return Pipeline{}, err
		}

		pipeline.EndpointItems[definition.Name] = items
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

// RefreshItems re-populates the Pipeline.EndpointItems collection
func (p *Pipeline) RefreshItems() error {
	definitions, err := EndpointDefinitions()
	if err != nil {
		return err
	}

	for _, definition := range definitions {
		items, err := loadItems(definition)
		if err != nil {
			return err
		}
		p.EndpointItems[definition.Name] = items
	}

	return nil
}

// Validate validates all the items in the pipeline and returns a Report
func (p *Pipeline) Validate() Report {
	results := make([]ValidatedEndpoint, 0)

	// for each endpoint validate all the items
	for endpoint, items := range p.EndpointItems {
		validatorResults := p.validateItems(items)
		results = append(results, ValidatedEndpoint{
			EndpointName: endpoint,
			Results:      validatorResults,
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

// validateItems validates a collection of items and returns the results
func (p *Pipeline) validateItems(items []PipelineItem) []Result {
	validatorResults := make([]Result, 0)

	// validate each single item and append to the results
	for _, item := range items {
		validatorOutputs := p.validateSingleItem(item)
		validatorResults = append(validatorResults, Result{
			Url:              item.Url,
			ValidatorsOutput: validatorOutputs,
		})
	}

	return validatorResults
}

// validateSingleItem validates a single item and returns the result of the validators
func (p *Pipeline) validateSingleItem(item PipelineItem) []ValidatorOutput {
	validatorOutputs := make([]ValidatorOutput, 0)

	for _, validator := range p.Validators {
		validatorOutput := ValidatorOutput{
			Validator: validator.Name(),
			Status:    "success",
			Error:     "",
		}

		if util.Contains(item.ExcludedValidators, validator.Name()) {
			log.DaemonLogger.Warnf("validator %s is excluded for %s", validator.Name(), item.Url)
			validatorOutput.Status = "skipped"
			validatorOutputs = append(validatorOutputs, validatorOutput)
			continue
		}

		err := validator.Validate(item)

		if err != nil {
			validatorOutput.Error = err.Error()
			validatorOutput.Status = "fail"
			validatorOutputs = append(validatorOutputs, validatorOutput)
			if validator.Fatal() {
				break
			}
		}

		validatorOutputs = append(validatorOutputs, validatorOutput)
	}

	return validatorOutputs
}

// loadItems parses the definition files and populates the Pipeline.EndpointItems collection
func loadItems(definition EndpointDefinition) ([]PipelineItem, error) {
	log.DaemonLogger.Infof("loading pipeline items for %s", definition.Name)
	var items []PipelineItem
	requests, err := parseRequests(definition)
	if err != nil {
		return nil, err
	}

	responses, err := send(requests)
	if err != nil {
		return nil, err
	}

	for _, response := range responses {
		items = append(items, PipelineItem{
			SchemaEntries:      definition.ResultSchema.Entries,
			Data:               response.RawData,
			Url:                response.Url,
			Code:               response.StatusCode,
			ExcludedValidators: definition.ExcludedValidators,
		})
	}

	return items, nil
}

package validation

import (
	"time"
)

type Validator interface {
	Name() string
	Validate(item PipelineItem) error
}

type PipelineItem struct {
	resultSchema
	Data     map[string]any
	Endpoint string
	ParamMap map[string][]string
	Code     int
}

type Pipeline struct {
	EndpointItems map[string][]PipelineItem
	Validators    []Validator
}

type ValidatedEndpoint struct {
	EndpointName string
	Results      []Result
}

type Result struct {
	Url              string
	ValidatorsOutput []ValidatorOutput
}

type ValidatorOutput struct {
	Validator string
	Error     string
}

func loadItems(definition endpointDefinition) ([]PipelineItem, error) {
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
			resultSchema: definition.ResultSchema,
			Data:         response.RawData,
			Endpoint:     response.Url,
			Code:         response.StatusCode,
		})
	}
	return items, nil
}

func NewPipelineV(validators ...Validator) (Pipeline, error) {
	pipeline, err := NewPipeline()
	if err != nil {
		return Pipeline{}, err
	}
	pipeline.Validators = validators
	return pipeline, nil
}

func NewPipeline() (Pipeline, error) {
	definitions, err := endpointDefinitions()
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

func (p Pipeline) Validate() Report {
	results := make([]ValidatedEndpoint, 0)
	for endpoint, items := range p.EndpointItems {
		validatorResults := make([]Result, 0)
		for _, item := range items {
			validatorOutputs := make([]ValidatorOutput, 0)
			for _, validator := range p.Validators {
				err := validator.Validate(item)
				if err != nil {
					validatorOutputs = append(validatorOutputs, ValidatorOutput{
						Validator: validator.Name(),
						Error:     err.Error(),
					})
					break
				} else {
					validatorOutputs = append(validatorOutputs, ValidatorOutput{
						Validator: validator.Name(),
						Error:     "",
					})
				}
			}
			validatorResults = append(validatorResults, Result{
				Url:              item.Endpoint,
				ValidatorsOutput: validatorOutputs,
			})
		}
		results = append(results, ValidatedEndpoint{
			EndpointName: endpoint,
			Results:      validatorResults,
		})
	}
	return Report{
		Time:    time.Now(),
		Results: results,
	}
}

func (p Pipeline) AddValidator(validator Validator) {
	p.Validators = append(p.Validators, validator)
}

func (p Pipeline) RemoveValidator(name string) {
	for i, v := range p.Validators {
		if v.Name() == name {
			p.Validators = append(p.Validators[:i], p.Validators[i+1:]...)
		}
	}
}

func (p Pipeline) RefreshItems() error {
	definitions, err := endpointDefinitions()
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

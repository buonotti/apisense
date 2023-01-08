package validation

type Validator interface {
	Name() string
	Validate(item Item) error
}

type Item struct {
	ResultSchema
	Data     map[string]any
	Endpoint string
	ParamMap map[string][]string
	Code     int
}

type Pipeline struct {
	EndpointItems map[string][]Item
	Validators    []Validator
}

type Result struct {
	EndpointName     string
	ValidatorResults []ValidatorResult
}

type ValidatorResult struct {
	Validator string
	Url       string
	Error     error
}

func loadItems(definition EndpointDefinition) ([]Item, error) {
	var items []Item
	requests, err := parseRequests(definition)
	if err != nil {
		return nil, err
	}
	responses, err := send(requests)
	if err != nil {
		return nil, err
	}
	for _, response := range responses {
		items = append(items, Item{
			ResultSchema: definition.ResultSchema,
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
		EndpointItems: make(map[string][]Item),
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

func (p Pipeline) Validate() []Result {
	var results []Result
	for endpoint, items := range p.EndpointItems {
		var errors []ValidatorResult
		for _, item := range items {
			for _, validator := range p.Validators {
				err := validator.Validate(item)
				errors = append(errors, ValidatorResult{
					Validator: validator.Name(),
					Url:       item.Endpoint,
					Error:     err,
				})
			}
		}
		results = append(results, Result{
			EndpointName:     endpoint,
			ValidatorResults: errors,
		})
	}
	return results
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

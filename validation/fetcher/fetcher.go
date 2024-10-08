package fetcher

import (
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/validation/definitions"
)

type TestCase struct {
	SchemaEntries      []definitions.SchemaEntry `json:"schemaEntries"`      // SchemaEntries are the schema definitions of every field in the items Data
	Data               map[string]any            `json:"data"`               // Data is the raw response data mapped in a map
	Url                string                    `json:"url"`                // Url is the request url of the item
	EndpointName       string                    `json:"endpointName"`       // EndpointName is the name of the endpoint in the definition file
	Code               int                       `json:"code"`               // Code is the response code of the request
	ExcludedValidators []string                  `json:"excludedValidators"` // ExcludedValidators is a list of validators that should be excluded from the validation
}

type Fetcher interface {
	Fetch(definition definitions.Endpoint) ([]TestCase, error)
}

type defaultFetcher struct{}

func (f *defaultFetcher) Fetch(definition definitions.Endpoint) ([]TestCase, error) {
	log.DaemonLogger().Info("Fetching data", "endpoint", definition.Name)
	var testCases []TestCase
	requests, err := parseRequests(definition)
	if err != nil {
		return nil, err
	}

	responses, err := send(requests)
	if err != nil {
		return nil, err
	}

	for _, resp := range responses {
		testCases = append(testCases, TestCase{
			SchemaEntries:      definition.ResponseSchema,
			Data:               resp.RawData,
			Url:                resp.Url,
			Code:               resp.StatusCode,
			ExcludedValidators: definition.ExcludedValidators,
			EndpointName:       definition.Name,
		})
	}

	return testCases, nil
}

func New() Fetcher {
	return &defaultFetcher{}
}

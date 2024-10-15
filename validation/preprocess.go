package validation

import (
	"strings"
	"text/template"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/util"
	"github.com/buonotti/apisense/validation/definitions"
)

type EndpointRequest struct {
	Url             string
	QueryParameters map[string]string
	Payload         string
}

type EndpointTestCases struct {
	Definition   definitions.Endpoint
	HttpRequests []EndpointRequest
}

func Preprocess(definition definitions.Endpoint) (EndpointTestCases, error) {
	var requests []EndpointRequest

	// normalize the variables
	vars, err := normalizeVariables(definition)
	if err != nil {
		return EndpointTestCases{}, err
	}

	// get the length of the first non-constant element of the variables. if there are only constants the valueCount is 1
	firstVariableVar := util.FindFirst(definition.Variables, func(param definitions.Variable) bool { return !param.IsConstant })
	valueCount := 1
	if firstVariableVar != nil {
		valueCount = len(firstVariableVar.Values)
	}

	// go through all the values of the variables and create a request for each step
	for valueCycle := 0; valueCycle < valueCount; valueCycle++ {
		// create a variable map where the key is the name of the variable and the value
		// is retrieved from the variables.EndpointParameter where the index is the
		// current cycle
		variableMap := definitions.VariableMap(make(map[string]any))
		for varIndex, variable := range vars {
			variableMap[definition.Variables[varIndex].Name] = variable.Value(valueCycle)
		}

		// create a new golang template
		temp := template.New(definition.Name)

		// parse the template on the base url of the definition
		temp, err := temp.Parse(definition.BaseUrl)
		if err != nil {
			return EndpointTestCases{}, errors.CannotApplyTemplateError.Wrap(err, "cannot parse endpoint: "+definition.BaseUrl)
		}

		// execute the template using the variable map of the current cycle
		var stringBuilder strings.Builder
		err = temp.Execute(&stringBuilder, variableMap)
		if err != nil {
			return EndpointTestCases{}, errors.CannotApplyTemplateError.Wrap(err, "cannot apply template to endpoint: "+definition.BaseUrl)
		}

		// get the url with the interpolated variables
		parsedUrl := stringBuilder.String()

		// interpolate the variable map with each query parameter
		queryParams := make(map[string]string)
		for _, queryParam := range definition.QueryParameters {
			temp, err := temp.Parse(queryParam.Value)
			if err != nil {
				return EndpointTestCases{}, errors.CannotApplyTemplateError.Wrap(err, "cannot parse query parameter: "+queryParam.Value)
			}

			var stringBuilder strings.Builder
			err = temp.Execute(&stringBuilder, variableMap)
			if err != nil {
				return EndpointTestCases{}, errors.CannotApplyTemplateError.Wrap(err, "cannot apply template to query parameter: "+queryParam.Value)
			}

			queryParam.Value = stringBuilder.String()
			queryParams[queryParam.Name] = queryParam.Value
		}

		// create a new request and add it to the list of requests
		requests = append(requests, EndpointRequest{
			Url:             parsedUrl,
			QueryParameters: queryParams,
		})
	}
	return EndpointTestCases{
		Definition:   definition,
		HttpRequests: requests,
	}, nil
}

// normalizeVariables converts all given variables in the definition of an endpoint to a collection of variables.EndpointParameter
func normalizeVariables(definition definitions.Endpoint) ([]definitions.Variable, error) {
	// create a new variables.EndpointParameter for each variable. a constant
	// parameter is created for constants and a variable parameter for variable
	// values
	var params []definitions.Variable
	for _, param := range definition.Variables {
		params = append(params, definitions.Variable{
			Name:       param.Name,
			IsConstant: param.IsConstant,
			Values:     param.Values,
		})
	}

	return params, nil
}

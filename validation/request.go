package validation

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"text/template"

	"github.com/go-resty/resty/v2"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/util"
	"github.com/buonotti/apisense/validation/variables"
)

// rest return the rest client with the set headers
func rest() *resty.Client {
	return resty.
		New().
		SetHeader("Accept", "*/*").
		SetDisableWarn(true)
}

// endpointRequest is a single request to an endpoint
type endpointRequest struct {
	Url             string            // Url is the base url of the request
	QueryParameters map[string]string // QueryParameters are the query parameters of the request that will be encoded
	ResponseFormat  string            // ResponseFormat is the format of the response (json or xml)
}

// endpointResponse is a single response
type endpointResponse struct {
	StatusCode          int                 // StatusCode is the status code of the response
	RawData             map[string]any      // RawData is the raw data of the response mapped to a map
	Url                 string              // Url is the full url of the request
	UsedQueryParameters map[string][]string // UsedQueryParameters are the query parameters that were used in the request
}

// requestData sends based on the given request a http GET request and returns the response
func requestData(definition endpointRequest) (endpointResponse, error) {
	var data map[string]any

	resp, err := rest().R().SetQueryParams(definition.QueryParameters).Get(definition.Url)
	if err != nil {
		return endpointResponse{
			RawData:    nil,
			StatusCode: -1,
			Url:        "",
		}, errors.CannotRequestDataError.Wrap(err, "cannot request data from "+definition.Url)
	}

	loc := resp.RawResponse.Request.URL.String()
	log.DaemonLogger.Infof("sent request to %s", resp.Request.URL)

	if resp.StatusCode() != 200 {
		return endpointResponse{
			RawData:    nil,
			StatusCode: resp.StatusCode(),
			Url:        loc,
		}, errors.NewF(errors.CannotRequestDataError, "cannot send request to: %s (status code %d)", definition.Url, resp.StatusCode())
	}

	switch definition.ResponseFormat {
	case "json":
		err = json.Unmarshal(resp.Body(), &data)
		if err != nil {
			return endpointResponse{
				RawData:    nil,
				StatusCode: -1,
				Url:        loc,
			}, errors.CannotParseDataError.Wrap(err, "cannot parse data from: "+definition.Url)
		}
	case "xml":
		err = xml.Unmarshal(resp.Body(), &data)
		if err != nil {
			return endpointResponse{
				RawData:    nil,
				StatusCode: -1,
				Url:        loc,
			}, errors.CannotParseDataError.Wrap(err, "cannot parse data from: "+definition.Url)
		}
	default:
		return endpointResponse{
			RawData:    nil,
			StatusCode: -1,
			Url:        loc,
		}, errors.UnknownFormatError.Wrap(err, "unknown format: "+definition.ResponseFormat)
	}
	return endpointResponse{
		RawData:    data,
		StatusCode: resp.StatusCode(),
		Url:        loc,
	}, nil
}

// normalizeVariables converts all given variables in the definition of an endpoint to a collection of variables.EndpointParameter
func normalizeVariables(definition EndpointDefinition) ([]variables.EndpointParameter, error) {
	// get the length of the first non-constant element of the variables. if there are only constants the valueCount is 1
	firstVariableVar := util.FindFirst(definition.Variables, func(param variableSchema) bool { return !param.IsConstant })
	valueCount := 1
	if firstVariableVar != nil {
		valueCount = len(firstVariableVar.Values)
	}

	// check if any of the non-constant variables has a different length of values than the first non-constant variable
	for _, param := range definition.Variables {
		if !param.IsConstant && len(param.Values) != valueCount {
			return nil, errors.NewF(errors.VariableValueLengthMismatchError, "variable %s has %d values, but %d are expected", param.Name, len(param.Values), valueCount)
		}
	}

	// create a new variables.EndpointParameter for each variable. a constant
	// parameter is created for constants and a variable parameter for variable
	// values
	var params []variables.EndpointParameter
	for _, param := range definition.Variables {
		if param.IsConstant {
			params = append(params, variables.NewConstantEndpointParameter(param.Values[0]))
		} else {
			params = append(params, variables.NewVariableEndpointParameter(param.Values))
		}
	}
	return params, nil
}

// parseRequests reads in the definition of an endpoint and returns a list of requests that need to be sent
func parseRequests(definition EndpointDefinition) ([]endpointRequest, error) {
	var requests []endpointRequest

	// normalize the variables
	vars, err := normalizeVariables(definition)
	if err != nil {
		return nil, err
	}

	// get the length of the first non-constant element of the variables. if there are only constants the valueCount is 1
	firstVariableVar := util.FindFirst(definition.Variables, func(param variableSchema) bool { return !param.IsConstant })
	valueCount := 1
	if firstVariableVar != nil {
		valueCount = len(firstVariableVar.Values)
	}

	// go through all the values of the variables and create a request for each step
	for valueCycle := 0; valueCycle < valueCount; valueCycle++ {
		// create a variable map where the key is the name of the variable and the value
		// is retrieved from the variables.EndpointParameter where the index is the
		// current cycle
		variableMap := variables.VariableMap(make(map[string]any))
		for varIndex, variable := range vars {
			variableMap[definition.Variables[varIndex].Name] = variable.Value(valueCycle)
		}

		// create a new golang template
		temp := template.New(definition.Name)

		// parse the template on the base url of the definition
		temp, err := temp.Parse(definition.BaseUrl)
		if err != nil {
			return nil, errors.CannotApplyTemplateError.Wrap(err, "cannot parse endpoint: "+definition.BaseUrl)
		}

		// execute the template using the variable map of the current cycle
		var stringBuilder strings.Builder
		err = temp.Execute(&stringBuilder, variableMap)
		if err != nil {
			return nil, errors.CannotApplyTemplateError.Wrap(err, "cannot apply template to endpoint: "+definition.BaseUrl)
		}

		// get the url with the interpolated variables
		parsedUrl := stringBuilder.String()

		// interpolate the variable map with each query parameter
		var queryParams = make(map[string]string)
		for _, queryParam := range definition.QueryParameters {
			temp, err := temp.Parse(queryParam.Value)
			if err != nil {
				return nil, errors.CannotApplyTemplateError.Wrap(err, "cannot parse query parameter: "+queryParam.Value)
			}

			var stringBuilder strings.Builder
			err = temp.Execute(&stringBuilder, variableMap)
			if err != nil {
				return nil, errors.CannotApplyTemplateError.Wrap(err, "cannot apply template to query parameter: "+queryParam.Value)
			}

			queryParam.Value = stringBuilder.String()
			queryParams[queryParam.Name] = queryParam.Value
		}

		// create a new request and add it to the list of requests
		requests = append(requests, endpointRequest{
			Url:             parsedUrl,
			QueryParameters: queryParams,
			ResponseFormat:  "json",
		})
	}
	return requests, nil
}

// send sends a collection of endpointRequest and returns the responses
func send(requests []endpointRequest) ([]endpointResponse, error) {
	var responses []endpointResponse
	for _, request := range requests {
		resp, err := requestData(request)
		if err != nil {
			return nil, err
		}

		responses = append(responses, resp)
	}
	return responses, nil
}

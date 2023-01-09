package validation

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
	"text/template"

	"github.com/go-resty/resty/v2"

	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/log"
	"github.com/buonotti/odh-data-monitor/util"
	"github.com/buonotti/odh-data-monitor/validation/variables"
)

func rest() *resty.Client {
	return resty.New().SetHeader("Accept", "*/*")
}

type endpointRequest struct {
	Url             string
	QueryParameters map[string]string
	ResponseFormat  string
}

type endpointResponse struct {
	StatusCode          int
	RawData             map[string]any
	Url                 string
	UsedQueryParameters map[string][]string
}

func requestData(definition endpointRequest) (endpointResponse, error) {
	var data map[string]any
	resp, err := rest().R().SetQueryParams(definition.QueryParameters).Get(definition.Url)
	if err != nil {
		return endpointResponse{RawData: nil, StatusCode: -1, Url: ""}, errors.CannotRequestDataError.Wrap(err, "Cannot request data from "+definition.Url)
	}
	loc := resp.RawResponse.Request.URL.String()
	log.DaemonLogger.Infof("Sent request to %s", resp.Request.URL)
	if resp.StatusCode() != 200 {
		return endpointResponse{RawData: nil, StatusCode: resp.StatusCode(), Url: loc}, errors.CannotRequestDataError.Wrap(nil, fmt.Sprintf("Cannot send request to: %s (status code %d)", definition.Url, resp.StatusCode()))
	}
	switch definition.ResponseFormat {
	case "json":
		err = json.Unmarshal(resp.Body(), &data)
		if err != nil {
			return endpointResponse{RawData: nil, StatusCode: -1, Url: loc}, errors.CannotParseDataError.Wrap(err, "Cannot parse data from: "+definition.Url)
		}
	case "xml":
		err = xml.Unmarshal(resp.Body(), &data)
		if err != nil {
			return endpointResponse{RawData: nil, StatusCode: -1, Url: loc}, errors.CannotParseDataError.Wrap(err, "Cannot parse data from: "+definition.Url)
		}
	default:
		return endpointResponse{RawData: nil, StatusCode: -1, Url: loc}, errors.UnknownFormatError.Wrap(err, "Unknown format: "+definition.ResponseFormat)
	}
	return endpointResponse{RawData: data, StatusCode: resp.StatusCode(), Url: loc}, nil
}

func normalizeVariables(definition endpointDefinition) ([]variables.EndpointParameter, error) {
	valueCount := len(util.FindFirst(definition.Variables, func(param variableSchema) bool { return !param.IsConstant }).Values)
	for _, param := range definition.Variables {
		if len(param.Values) != valueCount && !param.IsConstant {
			return nil, errors.Newf(errors.VariableValueLengthMismatchError, "Variable %s has %d values, but %d are expected", param.Name, len(param.Values), valueCount)
		}
	}
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

func parseRequests(definition endpointDefinition) ([]endpointRequest, error) {
	var requests []endpointRequest
	vars, err := normalizeVariables(definition)
	if err != nil {
		return nil, err
	}
	valueCount := len(util.FindFirst(definition.Variables, func(param variableSchema) bool { return !param.IsConstant }).Values)
	for valueCycle := 0; valueCycle < valueCount; valueCycle++ {
		variableMap := variables.VariableMap(make(map[string]any))
		for varIndex, variable := range vars {
			variableMap[definition.Variables[varIndex].Name] = variable.Value(valueCycle)
		}
		temp := template.New(definition.Name)
		temp, err := temp.Parse(definition.BaseUrl)
		if err != nil {
			return nil, errors.CannotApplyTemplateError.Wrap(err, "Cannot parse endpoint: "+definition.BaseUrl)
		}
		var stringBuilder strings.Builder
		err = temp.Execute(&stringBuilder, variableMap)
		if err != nil {
			return nil, errors.CannotApplyTemplateError.Wrap(err, "Cannot apply template to endpoint: "+definition.BaseUrl)
		}
		parsedUrl := stringBuilder.String()
		var queryParams = make(map[string]string)
		for _, queryParam := range definition.QueryParameters {
			temp, err := temp.Parse(queryParam.Value)
			if err != nil {
				return nil, errors.CannotApplyTemplateError.Wrap(err, "Cannot parse query parameter: "+queryParam.Value)
			}
			var stringBuilder strings.Builder
			err = temp.Execute(&stringBuilder, variableMap)
			if err != nil {
				return nil, errors.CannotApplyTemplateError.Wrap(err, "Cannot apply template to query parameter: "+queryParam.Value)
			}
			queryParam.Value = stringBuilder.String()
			queryParams[queryParam.Name] = queryParam.Value
		}
		requests = append(requests, endpointRequest{
			Url:             parsedUrl,
			QueryParameters: queryParams,
			ResponseFormat:  "json",
		})
	}
	return requests, nil
}

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

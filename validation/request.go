package validation

import (
	"encoding/json"
	"fmt"
	"strings"
	"text/template"

	"github.com/go-resty/resty/v2"

	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/log"
	"github.com/buonotti/odh-data-monitor/util"
)

func rest() *resty.Client {
	return resty.New().SetHeader("Accept", "*/*")
}

type EndpointRequest struct {
	Url             string
	QueryParameters map[string]string
	ResponseFormat  string
}

type EndpointResponse struct {
	StatusCode          int
	RawData             map[string]any
	Url                 string
	UsedQueryParameters map[string][]string
}

func requestData(definition EndpointRequest) (EndpointResponse, error) {
	var data map[string]any
	resp, err := rest().R().SetQueryParams(definition.QueryParameters).Get(definition.Url)
	if err != nil {
		return EndpointResponse{RawData: nil, StatusCode: -1, Url: ""}, errors.CannotRequestDataError.Wrap(err, "Cannot request data from "+definition.Url)
	}
	loc := resp.RawResponse.Request.URL.String()
	log.DaemonLogger.Infof("Sent request to %s", resp.Request.URL)
	if resp.StatusCode() != 200 {
		return EndpointResponse{RawData: nil, StatusCode: resp.StatusCode(), Url: loc}, errors.CannotRequestDataError.Wrap(nil, fmt.Sprintf("Cannot send request to: %s (status code %d)", definition.Url, resp.StatusCode()))
	}
	switch definition.ResponseFormat {
	case "json":
		err = json.Unmarshal(resp.Body(), &data)
		if err != nil {
			return EndpointResponse{RawData: nil, StatusCode: -1, Url: loc}, errors.CannotParseDataError.Wrap(err, "Cannot parse data from: "+definition.Url)
		}
	default:
		return EndpointResponse{RawData: nil, StatusCode: -1, Url: loc}, errors.UnknownFormatError.Wrap(err, "Unknown format: "+definition.ResponseFormat)
	}
	return EndpointResponse{RawData: data, StatusCode: resp.StatusCode(), Url: loc}, nil
}

func normalizeVariables(definition EndpointDefinition) ([]EndpointParameter, error) {
	valueCount := len(util.FindFirst(definition.Variables, func(param VariableSchema) bool { return !param.IsConstant }).Values)
	for _, param := range definition.Variables {
		if len(param.Values) != valueCount && !param.IsConstant {
			return nil, errors.Newf(errors.VariableValueLengthMismatchError, "Variable %s has %d values, but %d are expected", param.Name, len(param.Values), valueCount)
		}
	}
	var params []EndpointParameter
	for _, param := range definition.Variables {
		if param.IsConstant {
			params = append(params, ConstantEndpointParameter{
				value: param.Values[0],
			})
		} else {
			params = append(params, VariableEndpointParameter{
				values: param.Values,
			})
		}
	}
	return params, nil
}

func parseRequests(definition EndpointDefinition) ([]EndpointRequest, error) {
	var requests []EndpointRequest
	variables, err := normalizeVariables(definition)
	if err != nil {
		return nil, err
	}
	valueCount := len(util.FindFirst(definition.Variables, func(param VariableSchema) bool { return !param.IsConstant }).Values)
	for valueCycle := 0; valueCycle < valueCount; valueCycle++ {
		variableMap := VariableMap(make(map[string]any))
		for varIndex, variable := range variables {
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
		requests = append(requests, EndpointRequest{
			Url:             parsedUrl,
			QueryParameters: queryParams,
			ResponseFormat:  "json",
		})
	}
	return requests, nil
}

func send(requests []EndpointRequest) ([]EndpointResponse, error) {
	var responses []EndpointResponse
	for _, request := range requests {
		resp, err := requestData(request)
		if err != nil {
			return nil, err
		}
		responses = append(responses, resp)
	}
	return responses, nil
}

package validation

import (
	"encoding/xml"
	"fmt"
	"strings"
	"text/template"

	"github.com/buonotti/apisense/v2/errors"
	"github.com/buonotti/apisense/v2/log"
	"github.com/buonotti/apisense/v2/validation/definitions"
	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
)

// rest return the rest client with the set headers
func rest() *resty.Client {
	return resty.
		New().
		SetHeader("Accept", "*/*").
		SetDisableWarn(true)
}

type EndpointResponse struct {
	StatusCode int    // StatusCode is the status code of the response
	RawData    any    // RawData is the raw data of the response
	Url        string // Url is the full url of the request
}

func applyPayload(payload any, varMap definitions.VariableMap) ([]byte, error) {
	payloadString, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.CannotSerializeItemError.Wrap(err, "cannot serialize login payload")
	}
	return applyPayloadString(string(payloadString), varMap)
}

func applyPayloadString(payloadString string, varMap definitions.VariableMap) ([]byte, error) {
	templ, err := template.New("template").Parse(payloadString)
	if err != nil {
		return nil, errors.CannotApplyTemplateError.Wrap(err, "cannot parse template")
	}
	payloadBuilder := strings.Builder{}
	err = templ.Execute(&payloadBuilder, varMap)
	if err != nil {
		return nil, errors.CannotApplyTemplateError.Wrap(err, "cannot apply template")
	}
	return []byte(payloadBuilder.String()), nil
}

func retrieveToken(definition definitions.Endpoint) (string, error) {
	log.DaemonLogger().Debug("Retrieving jwt token")
	bodyContent, err := applyPayload(definition.JwtLogin.LoginPayload, definitions.VariableMap(definition.Secrets))
	if err != nil {
		return "", err
	}

	resp, err := rest().R().SetBody(bodyContent).SetHeader("Content-Type", "application/json").Post(definition.JwtLogin.Url)
	if err != nil {
		return "", errors.CannotRequestDataError.Wrap(err, "cannot send login request")
	}
	if resp.StatusCode() != 200 {
		return "", errors.CannotRequestDataError.New("login request failed with status: %d", resp.StatusCode())
	}
	var response map[string]any
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return "", errors.CannotParseDataError.Wrap(err, "cannot deserialize login response")
	}
	token := response[definition.JwtLogin.TokenKeyName].(string)
	log.DaemonLogger().Debug("Got token", "token", token)
	return token, nil
}

func prepareRequest(request EndpointRequest, definition definitions.Endpoint) (*resty.Request, error) {
	req := rest().R()

	for hName, hValue := range definition.Headers {
		req.SetHeader(hName, hValue)
	}

	req.SetQueryParams(request.QueryParameters)

	if definition.JwtLogin.Url != "" {
		var (
			token string
			err   error
		)
		token, err = retrieveToken(definition)
		if err != nil {
			return nil, err
		}

		req.SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	} else if definition.Authorization != "" {
		authHeaderVal, err := applyPayloadString(definition.Authorization, definitions.VariableMap(definition.Secrets))
		if err != nil {
			return nil, err
		}

		req.SetHeader("Authorization", string(authHeaderVal))
	}

	return req, nil
}

// requestData sends based on the given request a http GET request and returns the response
func Fetch(request EndpointRequest, definition definitions.Endpoint) (EndpointResponse, error) {
	var data any

	req, err := prepareRequest(request, definition)
	log.DaemonLogger().Debug("Prepared request for url", "url", request.Url)
	if err != nil {
		return EndpointResponse{}, err
	}

	var resp *resty.Response
	log.DaemonLogger().Debug("Starting request for url", "url", request.Url)
	switch definition.Method {
	case "POST":
		{
			payload, err := applyPayload(definition.Payload, request.Variables)
			if err != nil {
				return EndpointResponse{}, err
			}
			resp, err = req.SetBody(payload).Post(request.Url)
			break
		}
	case "PUT":
		{
			payload, err := applyPayload(definition.Payload, request.Variables)
			if err != nil {
				return EndpointResponse{}, err
			}
			resp, err = req.SetBody(payload).Post(request.Url)
			break
		}
	case "DELETE":
		{
			resp, err = req.Delete(request.Url)
			break
		}
	default:
		{
			resp, err = req.Get(request.Url)
			break
		}
	}
	log.DaemonLogger().Debug("Completed request for url", "url", request.Url)

	if err != nil {
		return EndpointResponse{}, errors.CannotRequestDataError.Wrap(err, "cannot request data from: %s", request.Url)
	}

	loc := resp.RawResponse.Request.URL.String()
	log.DaemonLogger().Info("Fetched data", "url", loc)

	okCode := 200
	if definition.OkCode > 0 {
		okCode = definition.OkCode
	}

	if resp.StatusCode() != okCode {
		return EndpointResponse{
			Url:        loc,
			StatusCode: resp.StatusCode(),
		}, nil
	}

	if definition.ResponseSchema != nil {
		switch definition.Format {
		case "json":
			err = json.Unmarshal(resp.Body(), &data)
			if err != nil {
				return EndpointResponse{}, errors.CannotParseDataError.Wrap(err, "cannot parse data from: "+request.Url)
			}
		case "xml":
			err = xml.Unmarshal(resp.Body(), &data)
			if err != nil {
				return EndpointResponse{}, errors.CannotParseDataError.Wrap(err, "cannot parse data from: "+request.Url)
			}
		default:
			return EndpointResponse{}, errors.UnknownFormatError.Wrap(err, "unknown format: "+definition.Format)
		}
	} else {
		data = make(map[string]any)
	}

	return EndpointResponse{
		RawData:    data,
		StatusCode: resp.StatusCode(),
		Url:        loc,
	}, nil
}

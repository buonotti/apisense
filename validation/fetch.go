package validation

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"strings"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/validation/definitions"
	"github.com/go-resty/resty/v2"
)

// rest return the rest client with the set headers
func rest() *resty.Client {
	return resty.
		New().
		SetHeader("Accept", "*/*").
		SetDisableWarn(true)
}

type EndpointResponse struct {
	StatusCode          int                 // StatusCode is the status code of the response
	RawData             map[string]any      // RawData is the raw data of the response mapped to a map
	Url                 string              // Url is the full url of the request
	UsedQueryParameters map[string][]string // UsedQueryParameters are the query parameters that were used in the request
}

func retrieveToken(definition definitions.Endpoint) (string, error) {
	log.DaemonLogger().Debug("Retrieving jwt token")
	payload, err := json.Marshal(definition.JwtLogin.LoginPayload)
	if err != nil {
		return "", errors.CannotSerializeItemError.Wrap(err, "cannot serialize login payload")
	}
	temp := template.New("Login" + definition.JwtLogin.Url)

	temp, err = temp.Parse(string(payload))
	if err != nil {
		return "", errors.CannotApplyTemplateError.Wrap(err, "cannot parse template")
	}

	payloadBuilder := strings.Builder{}
	err = temp.Execute(&payloadBuilder, definition.Secrets)
	if err != nil {
		return "", errors.CannotApplyTemplateError.Wrap(err, "cannot apply template")
	}
	resp, err := rest().R().SetBody(payloadBuilder.String()).SetHeader("Content-Type", "application/json").Post(definition.JwtLogin.Url)
	if err != nil {
		return "", errors.CannotRequestDataError.Wrap(err, "cannot send login request")
	}
	if resp.StatusCode() != 200 {
		return "", errors.NewF(errors.CannotRequestDataError, "login request failed with status: %d", resp.StatusCode())
	}
	var response map[string]any
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return "", errors.CannotParseDataError.Wrap(err, "cannot deserialize login response")
	}
	var token string = response[definition.JwtLogin.TokenKeyName].(string)
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
		temp := template.New("Authorization" + request.Url)
		temp, err := temp.Parse(definition.Authorization)
		if err != nil {
			return nil, errors.CannotApplyTemplateError.Wrap(err, "cannot parse tempalte")
		}
		headerBuilder := strings.Builder{}
		err = temp.Execute(&headerBuilder, definition.Secrets)
		if err != nil {
			return nil, errors.CannotApplyTemplateError.Wrap(err, "cannot apply template")
		}

		req.SetHeader("Authorization", headerBuilder.String())
	}

	return req, nil
}

// requestData sends based on the given request a http GET request and returns the response
func Fetch(request EndpointRequest, definition definitions.Endpoint) (EndpointResponse, error) {
	var data map[string]any

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
			resp, err = req.SetBody(definition.Payload).Post(request.Url)
			break
		}
	case "PUT":
		{
			resp, err = req.SetBody(definition.Payload).Post(request.Url)
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
		err = errors.SafeWrapF(errors.CannotRequestDataError, err, "cannot request data from: %s", request.Url)
		return EndpointResponse{}, err
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

	if len(definition.ResponseSchema) != 0 {
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

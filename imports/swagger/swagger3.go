package swagger

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/buonotti/apisense/v2/errors"
	"github.com/buonotti/apisense/v2/filesystem/locations"
	"github.com/buonotti/apisense/v2/log"
	"github.com/buonotti/apisense/v2/util"
	"github.com/buonotti/apisense/v2/validation/definitions"
	"github.com/goccy/go-yaml"
	"github.com/pb33f/libopenapi"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

type V3Importer struct{}

func addPathVariablesV3(host string, path string, endpoint *definitions.Endpoint) {
	pathSegments := strings.Split(path, "/")
	for i, seg := range pathSegments {
		if len(seg) > 0 && seg[0] == '{' {
			noVarSeg := seg[1 : len(seg)-1]
			varName := util.Capitalize(noVarSeg)
			pathSegments[i] = fmt.Sprintf("{{ .%s }}", varName)
			(*endpoint).Variables = append(endpoint.Variables, definitions.Variable{
				Name:       varName,
				IsConstant: true,
				Values:     []string{"TODO: edit me"},
			})
		}
	}
	(*endpoint).BaseUrl = fmt.Sprintf("http://%s%s", host, strings.Join(pathSegments, "/"))
}

func addQueryParametersV3(parameters []*v3.Parameter, endpoint *definitions.Endpoint) {
	for _, parameter := range parameters {
		if parameter.In == "query" {
			endpoint.QueryParameters = append(endpoint.QueryParameters, definitions.QueryDefinition{
				Name:  parameter.Name,
				Value: "TODO: edit me",
			})
		}
	}
}

func addResponseSchemaV3(response *v3.Response, endpoint *definitions.Endpoint) error {
	for pair := response.Content.Oldest(); pair != nil; pair = pair.Next() {
		media := pair.Key
		schema := pair.Value.Schema
		if media == "application/json" {
			rendered, err := schema.Schema().RenderInline()
			if err != nil {
				return errors.CannotSerializeItemError.Wrap(err, "cannot render schema to yaml")
			}
			var marshalled any
			err = yaml.Unmarshal(rendered, &marshalled)
			if err != nil {
				return errors.CannotUmarshalError.Wrap(err, "cannot remarshal schema")
			}

			endpoint.ResponseSchema = marshalled
			endpoint.Format = "json"

			return nil
		}
		if media == "application/xml" {
			rendered, err := schema.Schema().RenderInline()
			if err != nil {
				return errors.CannotSerializeItemError.Wrap(err, "cannot render schema to yaml")
			}
			var marshalled any
			err = yaml.Unmarshal(rendered, &marshalled)
			if err != nil {
				return errors.CannotUmarshalError.Wrap(err, "cannot remarshal schema")
			}

			endpoint.ResponseSchema = marshalled
			endpoint.Format = "xml"

			return nil
		}
	}

	return errors.InvalidContentTypeError.New("Found no supported content type in endpoint " + endpoint.Name)
}

func generateDefinitionNameV3(title string, path string, method string) string {
	titleRegex, _ := regexp.Compile(`\W`)
	title = titleRegex.ReplaceAllString(title, "")

	path, _ = strings.CutPrefix(path, "/")
	pathRegex, _ := regexp.Compile(`(/{(\w*)}|/)`)
	path = pathRegex.ReplaceAllString(path, "_$2")

	return fmt.Sprintf("%s_%s_%s", title, method, path)
}

func convertPathV3(results *[]definitions.Endpoint, title string, host string, path string, method string, operation *v3.Operation) error {
	if operation == nil || operation.Responses == nil || operation.Responses.Codes == nil {
		log.DefaultLogger().Debug("Skipping method with no response", "path", path, "method", method)
		return nil
	}
	responses := operation.Responses.Codes
	var response *v3.Response
	var okCode int
	for pair := responses.Oldest(); pair != nil; pair = pair.Next() {
		intCode, err := strconv.Atoi(pair.Key)
		if err == nil && intCode >= 200 && intCode < 300 {
			log.DefaultLogger().Debug("Found first ok code", "path", path, "method", method, "code", intCode)
			response = pair.Value
			okCode = intCode
		}
	}
	if response == nil {
		log.DefaultLogger().Warn("Skipping method on path with no response code matching 200 >= response < 300", "path", path, "method", method)
		return nil
	}

	result := definitions.Endpoint{}
	result.Version = 1
	result.FileName = generateDefinitionNameV3(title, path, method)
	result.FullPath = locations.Definition(result.FileName)
	result.OkCode = okCode
	result.Name = result.FileName
	addPathVariablesV3(host, path, &result)
	result.Method = method
	result.IsEnabled = true
	addResponseSchemaV3(response, &result)
	addQueryParametersV3(operation.Parameters, &result)
	*results = append(*results, result)
	return nil
}

func convertModelV3(model libopenapi.DocumentModel[v3.Document]) ([]definitions.Endpoint, error) {
	title := model.Model.Info.Title
	host := "undefined"
	if len(model.Model.Servers) > 0 {
		host = model.Model.Servers[0].URL
	}
	fmt.Println(title, host)
	converted := make([]definitions.Endpoint, 0)
	paths := model.Model.Paths.PathItems
	for pair := paths.Oldest(); pair != nil; pair = pair.Next() {
		err := convertPathV3(&converted, title, host, pair.Key, "GET", pair.Value.Get)
		if err != nil {
			return nil, err
		}
		err = convertPathV3(&converted, title, host, pair.Key, "POST", pair.Value.Post)
		if err != nil {
			return nil, err
		}
		err = convertPathV3(&converted, title, host, pair.Key, "PUT", pair.Value.Put)
		if err != nil {
			return nil, err
		}
		err = convertPathV3(&converted, title, host, pair.Key, "DELETE", pair.Value.Delete)
		if err != nil {
			return nil, err
		}
	}
	return converted, nil
}

func (_ *V3Importer) Import(content []byte) ([]definitions.Endpoint, error) {
	doc, err := libopenapi.NewDocument(content)
	if err != nil {
		return nil, errors.CannotUmarshalError.Wrap(err, "cannot create openapi document")
	}
	v3Model, errors := doc.BuildV3Model()
	if len(errors) > 0 {
		for _, err := range errors {
			log.DefaultLogger().Error("Error: " + err.Error())
		}
		log.DefaultLogger().Fatal("One or more errors occured while creating model", "amount", len(errors))
	}

	return convertModelV3(*v3Model)
}

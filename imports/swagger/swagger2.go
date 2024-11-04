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
	v2 "github.com/pb33f/libopenapi/datamodel/high/v2"
)

type V2Importer struct{}

func generateDefinitionNameV2(title string, path string, method string) string {
	titleRegex, _ := regexp.Compile(`\W`)
	title = titleRegex.ReplaceAllString(title, "")

	path, _ = strings.CutPrefix(path, "/")
	pathRegex, _ := regexp.Compile(`(/:(\w*)|/)`)
	path = pathRegex.ReplaceAllString(path, "_$2")

	return fmt.Sprintf("%s_%s_%s", title, method, path)
}

func setFormatV2(products []string, endpoint *definitions.Endpoint) {
	(*endpoint).Format = "json"
	for _, prod := range products {
		if prod == "application/xml" {
			(*endpoint).Format = "xml"
		}
	}
}

func addPathVariablesV2(host string, basePath string, path string, endpoint *definitions.Endpoint) {
	pathSegments := strings.Split(path, "/")
	for i, seg := range pathSegments {
		if len(seg) > 0 && seg[0] == ':' {
			noColonSeg := seg[1:]
			varName := util.Capitalize(noColonSeg)
			pathSegments[i] = fmt.Sprintf("{{ .%s }}", varName)
			(*endpoint).Variables = append(endpoint.Variables, definitions.Variable{
				Name:       varName,
				IsConstant: true,
				Values:     []string{"TODO: edit me"},
			})
		}
	}
	(*endpoint).BaseUrl = fmt.Sprintf("http://%s%s%s", host, basePath, strings.Join(pathSegments, "/"))
}

func addQueryParametersV2(parameters []*v2.Parameter, endpoint *definitions.Endpoint) {
	for _, param := range parameters {
		if param.In == "query" {
			(*endpoint).QueryParameters = append(endpoint.QueryParameters, definitions.QueryDefinition{
				Name:  param.Name,
				Value: "TODO: edit me",
			})
		}
	}
}

func addResponseSchemaV2(response *v2.Response, endpoint *definitions.Endpoint) error {
	rendered, err := response.Schema.Schema().RenderInline()
	if err != nil {
		return errors.CannotSerializeItemError.Wrap(err, "cannot render schema to yaml")
	}
	var marshalled any
	err = yaml.Unmarshal(rendered, &marshalled)
	if err != nil {
		return errors.CannotUmarshalError.Wrap(err, "cannot remarshal schema")
	}

	endpoint.ResponseSchema = marshalled

	return nil
}

func convertPathV2(results *[]definitions.Endpoint, title string, host string, basePath string, path string, method string, operation *v2.Operation) error {
	if operation == nil || operation.Responses == nil || operation.Responses.Codes == nil {
		log.DefaultLogger().Debug("Skipping method with no response", "path", path, "method", method)
		return nil
	}
	responses := operation.Responses.Codes
	var response *v2.Response
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
	result.FileName = generateDefinitionNameV2(title, path, method)
	result.FullPath = locations.Definition(result.FileName)
	result.OkCode = okCode
	result.Name = result.FileName
	addPathVariablesV2(host, path, basePath, &result)
	result.Method = method
	result.IsEnabled = true
	addResponseSchemaV2(response, &result)
	addQueryParametersV2(operation.Parameters, &result)
	setFormatV2(operation.Produces, &result)
	*results = append(*results, result)
	return nil
}

func convertModelV2(model libopenapi.DocumentModel[v2.Swagger]) ([]definitions.Endpoint, error) {
	title := model.Model.Info.Title
	host := model.Model.Host
	basePath := model.Model.BasePath
	converted := make([]definitions.Endpoint, 0)
	paths := model.Model.Paths.PathItems
	for pair := paths.Oldest(); pair != nil; pair = pair.Next() {
		err := convertPathV2(&converted, title, host, basePath, pair.Key, "GET", pair.Value.Get)
		if err != nil {
			return nil, err
		}
		err = convertPathV2(&converted, title, host, basePath, pair.Key, "POST", pair.Value.Post)
		if err != nil {
			return nil, err
		}
		err = convertPathV2(&converted, title, host, basePath, pair.Key, "PUT", pair.Value.Put)
		if err != nil {
			return nil, err
		}
		err = convertPathV2(&converted, title, host, basePath, pair.Key, "DELETE", pair.Value.Delete)
		if err != nil {
			return nil, err
		}
	}
	return converted, nil
}

func (_ *V2Importer) Import(content []byte) ([]definitions.Endpoint, error) {
	doc, err := libopenapi.NewDocument(content)
	if err != nil {
		return nil, errors.CannotUmarshalError.Wrap(err, "cannot create openapi document")
	}
	v2Model, errors := doc.BuildV2Model()
	if len(errors) > 0 {
		for _, err := range errors {
			log.DefaultLogger().Error("Error: " + err.Error())
		}
		log.DefaultLogger().Fatal("One or more errors occured while creating model", "amount", len(errors))
	}

	return convertModelV2(*v2Model)
}

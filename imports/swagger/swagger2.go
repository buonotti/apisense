package swagger

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/util"
	"gopkg.in/yaml.v3"

	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/validation/definitions"
)

type parameterSpecV2 struct {
	Name string
	In   string
}

type pathSpecV2 struct {
	Responses  map[string]responseSpecV2
	Produces   []string
	Parameters []parameterSpecV2
}

type responseSpecV2 struct {
	Description string
	Schema      map[string]any
}

type specV2 struct {
	Swagger string
	Info    struct {
		Title string
	}
	Host        string
	BasePath    string
	Paths       map[string]map[string]pathSpecV2
	Definitions map[string]map[string]any
}

type V2Importer struct{}

func setOkCodeV2(responses map[string]responseSpecV2, endpoint *definitions.Endpoint) error {
	for code := range responses {
		statusCode, err := strconv.Atoi(code)
		if err != nil {
			return err
		}
		if statusCode >= 200 && statusCode < 300 {
			(*endpoint).OkCode = statusCode
			break
		}
	}
	if endpoint.OkCode == 0 {
		(*endpoint).OkCode = 200
	}
	return nil
}

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

func addQueryParametersV2(parameters []parameterSpecV2, endpoint *definitions.Endpoint) {
	for _, param := range parameters {
		if param.In == "query" {
			(*endpoint).QueryParameters = append(endpoint.QueryParameters, definitions.QueryDefinition{
				Name:  param.Name,
				Value: "TODO: edit me",
			})
		}
	}
}

func addResponseSchemaV2(response responseSpecV2, defs map[string]map[string]any, endpoint *definitions.Endpoint) error {
	endpoint.ResponseSchema = response.Schema
	if response.Schema != nil {
		endpoint.ResponseSchema["$schema"] = "https://json-schema.org/draft/2020-12/schema"
		endpoint.ResponseSchema["definitions"] = defs
	}

	return nil
}

func convertPathMethod(title string, host string, basePath string, path string, method string, spec pathSpecV2, defs map[string]map[string]any) (definitions.Endpoint, error) {
	apiPath := fmt.Sprintf("http://%s%s%s", host, basePath, path)
	log.DefaultLogger().Info("Converting call to path", "path", apiPath, "method", method)

	result := definitions.Endpoint{}
	result.IsEnabled = true
	result.Method = strings.ToUpper(method)
	result.Variables = make([]definitions.Variable, 0)
	result.FileName = generateDefinitionNameV2(title, path, method)
	result.Name = result.FileName
	result.FullPath = filepath.FromSlash(directories.DefinitionsDirectory() + "/" + result.FileName + ".apisensedef.yml")

	err := setOkCodeV2(spec.Responses, &result)
	if err != nil {
		return definitions.Endpoint{}, err
	}
	setFormatV2(spec.Produces, &result)
	addPathVariablesV2(host, basePath, path, &result)
	addQueryParametersV2(spec.Parameters, &result)
	err = addResponseSchemaV2(spec.Responses[strconv.Itoa(result.OkCode)], defs, &result)
	if err != nil {
		return definitions.Endpoint{}, err
	}

	return result, nil
}

func convertPaths(title string, host string, basePath string, path string, spec map[string]pathSpecV2, defs map[string]map[string]any) ([]definitions.Endpoint, error) {
	converted := make([]definitions.Endpoint, 0)
	for method, pathSpec := range spec {
		conv, err := convertPathMethod(title, host, basePath, path, method, pathSpec, defs)
		if err != nil {
			return nil, err
		}
		converted = append(converted, conv)
	}
	return converted, nil
}

func (_ *V2Importer) Import(file string, content []byte) ([]definitions.Endpoint, error) {
	var swaggerSpec specV2
	if strings.HasSuffix(file, ".json") {
		err := json.Unmarshal(content, &swaggerSpec)
		if err != nil {
			return []definitions.Endpoint{}, err
		}
	} else {
		err := yaml.Unmarshal(content, &swaggerSpec)
		if err != nil {
			return []definitions.Endpoint{}, err
		}
	}

	converted := make([]definitions.Endpoint, 0)

	for name, pathSpec := range swaggerSpec.Paths {
		conv, err := convertPaths(swaggerSpec.Info.Title, swaggerSpec.Host, swaggerSpec.BasePath, name, pathSpec, swaggerSpec.Definitions)
		if err != nil {
			return nil, err
		}

		for _, c := range conv {
			converted = append(converted, c)
		}
	}

	slices.SortFunc(converted, func(e1, e2 definitions.Endpoint) int {
		urlCmp := strings.Compare(e1.BaseUrl, e2.BaseUrl)
		if urlCmp == 0 {
			return strings.Compare(e1.Method, e2.Method)
		}
		return urlCmp
	})

	for _, c := range converted {
		b, _ := yaml.Marshal(c)
		fmt.Println(c.FullPath)
		fmt.Println(string(b))
		fmt.Println("---")
	}

	return converted, nil
}

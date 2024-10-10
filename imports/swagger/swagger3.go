package swagger

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/util"
	"github.com/buonotti/apisense/validation/definitions"
	"gopkg.in/yaml.v3"
)

type pathSpecV3 struct {
	Responses  map[string]responseSpecV3
	Parameters []any
}

type responseSpecV3 struct {
	Description string
	Content     map[string]struct {
		Schema map[string]any
	}
}

type componentsSpecV3 struct {
	Parameters map[string]any
	Schemas    map[string]map[string]any
}

type specV3 struct {
	Openapi string
	Info    struct {
		Title string
	}
	Servers []struct {
		Url string
	}
	Paths      map[string]map[string]pathSpecV3
	Components componentsSpecV3
}

type V3Importer struct{}

func setOkCodeV3(responses map[string]responseSpecV3, endpoint *definitions.Endpoint) error {
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

func addQueryParametersV3(parameters []any, defs componentsSpecV3, endpoint *definitions.Endpoint) {
	// TODO
}

func addResponseSchemaV3(response responseSpecV3, defs map[string]map[string]any, endpoint *definitions.Endpoint) error {
	for media, schema := range response.Content {
		if media == "application/json" {
			endpoint.Format = "json"
			endpoint.ResponseSchema = schema.Schema
			if endpoint.ResponseSchema != nil {
				endpoint.ResponseSchema["definitions"] = defs
			}
			return nil
		} else if media == "appliaction/xml" {
			endpoint.Format = "xml"
			endpoint.ResponseSchema = schema.Schema
			if endpoint.ResponseSchema != nil {
				endpoint.ResponseSchema["definitions"] = defs
			}
			return nil
		}
	}

	return errors.InvalidContentTypeError.New("Found no supported content type in endpoint " + endpoint.Name)
}

func convertPathMethodV3(title string, host string, path string, method string, spec pathSpecV3, defs componentsSpecV3) (definitions.Endpoint, error) {
	apiPath := fmt.Sprintf("http://%s%s", host, path)
	log.DefaultLogger().Info("Converting call to path", "path", apiPath, "method", method)

	result := definitions.Endpoint{}
	result.IsEnabled = true
	result.Method = strings.ToUpper(method)
	result.Variables = make([]definitions.Variable, 0)
	result.FileName = generateDefinitionNameV3(title, path, method)
	result.Name = result.FileName
	result.FullPath = filepath.FromSlash(directories.DefinitionsDirectory() + "/" + result.FileName + ".apisensedef.yml")

	err := setOkCodeV3(spec.Responses, &result)
	if err != nil {
		return definitions.Endpoint{}, err
	}
	addPathVariablesV3(host, path, &result)
	addQueryParametersV3(spec.Parameters, defs, &result)
	err = addResponseSchemaV3(spec.Responses[strconv.Itoa(result.OkCode)], defs.Schemas, &result)
	if err != nil {
		return definitions.Endpoint{}, err
	}

	return result, nil
}

func generateDefinitionNameV3(title string, path string, method string) string {
	titleRegex, _ := regexp.Compile(`\W`)
	title = titleRegex.ReplaceAllString(title, "")

	path, _ = strings.CutPrefix(path, "/")
	pathRegex, _ := regexp.Compile(`(/:(\w*)|/)`)
	path = pathRegex.ReplaceAllString(path, "_$2")

	return fmt.Sprintf("%s_%s_%s", title, method, path)
}

func convertPathsV3(title string, host string, path string, spec map[string]pathSpecV3, defs componentsSpecV3) ([]definitions.Endpoint, error) {
	converted := make([]definitions.Endpoint, 0)
	for method, pathSpec := range spec {
		conv, err := convertPathMethodV3(title, host, path, method, pathSpec, defs)
		if err != nil {
			return nil, err
		}
		converted = append(converted, conv)
	}
	return converted, nil
}

func (_ *V3Importer) Import(file string, content []byte) ([]definitions.Endpoint, error) {
	var swaggerSpec specV3
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
	server := "/"
	if len(swaggerSpec.Servers) > 0 {
		server = swaggerSpec.Servers[0].Url
	}

	for name, pathSpec := range swaggerSpec.Paths {
		conv, err := convertPathsV3(swaggerSpec.Info.Title, server, name, pathSpec, swaggerSpec.Components)
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

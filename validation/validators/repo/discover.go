package repo

import (
	"strings"

	"github.com/buonotti/apisense/errors"
	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
)

const repoEndpoint string = "https://api.github.com/orgs/buonotti/repos"

func DiscoverTemplates() (map[string]string, error) {
	resp, err := resty.New().R().Get(repoEndpoint)
	if err != nil {
		return nil, errors.GithubUnreachableError.Wrap(err, "failed to request endpoints")
	}
	var repos []map[string]any
	err = json.Unmarshal(resp.Body(), &repos)
	if err != nil {
		return nil, errors.CannotUmarshalError.Wrap(err, "failed to umarshal github response")
	}

	res := make(map[string]string)

	for _, repo := range repos {
		if name, ok := repo["name"].(string); ok {
			if strings.HasPrefix(name, "validator-template-") {
				lang, _ := strings.CutPrefix(name, "validator-template-")
				res[lang] = repo["html_url"].(string)
			}
		} else {
			return nil, errors.InvalidGithubResponseError.New("no field 'name' of type string")
		}
	}
	return res, nil
}

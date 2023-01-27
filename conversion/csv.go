package conversion

import (
	"fmt"
	"strings"

	"github.com/buonotti/apisense/validation"
)

func Csv() Converter {
	return csvConverter{}
}

type csvConverter struct{}

func (c csvConverter) ConvertMany(reports []validation.Report) ([]byte, error) {
	lines := strings.Builder{}
	for _, r := range reports {
		d, err := c.Convert(r)
		if err != nil {
			return nil, err
		}
		lines.Write(d)
		lines.Write([]byte("\n"))
	}
	return []byte(lines.String()), nil
}

func (csvConverter) Convert(reports ...validation.Report) ([]byte, error) {
	lines := strings.Builder{}
	for _, report := range reports {
		for _, validatedEndpoint := range report.Endpoints {
			for _, endpointResult := range validatedEndpoint.Results {
				for _, validatorOutput := range endpointResult.ValidatorsOutput {
					line := fmt.Sprintf("%s;%s;%s;%s;%s;%s\n",
						report.Time.String(),
						validatedEndpoint.EndpointName,
						endpointResult.Url,
						validatorOutput.Validator,
						validatorOutput.Status,
						validatorOutput.Error,
					)
					lines.Write([]byte(line))
				}
			}
		}
	}
	return []byte(lines.String()), nil
}

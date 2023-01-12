package conversion

import (
	"fmt"

	"github.com/buonotti/odh-data-monitor/util"
	"github.com/buonotti/odh-data-monitor/validation"
)

func Csv() Converter {
	return csvConverter{}
}

type csvConverter struct{}

func (csvConverter) Convert(report validation.Report) ([]byte, error) {
	lines := make([]string, 1)
	header := "date;endpoint;url;validator;status;error"
	lines[0] = header
	for _, validatedEndpoint := range report.Results {
		for _, endpointResult := range validatedEndpoint.Results {
			for _, validatorOutput := range endpointResult.ValidatorsOutput {
				line := fmt.Sprintf("%s;%s;%s;%s;%s;%s",
					report.Time.String(),
					validatedEndpoint.EndpointName,
					endpointResult.Url,
					validatorOutput.Validator,
					validatorOutput.Status,
					validatorOutput.Error,
				)
				lines = append(lines, line)
			}
		}
	}
	return []byte(util.Join(lines, "\n")), nil
}

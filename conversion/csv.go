package conversion

import (
	"fmt"
	"strings"

	"github.com/buonotti/apisense/v2/validation/pipeline"
)

// Csv returns a new csvConverter
func Csv() Converter {
	return csvConverter{}
}

type csvConverter struct{}

func (csvConverter) Convert(reports ...pipeline.Report) ([]byte, error) {
	lines := strings.Builder{}
	for _, report := range reports {
		for _, validatedEndpoint := range report.Endpoints {
			for _, endpointResult := range validatedEndpoint.TestCaseResults {
				for _, validatorOutput := range endpointResult.ValidatorResults {
					line := fmt.Sprintf("%s;%s;%s;%s;%s;%s\n",
						report.Time.String(),
						validatedEndpoint.EndpointName,
						endpointResult.Url,
						validatorOutput.Name,
						validatorOutput.Status,
						validatorOutput.Message,
					)
					lines.Write([]byte(line))
				}
			}
		}
	}
	return []byte(lines.String()), nil
}

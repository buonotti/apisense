package controllers

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/buonotti/apisense/api/filter"
	"github.com/buonotti/apisense/conversion"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/util"
	"github.com/buonotti/apisense/validation/pipeline"
)

// AllReports godoc
//
//	@Summary		Get all the reports
//	@Description	Gets a list of all reports that can be filtered with a query
//	@ID				all-reports
//	@Tags			reports
//	@Param			where	query		string	false	"Query in the format: field.op.value (optional)"
//	@Param			format	query		string	false	"Return format: json or csv (default: json)"
//	@Success		200		array		validation.Report
//	@Failure		500		{object}	ErrorResponse
//	@Router			/reports [get]
func AllReports(c *gin.Context) {
	allReports, err := pipeline.Reports()
	if err != nil {
		c.AbortWithStatusJSON(500, ErrorResponse{Message: err.Error()})
		return
	}

	whereFilter, err := filter.ParseFromContext[pipeline.Report](c)
	allReports = whereFilter.Apply(allReports)

	if err != nil {
		c.AbortWithStatusJSON(500, ErrorResponse{Message: err.Error()})
		return
	}

	writeFormattedReport(c, allReports...)
}

// Report godoc
//
//	@Summary		Get one report
//	@Description	Gets a single report identified by his id
//	@ID				report
//	@Tags			reports
//	@Param			format	query		string	false	"json"
//	@Param			id		path		string	true	"qNg8rJX"
//	@Success		200		{object}	validation.Report
//	@Failure		404		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/reports/:id [get]
func Report(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := errors.IdRequiredError.New("")
		c.AbortWithStatusJSON(400, ErrorResponse{Message: err.Error()})
		return
	}

	reports, err := pipeline.Reports()
	if err != nil {
		c.AbortWithStatusJSON(500, ErrorResponse{Message: err.Error()})
		return
	}

	report := util.FindFirst(reports, func(report pipeline.Report) bool {
		return report.Id == id
	})

	if report == nil {
		err = errors.CannotFindReportError.New("cannot find report with id: " + id)
		c.AbortWithStatusJSON(404, ErrorResponse{Message: err.Error()})
		return
	}

	writeFormattedReport(c, *report)
}

func writeFormattedReport(c *gin.Context, reports ...pipeline.Report) {
	body := strings.Builder{}
	format := c.Query("format")
	if format == "" {
		format = "json"
	}

	formatter := conversion.Get(format)
	if formatter == nil {
		err := errors.UnknownFormatError.New("unknown format: " + format)
		c.AbortWithStatusJSON(400, ErrorResponse{Message: err.Error()})
		return
	}

	convertedReports, err := formatter.Convert(reports...)
	if err != nil {
		c.AbortWithStatusJSON(500, ErrorResponse{Message: err.Error()})
	}

	body.Write(convertedReports)
	c.Data(200, "application/"+format, []byte(body.String()))
}

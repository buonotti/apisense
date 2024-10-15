package controllers

import (
	"net/http"
	"strings"

	"github.com/buonotti/apisense/api/filter"
	"github.com/buonotti/apisense/conversion"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/util"
	"github.com/buonotti/apisense/validation/pipeline"
	"github.com/gofiber/fiber/v2"
)

// AllReports godoc
//
//	@Summary		Get all the reports
//	@Description	Gets a list of all reports that can be filtered with a query
//	@ID				all-reports
//	@Tags			reports
//	@Produce		json
//	@Param			where	query		string	false	"Query in the format: field.op.value (optional)"
//	@Param			format	query		string	false	"Return format: json or csv (default: json)"
//	@Success		200		array		pipeline.Report
//	@Failure		500		{object}	ErrorResponse
//	@Router			/reports [get]
func AllReports(c *fiber.Ctx) error {
	allReports, err := pipeline.Reports()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Err(err))
	}

	whereFilter, err := filter.ParseFromContext[pipeline.Report](c)
	allReports = whereFilter.Apply(allReports)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Err(err))
	}

	return writeFormattedReport(c, allReports...)
}

// Report godoc
//
//	@Summary		Get one report
//	@Description	Gets a single report identified by his id
//	@ID				report
//	@Tags			reports
//	@Produce		json
//	@Param			format	query		string	false	"json"
//	@Param			id		path		string	true	"qNg8rJX"
//	@Success		200		{object}	pipeline.Report
//	@Failure		404		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/reports/:id [get]
func Report(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		err := errors.IdRequiredError.New("")
		return c.Status(http.StatusBadRequest).JSON(Err(err))
	}

	reports, err := pipeline.Reports()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Err(err))
	}

	report := util.FindFirst(reports, func(report pipeline.Report) bool {
		return report.Id == id
	})

	if report == nil {
		err = errors.CannotFindReportError.New("cannot find report with id: " + id)
		return c.Status(http.StatusNotFound).JSON(Err(err))
	}

	return writeFormattedReport(c, *report)
}

func writeFormattedReport(c *fiber.Ctx, reports ...pipeline.Report) error {
	body := strings.Builder{}
	format := c.Query("format")
	if format == "" {
		format = "json"
	}

	formatter := conversion.Get(format)
	if formatter == nil {
		err := errors.UnknownFormatError.New("unknown format: " + format)
		return c.Status(http.StatusBadRequest).JSON(ErrorResponse{Message: err.Error()})
	}

	convertedReports, err := formatter.Convert(reports...)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Err(err))
	}

	body.Write(convertedReports)
	c.Set("Content-Type", "application/"+format)
	return c.Status(200).Send([]byte(body.String()))
}

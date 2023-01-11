package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/buonotti/odh-data-monitor/util"
	"github.com/buonotti/odh-data-monitor/validation"
)

func AllReports(c *gin.Context) {
	allReports, err := validation.Reports()
	allReports = filterReports(c, allReports)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, allReports)
}

func filterReports(c *gin.Context, reports []validation.Report) []validation.Report {
	return reports // TODO
}

func Report(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"message": "id is required"})
		return
	}
	reports, err := validation.Reports()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	report := util.FindFirst(reports, func(report validation.Report) bool {
		return report.Id == id
	})
	if report == nil {
		c.AbortWithStatusJSON(404, gin.H{"message": "report not found"})
		return
	}
	c.JSON(200, report)
}

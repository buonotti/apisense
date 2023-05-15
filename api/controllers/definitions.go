package controllers

import (
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/util"
	"github.com/buonotti/apisense/validation/definitions"
	"github.com/gin-gonic/gin"
)

// AllDefinitions godoc
//
//	@Summary		Get all the definitions
//	@Description	Gets a list of all definitions
//	@ID				all-definitions
//	@Tags			definitions
//	@Success		200	array		definitions.Endpoint
//	@Failure		500	{object}	ErrorResponse
//	@Router			/definitions [get]
func AllDefinitions(c *gin.Context) {
	allDefinitions, err := definitions.Endpoints()
	if err != nil {
		c.AbortWithStatusJSON(500, ErrorResponse{Message: err.Error()})
		return
	}

	if len(allDefinitions) == 0 {
		allDefinitions = []definitions.Endpoint{}
	}

	c.JSON(200, allDefinitions)
}

// Definition godoc
//
//	@Summary		Get one definition
//	@Description	Gets a single definition identified by his endpoint name
//	@ID				definition
//	@Tags			definitions
//	@Param			name	path		string	true	"Bluetooth"
//	@Success		200		{object}	definitions.Endpoint
//	@Failure		404		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/definitions/:name [get]
func Definition(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		err := errors.NameRequiredError.New("")
		c.AbortWithStatusJSON(400, ErrorResponse{Message: err.Error()})
		return
	}

	defs, err := definitions.Endpoints()
	if err != nil {
		c.AbortWithStatusJSON(500, ErrorResponse{Message: err.Error()})
		return
	}

	definition := util.FindFirst(defs, func(endpoint definitions.Endpoint) bool {
		return endpoint.Name == name
	})

	if definition == nil {
		err := errors.CannotFindDefinitionError.New("")
		c.AbortWithStatusJSON(404, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(200, definition)
}

package controllers

import (
	"os"
	"path/filepath"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/util"
	"github.com/buonotti/apisense/validation/definitions"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

// AllDefinitions godoc
//
//	@Summary		Get all the definitions
//	@Description	Gets a list of all definitions
//	@ID				all-definitions
//	@Tags			definitions
//	@Security		ApiKeyAuth
//	@param			Authorization	header	string	true	"Authorization"
//	@Produce		json
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
//	@Security		ApiKeyAuth
//	@param			Authorization	header	string	true	"Authorization"
//	@Produce		json
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

// CreateDefinition godoc
//
//	@Summary		Create a definition
//	@Description	Creates a new definition
//	@ID				create-definition
//	@Tags			definitions
//	@Security		ApiKeyAuth
//	@param			Authorization	header	string	true	"Authorization"
//	@Accept			json
//	@Produce		json
//	@Param			definition	body		definitions.Endpoint	true	"Endpoint definition"
//	@Success		200			{object}	definitions.Endpoint
//	@Failure		400			{object}	ErrorResponse
//	@Failure		500			{object}	ErrorResponse
//	@Router			/definitions [post]
func CreateDefinition(c *gin.Context) {
	var definition definitions.Endpoint
	if !requestValid(c, &definition) {
		return
	}

	allDefinitions, err := definitions.Endpoints()
	if err != nil {
		c.AbortWithStatusJSON(500, ErrorResponse{Message: err.Error()})
		return
	}

	if util.FindFirst(allDefinitions, func(endpoint definitions.Endpoint) bool {
		return endpoint.Name == definition.Name
	}) != nil {
		err := errors.DefinitionAlreadyExistsError.New("")
		c.AbortWithStatusJSON(409, ErrorResponse{Message: err.Error()})
		return
	}

	fileName := filepath.FromSlash(directories.DefinitionsDirectory() + "/" + definition.Name + ".apisensedef.yml")
	data, err := yaml.Marshal(definition)
	if err != nil {
		c.AbortWithStatusJSON(500, ErrorResponse{Message: err.Error()})
		return
	}

	err = os.WriteFile(fileName, data, os.ModePerm)
	if err != nil {
		c.AbortWithStatusJSON(500, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(200, definition)
}

package controllers

import (
	"net/http"
	"os"

	"github.com/buonotti/apisense/v2/errors"
	"github.com/buonotti/apisense/v2/filesystem/locations"
	"github.com/buonotti/apisense/v2/util"
	"github.com/buonotti/apisense/v2/validation/definitions"
	"github.com/goccy/go-yaml"
	"github.com/gofiber/fiber/v2"
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
func AllDefinitions(c *fiber.Ctx) error {
	allDefinitions, err := definitions.Endpoints()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Err(err))
	}

	if len(allDefinitions) == 0 {
		allDefinitions = []definitions.Endpoint{}
	}

	return c.JSON(allDefinitions)
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
func Definition(c *fiber.Ctx) error {
	name := c.Params("name")
	if name == "" {
		err := errors.NameRequiredError.New("")
		return c.Status(http.StatusBadRequest).JSON(Err(err))
	}

	defs, err := definitions.Endpoints()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Err(err))
	}

	definition := util.FindFirst(defs, func(endpoint definitions.Endpoint) bool {
		return endpoint.Name == name
	})

	if definition == nil {
		err := errors.CannotFindDefinitionError.New("")
		return c.Status(http.StatusNotFound).JSON(Err(err))
	}

	return c.JSON(definition)
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
//	@Success		201			{object}	definitions.Endpoint
//	@Failure		400			{object}	ErrorResponse
//	@Failure		500			{object}	ErrorResponse
//	@Router			/definitions [post]
func CreateDefinition(c *fiber.Ctx) error {
	var definition definitions.Endpoint
	if err := requestValid(c, &definition); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Err(err))
	}

	allDefinitions, err := definitions.Endpoints()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Err(err))
	}

	if util.FindFirst(allDefinitions, func(endpoint definitions.Endpoint) bool {
		return endpoint.Name == definition.Name
	}) != nil {
		err := errors.DefinitionAlreadyExistsError.New("")
		return c.Status(http.StatusConflict).JSON(Err(err))
	}

	err = definitions.ValidateDefinition(&definition)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Err(err))
	}

	fileName := locations.Definition(definition.Name)
	data, err := yaml.Marshal(definition)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Err(err))
	}

	err = os.WriteFile(fileName, data, os.ModePerm)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Err(err))
	}

	return c.Status(http.StatusCreated).JSON(definition)
}

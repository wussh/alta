package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"gofrendi/structureExample/model"

	"github.com/labstack/echo/v4"
)

type PersonController struct {
	model model.PersonModel
}

func NewPersonController(m model.PersonModel) PersonController {
	return PersonController{model: m}
}

func (pc PersonController) GetAll(c echo.Context) error {
	allPersons, err := pc.model.GetAll()
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot get persons")
	}
	return c.JSON(http.StatusOK, allPersons)
}

func (pc PersonController) Add(c echo.Context) error {
	var person model.Person
	if err := c.Bind(&person); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid person data")
	}
	person, err := pc.model.Add(person)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot add person")
	}
	return c.JSON(http.StatusOK, person)
}

func (pc PersonController) Edit(c echo.Context) error {
	personId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid person id")
	}
	var person model.Person
	if err := c.Bind(&person); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid person data")
	}
	person, err = pc.model.Edit(personId, person)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot edit person")
	}
	return c.JSON(http.StatusOK, person)
}

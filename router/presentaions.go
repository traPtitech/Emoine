package router

import (
	"net/http"
	"strconv"

	"github.com/FujishigeTemma/Emoine/repository"
	"github.com/labstack/echo/v4"
)

type PostPresentationsStruct struct {
	Name        string `json:"name"`
	Speakers    string `json:"speakers"`
	Description string `json:"description"`
}

// GetPresentations GET /presentations
func (h *Handlers) GetPresentations(c echo.Context) error {
	presentations, err := h.Repo.GetPresentations()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, presentations)
}

// PostPresentations POST /presentations
func (h *Handlers) PostPresentations(c echo.Context) error {
	posted := PostPresentationsStruct{}
	if err := c.Bind(&posted); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	createStruct := repository.CreatePresentation{
		Name:        posted.Name,
		Speakers:    posted.Speakers,
		Description: posted.Description,
	}

	err := h.Repo.CreatePresentation(&createStruct)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

// GetPresentations GET /presentations
func (h *Handlers) GetPresentation(c echo.Context) error {
	presentationID, err := strconv.Atoi(c.Param("presentationID"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	presentation, err := h.Repo.GetPresentation(presentationID)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, presentation)
}

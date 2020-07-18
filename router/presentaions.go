package router

import (
	"net/http"
	"strconv"

	"github.com/FujishigeTemma/Emoine/repository"
	"github.com/FujishigeTemma/Emoine/utils"
	"github.com/labstack/echo/v4"
)

type PostPresentationsStruct struct {
	Name        string `json:"name"`
	Speakers    string `json:"speakers"`
	Description string `json:"description"`
}

type PatchPresentationsStruct struct {
	Name        utils.String `json:"name"`
	Speakers    utils.String `json:"speakers"`
	Description utils.String `json:"description"`
	Prev        utils.Int    `json:"prev"`
	Next        utils.Int    `json:"next"`
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

// GetPresentation GET /presentations/:presentationID
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

// PatchPresentation PATCH /presentations/:presentationID
func (h *Handlers) PatchPresentation(c echo.Context) error {
	presentationID, err := strconv.Atoi(c.Param("presentationID"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	presentation, err := h.Repo.GetPresentation(presentationID)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	patchStruct := PatchPresentationsStruct{}
	if err := c.Bind(&patchStruct); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	presentation.Name = patchStruct.Name
	presentation.Description = patchStruct.Description
	presentation.Speakers = patchStruct.Speakers

	if err = h.Repo.UpdatePresentation(presentation); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, presentation)
}



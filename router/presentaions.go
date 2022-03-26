package router

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/traPtitech/Emoine/repository"
	"github.com/traPtitech/Emoine/utils"
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
}

// GetPresentations GET /presentations
func (h *Handlers) GetPresentations(c echo.Context) error {
	presentations, err := h.repo.GetPresentations()
	if err != nil {
		return err
	}
	if presentations == nil {
		return c.JSON(http.StatusOK, []repository.Presentation{})
	}
	return c.JSON(http.StatusOK, presentations)
}

// PostPresentations POST /presentations
func (h *Handlers) PostPresentations(c echo.Context) error {
	var req PostPresentationsStruct
	if err := c.Bind(&req); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	err := h.repo.CreatePresentation(req.Name, req.Speakers, req.Description)
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

	presentation, err := h.repo.GetPresentation(presentationID)
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

	presentation, err := h.repo.GetPresentation(presentationID)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	var req PatchPresentationsStruct
	if err := c.Bind(&req); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if req.Name.Valid {
		presentation.Name = req.Name
	}
	if req.Description.Valid {
		presentation.Description = req.Description
	}
	if req.Speakers.Valid {
		presentation.Speakers = req.Speakers
	}

	if err = h.repo.UpdatePresentation(presentation.Name.String, presentation.Speakers.String, presentation.Description.String, int(presentation.Prev.Int64), int(presentation.Next.Int64), presentationID); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, presentation)
}

// DeletePresentation DELETE /presentations/:presentationID
func (h *Handlers) DeletePresentation(c echo.Context) error {
	presentationID, err := strconv.Atoi(c.Param("presentationID"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	err = h.repo.DeletePresentation(presentationID)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

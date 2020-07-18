package router

import (
	"net/http"
	"strconv"

	"github.com/FujishigeTemma/Emoine/repository"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
)

type PostReviewStruct struct {
	Skill        int `json:"skill" db:"skill"`
	Artistry     int `json:"artistry" db:"artistry"`
	Idea         int `json:"idea" db:"idea"`
	Presentation int `json:"presentation" db:"presentation"`
}

// GetPresentationReview GET /presentations/:presentationID/review
func (h *Handlers) GetPresentationReview(c echo.Context) error {
	presentationID, err := strconv.Atoi(c.Param("presentationID"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	res, err := h.Repo.GetReviewStatistics(presentationID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

// PostPresentationReview POST /presentations/:presentationID/review
func (h *Handlers) PostPresentationReview(c echo.Context) error {
	presentationID, err := strconv.Atoi(c.Param("presentationID"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	_, err = h.Repo.GetPresentation(presentationID)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	posted := PostReviewStruct{}
	if err := c.Bind(&posted); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	
	userID := uuid.Nil
	createReview := repository.Review {
		UserId: userID,
		PresentationId: presentationID,
		Score: repository.Score{
			Skill: posted.Skill,
			Artistry: posted.Artistry,
			Idea: posted.Idea,
			Presentation: posted.Presentation,
		},
	}

	err = h.Repo.CreateReview(&createReview)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, createReview)
}

package router

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/traPtitech/Emoine/repository"
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

	_, err = h.Repo.GetPresentation(presentationID)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
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

	userID, err := getUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}
	createReview := repository.Review{
		UserId:         userID,
		PresentationId: presentationID,
		Score: repository.Score{
			Skill:        posted.Skill,
			Artistry:     posted.Artistry,
			Idea:         posted.Idea,
			Presentation: posted.Presentation,
		},
	}

	isExist, err := h.Repo.IsExistReview(userID, presentationID)
	if err != nil {
		return err
	}
	if isExist {
		return c.NoContent(http.StatusConflict)
	}

	err = h.Repo.CreateReview(&createReview)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, createReview)
}

// PatchPresentationReview PATCH /presentations/:presentationID/review
func (h *Handlers) PatchPresentationReview(c echo.Context) error {
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

	userID, err := getUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}
	updateReview := repository.Review{
		UserId:         userID,
		PresentationId: presentationID,
		Score: repository.Score{
			Skill:        posted.Skill,
			Artistry:     posted.Artistry,
			Idea:         posted.Idea,
			Presentation: posted.Presentation,
		},
	}

	err = h.Repo.UpdateReview(&updateReview)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

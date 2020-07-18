package router

import (
	"database/sql"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/FujishigeTemma/Emoine/repository"
	"github.com/labstack/echo/v4"
)

type PostPresentationsStruct struct {
	Name        string `json:"name"`
	Speakers    string `json:"speakers"`
	Description string `json:"description"`
}

type PatchPresentationsStruct struct {
	Name        *string `json:"name"`
	Speakers    *string `json:"speakers"`
	Description *string `json:"description"`
	Prev				*int			`json:"prev"`
	Next				*int			`json:"next"`
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

	rtPatchStruct := reflect.TypeOf(patchStruct)
	rvPatchStruct := reflect.ValueOf(patchStruct)

	rvPresentation := reflect.ValueOf(&presentation).Elem()

	for i := 0; i < rtPatchStruct.NumField(); i++ {
		f := rtPatchStruct.Field(i)
		v := rvPatchStruct.FieldByName(f.Name).Interface()
		fmt.Println(v)
		str := ""
		integer := 0
		if v != nil {
			if f.Type == reflect.TypeOf(&str) {
				fmt.Println("string")
				rvPresentation.FieldByName(f.Name).Set(reflect.ValueOf(sql.NullString{ String: v.(string), Valid: true }))
			}
			if f.Type == reflect.TypeOf(&integer) {
				fmt.Println("int")
				rvPresentation.FieldByName(f.Name).Set(reflect.ValueOf(sql.NullInt32{ Int32: v.(int32), Valid: true }))
			}
		}
	}
	fmt.Println(presentation)

	if err = h.Repo.UpdatePresentation(presentation); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, presentation)
}

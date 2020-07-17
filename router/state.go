package router

import (
	"github.com/FujishigeTemma/Emoine/repository"
	"github.com/labstack/echo/v4"
)

// PostState POST /state
func (h *Handlers) PostState(c echo.Context) error {
	state := &repository.State{}


	h.Repo.UpdateState()
}

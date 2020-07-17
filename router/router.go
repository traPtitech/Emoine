package router

import (
	"github.com/FujishigeTemma/Emoine/repository"
	"github.com/labstack/echo"
)

type Handlers struct {
	Repo           repository.Repository
}

func Setup(r repository.Repository) *echo.Echo {
	e := echo.New()
	h := &Handlers{
		Repo: r,
	}

	api := e.Group("/api")
	{
		apiPresentaions := api.Group("/presentations")
		{
			apiPresentaions.GET("", h.GetPresentations)
		}
	}
	return e
}

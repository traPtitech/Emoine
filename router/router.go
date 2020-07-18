package router

import (
	"github.com/FujishigeTemma/Emoine/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handlers struct {
	Repo repository.Repository
}

func Setup(repo repository.Repository) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	h := &Handlers{
		Repo: repo,
	}

	api := e.Group("/api")
	{
		apiPresentations := api.Group("/presentations")
		{
			apiPresentations.GET("", h.GetPresentations)
			apiPresentations.POST("", h.PostPresentations)
			apiPresentationsID := apiPresentations.Group("/:presentationID")
			{
				apiPresentationsID.GET("", h.GetPresentation)
			}
		}
	}
	return e
}

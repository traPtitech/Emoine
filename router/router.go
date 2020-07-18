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

	s := NewStreamer(repo)

	api := e.Group("/api")
	{
		apiPresentations := api.Group("/presentations")
		{
			apiPresentations.GET("", h.GetPresentations)
			apiPresentations.POST("", h.PostPresentations)
			apiPresentationsID := apiPresentations.Group("/:presentationID")
			{
				apiPresentationsID.GET("", h.GetPresentation)
				apiPresentationsID.PATCH("", h.PatchPresentation)
				apiPresentationsID.DELETE("", h.DeletePresentation)
				apiPresentationsID.GET("/reaction", h.GetPresentationReaction)
				apiPresentationsID.GET("/review", h.GetPresentationReview)
			}
		}
		api.Group("/ws").GET("", func(c echo.Context) error {
			s.ServeHTTP(c)
			return nil
		})
	}
	return e
}

package router

import (
	"github.com/FujishigeTemma/Emoine/repository"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"strings"
)

type Handlers struct {
	Repo          repository.Repository
	SessionOption sessions.Options
	ClientID      string
}

func Setup(repo repository.Repository) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	h := &Handlers{
		Repo: repo,
		SessionOption: sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 30,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		},
		ClientID: os.Getenv("CLIENT_ID"),
	}
	e.Use(h.WatchCallbackMiddleware())

	// s := NewStreamer()

	api := e.Group("/api", h.IsTraQUserMiddleware)
	{
		isAdmin := h.AdminUserMiddleware

		apiPresentations := api.Group("/presentations")
		{
			apiPresentations.GET("", h.GetPresentations)
			apiPresentations.POST("", h.PostPresentations, isAdmin)
			apiPresentationsID := apiPresentations.Group("/:presentationID")
			{
				apiPresentationsID.GET("", h.GetPresentation)
				apiPresentationsID.PATCH("", h.PatchPresentation, isAdmin)
				apiPresentationsID.DELETE("", h.DeletePresentation, isAdmin)
				apiPresentationsID.GET("/reaction", h.GetPresentationReaction, isAdmin)
				apiPresentationsID.GET("/review", h.GetPresentationReview, isAdmin)
				apiPresentationsID.POST("/review", h.PostPresentationReview)
			}
		}
		// api.Group("/ws").GET("", func(c echo.Context) error {
		// 	s.ServeHTTP(c)
		// 	return nil
		// })
	}
	e.POST("/api/oauth2/code", h.GetGeneratedCode)

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Request().URL.Path, "/api")
		},
		Root:  "web/dist",
		HTML5: true,
	}))
	return e
}

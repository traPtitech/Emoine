package router

import (
	"net/http"
	"os"
	"strings"

	"github.com/FujishigeTemma/Emoine/repository"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handlers struct {
	Repo repository.Repository
	SessionOption     sessions.Options
	ClientID          string
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
		ClientID:   os.Getenv("CLIENT_ID"),
	}
	e.Use(h.WatchCallbackMiddleware())

	// s := NewStreamer()

	api := e.Group("/api", h.TraQUserMiddleware)
	{
		apiPresentations := api.Group("/presentations")
		{
			apiPresentations.GET("", h.GetPresentations)
			apiPresentations.POST("", h.PostPresentations)
			apiPresentationsID := apiPresentations.Group("/:presentationID")
			{
				apiPresentationsID.GET("", h.GetPresentation)
				apiPresentationsID.PATCH("", h.PatchPresentation)
			}
		}
		// api.Group("/ws").GET("", func(c echo.Context) error {
		// 	s.ServeHTTP(c)
		// 	return nil
		// })
	}
	e.POST("/api/authParams", h.HandlePostAuthParams)

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Request().URL.Path, "/api")
		},
		Root:  "web/dist",
		HTML5: true,
	}))
	return e
}

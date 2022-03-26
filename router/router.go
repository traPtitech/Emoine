package router

import (
	"github.com/traPtitech/Emoine/services/streamer"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/traPtitech/Emoine/repository"
)

type Handlers struct {
	repo          repository.Repository
	streamer      *streamer.Streamer
	SessionOption sessions.Options
	clientID      string
}

func Setup(repo repository.Repository) *echo.Echo {
	setDefaultStateData()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SECRET")))))
	s := streamer.NewStreamer(repo)

	h := &Handlers{
		repo: repo,
		SessionOption: sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 30,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		},
		clientID: os.Getenv("CLIENT_ID"),
		streamer: s,
	}

	api := e.Group("/api")
	{
		isAdmin := func(next echo.HandlerFunc) echo.HandlerFunc {
			return next
		}

		// TODO: グループだと動かない
		api.GET("/live-id", h.GetLiveID)
		api.PUT("/live-id", h.PutLiveID, isAdmin)
		api.GET("/viewer", h.GetViewer)

		api.POST("/state", h.PostState, isAdmin)

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
				apiPresentationsID.PUT("/review", h.PutPresentationReview)
				apiPresentationsID.GET("/comments", h.GetPresentationComments, isAdmin)
			}
		}
		api.GET("/ws", h.ConnectWebSocket)
	}
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Request().URL.Path, "/api")
		},
		Root:  "web/dist",
		HTML5: true,
	}))
	return e
}

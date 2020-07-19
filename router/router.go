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
	Repo          repository.Repository
	stream *Streamer
	SessionOption sessions.Options
	ClientID      string
}

func Setup(repo repository.Repository) *echo.Echo {
	setDefaultStateData()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	s := NewStreamer(repo)

	h := &Handlers{
		Repo: repo,
		SessionOption: sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 30,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		},
		ClientID: os.Getenv("CLIENT_ID"),
		stream: s,
	}
	e.Use(h.WatchCallbackMiddleware)


	api := e.Group("/api", h.IsTraQUserMiddleware)
	{
		isAdmin := h.AdminUserMiddleware

		// TODO: グループだと動かない
		api.GET("/live-id", h.GetLiveID)
		api.PUT("/live-id", h.PutLiveID, isAdmin)

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
				apiPresentationsID.POST("/review", h.PostPresentationReview)
				apiPresentationsID.PATCH("/review", h.PatchPresentationReview)
				apiPresentationsID.GET("/comments", h.GetPresentationComments, isAdmin)
			}
		}
		api.GET("/users/me", h.GetUserMe)
		api.GET("/ws", func(c echo.Context) error {
			s.ServeHTTP(c)
			return nil
		})
	}
	e.GET("/api/oauth2/code", h.GetGeneratedCode)
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Request().URL.Path, "/api")
		},
		Root:  "web/dist",
		HTML5: true,
	}))
	return e
}

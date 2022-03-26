package router

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const (
	sessionKey = "emoine_session"
	userIDKey  = "userID"
)

var sessionOptions = &sessions.Options{
	MaxAge:   86400 * 7,
	Secure:   true,
	HttpOnly: true,
	SameSite: http.SameSiteLaxMode,
}

func (h *Handlers) SessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := getSession(c)
		if err != nil {
			if err := setSession(c); err != nil {
				return c.NoContent(http.StatusInternalServerError)
			}
		}
		return next(c)
	}
}

func getSession(c echo.Context) (uuid.UUID, error) {
	sess, err := session.Get(sessionKey, c)
	if err != nil {
		return uuid.Nil, err
	}
	userID, ok := sess.Values[userIDKey]
	if !ok {
		return uuid.Nil, errors.New("userID not found")
	}
	return uuid.Parse(userID.(string))
}

func setSession(c echo.Context) error {
	sess, _ := session.Get(sessionKey, c)
	sess.Options = sessionOptions
	sess.Values[userIDKey] = uuid.New().String()
	return sess.Save(c.Request(), c.Response())
}

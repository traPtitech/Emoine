package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/traPtitech/Emoine/repository"
	"github.com/traPtitech/Emoine/utils"
)

type ResponseToken struct {
	Token string `json:"token"`
}

func (h *Handlers) PostToken(c echo.Context) error {
	userID, err := getRequestUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	generatedToken, err := utils.GenerateToken()
	if err != nil {
		return err
	}

	if deletedCount, err := h.Repo.CleanupExpiredTokens(); err != nil {
		log.Warnf("Failed to cleanup expired tokens: %v\n", err)
	} else {
		log.Infof("Cleaned expired %d tokens\n", deletedCount)
	}

	token := repository.Token{
		Token:  generatedToken,
		UserID: userID,
	}
	if err := h.Repo.CreateToken(&token); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ResponseToken{Token: generatedToken})
}

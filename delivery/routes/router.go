package routes

import (
	_authHandler "potentivio-app/delivery/handler/auth"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/login/artist", ah.LoginArtistHandler())
	e.POST("/login/cafe-owner", ah.LoginCafeHandler())
}

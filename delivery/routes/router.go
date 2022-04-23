package routes

import (
	_artistHandler "potentivio-app/delivery/handler/artist"
	_authHandler "potentivio-app/delivery/handler/auth"
	_middlewares "potentivio-app/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/login/artist", ah.LoginArtistHandler())
	e.POST("/login/cafe-owner", ah.LoginCafeHandler())
}

func RegisterArtistPath(e *echo.Echo, ah *_artistHandler.ArtistHandler) {
	e.POST("/artist", ah.CreateArtistHandler(), _middlewares.JWTMiddleware())
	e.GET("/artist/:id", ah.GetArtistByIdHandler(), _middlewares.JWTMiddleware())
}

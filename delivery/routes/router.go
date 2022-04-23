package routes

import (
	_authHandler "potentivio-app/delivery/handler/auth"
	_cafeHandler "potentivio-app/delivery/handler/cafe"
	_middlewares "potentivio-app/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/login/artist", ah.LoginArtistHandler())
	e.POST("/login/cafe-owner", ah.LoginCafeHandler())
}
func RegisterCafePath(e *echo.Echo, ch *_cafeHandler.CafeHandler) {
	e.GET("/cafe/:id", ch.GetCafeByIdHandler(), _middlewares.JWTMiddleware())
	e.GET("/cafe/profile", ch.GetCafeProfileHandler(), _middlewares.JWTMiddleware())
}

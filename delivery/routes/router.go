package routes

import (
	_artistHandler "potentivio-app/delivery/handler/artist"
	_authHandler "potentivio-app/delivery/handler/auth"
	_cafeHandler "potentivio-app/delivery/handler/cafe"
	_hirehandler "potentivio-app/delivery/handler/hire"

	_catagoryHandler "potentivio-app/delivery/handler/catagory"
	_imageCafeHandler "potentivio-app/delivery/handler/imageCafe"

	_middlewares "potentivio-app/delivery/middlewares"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/login/artist", ah.LoginArtistHandler())
	e.POST("/login/cafe-owner", ah.LoginCafeHandler())
}

func RegisterCafePath(e *echo.Echo, ch *_cafeHandler.CafeHandler) {
	e.GET("/cafe/:id", ch.GetCafeByIdHandler(), _middlewares.JWTMiddleware())
	e.GET("/cafe/profile", ch.GetCafeProfileHandler(), _middlewares.JWTMiddleware())
}

func RegisterArtistPath(e *echo.Echo, ah *_artistHandler.ArtistHandler) {
	e.POST("/artist", ah.CreateArtistHandler())
	e.GET("/artist/profile", ah.GetProfileArtistHandler(), _middlewares.JWTMiddleware())
	e.GET("/artist", ah.GetAllArtistHandler(), _middlewares.JWTMiddleware())
	e.GET("/artist/:id", ah.GetArtistByIdHandler(), _middlewares.JWTMiddleware())
	e.PUT("/artist/:id", ah.UpdateArtistHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/artist/:id", ah.DeleteArtistHandler(), _middlewares.JWTMiddleware())
}

func RegisterCatagoryPath(e *echo.Echo, ch *_catagoryHandler.CatagoryHandler) {
	e.GET("/catagory", ch.GetAllCatagoryHandler())
	e.POST("/catagory", ch.CreateCatagoryHandler())
}

func RegisterImageCafePath(e *echo.Echo, ich *_imageCafeHandler.ImageCafeHandler) {
	e.POST("/image/cafe", ich.CreateImageCafeHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/image/cafe/:id", ich.DeleteImageCafeHandler(), _middlewares.JWTMiddleware())

}

func HireArtistPath(e *echo.Echo, hh *_hirehandler.HireHandler) {
	e.POST("hire/:id", hh.CreateHire(), _middlewares.JWTMiddleware())
	e.GET("hire/artist", hh.GetHireByIdArtis(), _middlewares.JWTMiddleware())

}

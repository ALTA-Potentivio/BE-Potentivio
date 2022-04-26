package routes

import (
	_artistHandler "potentivio-app/delivery/handler/artist"
	_authHandler "potentivio-app/delivery/handler/auth"
	_cafeHandler "potentivio-app/delivery/handler/cafe"
	_hirehandler "potentivio-app/delivery/handler/hire"
	_videoHandler "potentivio-app/delivery/handler/videoArtist"

	"github.com/labstack/echo/v4"

	_catagoryHandler "potentivio-app/delivery/handler/catagory"
	_genreHandler "potentivio-app/delivery/handler/genre"
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
	e.POST("/cafe", ch.PostCafeHandler())
	e.GET("/cafe", ch.GetAllCafeHandler())
	e.DELETE("/cafe/:id", ch.DeleteCafeHandler())
	e.PUT("/cafe/:id", ch.UpdateCafeHandler(), _middlewares.JWTMiddleware())
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
	e.GET("/catagory", ch.GetAllCatagoryHandler(), _middlewares.JWTMiddleware())
	e.POST("/catagory", ch.CreateCatagoryHandler(), _middlewares.JWTMiddleware())
}

func RegisterGenrePath(e *echo.Echo, gh *_genreHandler.GenreHandler) {
	e.GET("/genre", gh.GetAllGenreHandler(), _middlewares.JWTMiddleware())
	e.POST("/genre", gh.CreateGenreHandler(), _middlewares.JWTMiddleware())
}

func RegisterImageCafePath(e *echo.Echo, ich *_imageCafeHandler.ImageCafeHandler) {
	e.POST("/image/cafe", ich.CreateImageCafeHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/image/cafe/:id", ich.DeleteImageCafeHandler(), _middlewares.JWTMiddleware())

}

func HireArtistPath(e *echo.Echo, hh *_hirehandler.HireHandler) {
	e.POST("hire/:id", hh.CreateHire(), _middlewares.JWTMiddleware())
	e.GET("hire/artist", hh.GetHireByIdArtis(), _middlewares.JWTMiddleware())
	e.GET("hire/cafe", hh.GetHireByIdCafe(), _middlewares.JWTMiddleware())
	e.POST("/accept", hh.AcceptHire(), _middlewares.JWTMiddleware())
	e.PUT("/cancel", hh.CancelHireByCafe(), _middlewares.JWTMiddleware())

}

func RegisterVideoArtistPath(e *echo.Echo, ich *_videoHandler.VideoHandler) {
	e.POST("/video/artist", ich.PostVideoHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/video/artist/:id", ich.DeleteVideoHandler(), _middlewares.JWTMiddleware())

}

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
	_notifHandler "potentivio-app/delivery/handler/notification"

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
	e.GET("/cafe", ch.GetAllCafeHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/cafe/profile", ch.DeleteCafeHandler(), _middlewares.JWTMiddleware())
	e.PUT("/cafe/profile", ch.UpdateCafeHandler(), _middlewares.JWTMiddleware())
}

func RegisterArtistPath(e *echo.Echo, ah *_artistHandler.ArtistHandler) {
	e.POST("/artist", ah.CreateArtistHandler())
	e.GET("/artist/profile", ah.GetProfileArtistHandler(), _middlewares.JWTMiddleware())
	e.GET("/artist", ah.GetAllArtistHandler(), _middlewares.JWTMiddleware())
	e.GET("/artist/:id", ah.GetArtistByIdHandler(), _middlewares.JWTMiddleware())
	e.PUT("/artist/profile", ah.UpdateArtistHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/artist/profile", ah.DeleteArtistHandler(), _middlewares.JWTMiddleware())
}

func RegisterCategoryPath(e *echo.Echo, ch *_catagoryHandler.CatagoryHandler) {
	e.GET("/category", ch.GetAllCatagoryHandler(), _middlewares.JWTMiddleware())
	e.POST("/category", ch.CreateCatagoryHandler(), _middlewares.JWTMiddleware())
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
	e.POST("/hire/:id", hh.CreateHire(), _middlewares.JWTMiddleware())
	e.GET("/hire/artist", hh.GetHireByIdArtis(), _middlewares.JWTMiddleware())
	e.GET("/hire/cafe", hh.GetHireByIdCafe(), _middlewares.JWTMiddleware())
	e.POST("/accept/:id", hh.AcceptHire(), _middlewares.JWTMiddleware())
	e.PUT("/cafe/cancel/:id", hh.CancelHireByCafe(), _middlewares.JWTMiddleware())
	e.PUT("/reject/:id", hh.RejectHire(), _middlewares.JWTMiddleware())
	e.PUT("/artist/cancel/:id", hh.CancelHireByArtis(), _middlewares.JWTMiddleware())
	e.PUT("/rating/:id", hh.Rating(), _middlewares.JWTMiddleware())
	// callback jangan di hit lewat postman
	e.POST("/callback", hh.CallBack())
	e.PUT("/cafe/done/:id", hh.Done(), _middlewares.JWTMiddleware())
}

func RegisterVideoArtistPath(e *echo.Echo, ich *_videoHandler.VideoHandler) {
	e.POST("/video/artist", ich.PostVideoHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/video/artist/:id", ich.DeleteVideoHandler(), _middlewares.JWTMiddleware())
}

func RegisterNotificationPath(e *echo.Echo, nh *_notifHandler.NotifHandler) {
	e.POST("/offer/:id", nh.CreateNotifHandler(), _middlewares.JWTMiddleware())
	e.GET("/offer", nh.GetAllNotifByIdCafe(), _middlewares.JWTMiddleware())
	e.DELETE("/offer/:id", nh.DeleteNotifHandler(), _middlewares.JWTMiddleware())
}

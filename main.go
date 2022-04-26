package main

import (
	"fmt"
	"log"
	"net/http"
	"potentivio-app/configs"

	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"

	_authHandler "potentivio-app/delivery/handler/auth"
	_authRepository "potentivio-app/repository/auth"
	_authUseCase "potentivio-app/usecase/auth"

	_cafeHandler "potentivio-app/delivery/handler/cafe"
	_cafeRepository "potentivio-app/repository/cafe"
	_cafeUseCase "potentivio-app/usecase/cafe"

	_artistHandler "potentivio-app/delivery/handler/artist"
	_artistRepository "potentivio-app/repository/artist"
	_artistUseCase "potentivio-app/usecase/artist"

	_hireHandler "potentivio-app/delivery/handler/hire"
	_hireRepository "potentivio-app/repository/hire"
	_hireUseCase "potentivio-app/usecase/hire"

	_catagoryHandler "potentivio-app/delivery/handler/catagory"
	_catagoryRepository "potentivio-app/repository/catagory"
	_catagoryUseCase "potentivio-app/usecase/catagory"

	_genreHandler "potentivio-app/delivery/handler/genre"
	_genreRepository "potentivio-app/repository/genre"
	_genreUseCase "potentivio-app/usecase/genre"

	_imageCafeHandler "potentivio-app/delivery/handler/imageCafe"
	_imageCafeRepository "potentivio-app/repository/imageCafe"
	_imageCafeUseCase "potentivio-app/usecase/imageCafe"

	_videoArtistHandler "potentivio-app/delivery/handler/videoArtist"
	_videoArtistRepository "potentivio-app/repository/videoArtist"
	_videoArtistUseCase "potentivio-app/usecase/videoArtist"

	_routes "potentivio-app/delivery/routes"
	_utils "potentivio-app/utils"
)

func main() {
	config := configs.GetConfig()
	db := _utils.InitDB(config)

	authRepo := _authRepository.NewAuthRepository(db)
	authUseCase := _authUseCase.NewAuthUseCase(authRepo)
	authHandler := _authHandler.NewAuthHandler(authUseCase)

	cafeRepo := _cafeRepository.NewCafeRepository(db)
	cafeUseCase := _cafeUseCase.NewCafeUseCase(cafeRepo)
	cafeHandler := _cafeHandler.NewCafeHandler(cafeUseCase)

	hireRepoArtist := _hireRepository.NewHireRepository(db)
	artistRepo := _artistRepository.NewArtistRepository(db)
	artistUseCase := _artistUseCase.NewArtistUseCase(artistRepo, hireRepoArtist)
	artistHandler := _artistHandler.NewArtistHandler(artistUseCase)

	hireRepo := _hireRepository.NewHireRepository(db)
	hireUseCase := _hireUseCase.NewHireUseCase(hireRepo, artistRepo, cafeRepo)
	hireHandler := _hireHandler.NewHireHandler(hireUseCase)

	catagoryRepo := _catagoryRepository.NewCatagoryRepository(db)
	catagoryUseCase := _catagoryUseCase.NewCatagoryUseCase(catagoryRepo)
	catagoryHandler := _catagoryHandler.NewCatagoryHandler(catagoryUseCase)

	genreRepo := _genreRepository.NewGenreRepository(db)
	genreUseCase := _genreUseCase.NewGenreUseCase(genreRepo)
	genreHandler := _genreHandler.NewGenreHandler(genreUseCase)

	imageCafeRepo := _imageCafeRepository.NewImageCafeRepository(db)
	imageCafeUseCase := _imageCafeUseCase.NewImageCafeUseCase(imageCafeRepo)
	imageCafeHandler := _imageCafeHandler.NewImageCafeHandler(imageCafeUseCase)

	videoArtistRepo := _videoArtistRepository.NewVideoRepository(db)
	videoArtistUseCase := _videoArtistUseCase.NewVideoUseCase(videoArtistRepo)
	videoArtistHandler := _videoArtistHandler.NewVideoHandler(videoArtistUseCase)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	_routes.RegisterAuthPath(e, authHandler)
	_routes.RegisterCafePath(e, cafeHandler)
	_routes.RegisterArtistPath(e, artistHandler)
	_routes.RegisterCatagoryPath(e, &catagoryHandler)
	_routes.RegisterImageCafePath(e, imageCafeHandler)
	_routes.HireArtistPath(e, hireHandler)
	_routes.RegisterGenrePath(e, &genreHandler)
	_routes.RegisterVideoArtistPath(e, videoArtistHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}

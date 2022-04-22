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

	_routes "potentivio-app/delivery/routes"
	_utils "potentivio-app/utils"
)

func main() {
	config := configs.GetConfig()
	db := _utils.InitDB(config)

	authRepo := _authRepository.NewAuthRepository(db)
	authUseCase := _authUseCase.NewAuthUseCase(authRepo)
	authHandler := _authHandler.NewAuthHandler(authUseCase)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	_routes.RegisterAuthPath(e, authHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}

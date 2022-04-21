package main

import (
	"fmt"
	"log"
	"net/http"
	"potentivio-app/configs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_utils "potentivio-app/utils"
)

func main() {
	config := configs.GetConfig()
	db := _utils.InitDB(config)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}

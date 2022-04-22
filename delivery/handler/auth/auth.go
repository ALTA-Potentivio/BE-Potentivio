package auth

import (
	"fmt"
	"net/http"
	"potentivio-app/delivery/helper"
	_entities "potentivio-app/entities"
	_authUseCase "potentivio-app/usecase/auth"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authUseCase _authUseCase.AuthUseCaseInterface
}

func NewAuthHandler(auth _authUseCase.AuthUseCaseInterface) *AuthHandler {
	return &AuthHandler{
		authUseCase: auth,
	}
}

func (ah *AuthHandler) LoginArtistHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var login _entities.Artist
		err := c.Bind(&login)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error bind data"))
		}
		token, idArtist, errorLogin := ah.authUseCase.LoginArtist(login.Email, login.Password)
		if errorLogin != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(fmt.Sprintf("%v", errorLogin)))
		}
		responseToken := map[string]interface{}{
			"token":     token,
			"id_artist": idArtist,
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success login", responseToken))
	}
}

func (ah *AuthHandler) LoginCafeHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var login _entities.Cafe
		err := c.Bind(&login)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error bind data"))
		}
		token, idCafe, errorLogin := ah.authUseCase.LoginCafe(login.Email, login.Password)
		if errorLogin != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(fmt.Sprintf("%v", errorLogin)))
		}
		responseToken := map[string]interface{}{
			"token":   token,
			"id_cafe": idCafe,
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success login", responseToken))
	}
}

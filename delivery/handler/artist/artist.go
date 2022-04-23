package artist

import (
	"net/http"
	"potentivio-app/delivery/helper"
	"potentivio-app/entities"
	_artistUseCase "potentivio-app/usecase/artist"

	"github.com/labstack/echo/v4"
)

type ArtistHandler struct {
	artistUseCase _artistUseCase.ArtistUseCaseInterface
}

func NewArtistHandler(artistUseCase _artistUseCase.ArtistUseCaseInterface) *ArtistHandler {
	return &ArtistHandler{
		artistUseCase: artistUseCase,
	}
}

func (ah *ArtistHandler) CreateArtistHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newArtist entities.Artist
		err := c.Bind(&newArtist)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error()))
		}
		_, error := ah.artistUseCase.CreateArtist(newArtist)
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(error.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success create user"))
	}
}

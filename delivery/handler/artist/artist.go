package artist

import (
	"net/http"
	"potentivio-app/delivery/helper"
	"potentivio-app/entities"
	_artistUseCase "potentivio-app/usecase/artist"
	"strconv"

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

func (ah *ArtistHandler) GetArtistByIdHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}

		artist, hireNotAvailable, hireHistory, rows, err := ah.artistUseCase.GetArtistById(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}

		notAvailable := []map[string]interface{}{}
		for i := 0; i < len(hireNotAvailable); i++ {
			response := map[string]interface{}{
				"date": hireNotAvailable[i].Date,
			}
			notAvailable = append(notAvailable, response)
		}

		history := []map[string]interface{}{}
		for i := 0; i < len(hireHistory); i++ {
			response := map[string]interface{}{
				"cafe_name": hireHistory[i].Cafe.Name,
				"date":      hireHistory[i].Date,
				"rating":    hireHistory[i].Rating,
				"comment":   hireHistory[i].Comment,
			}
			history = append(history, response)
		}

		responseArtist := map[string]interface{}{
			"id":             artist.ID,
			"artist_name":    artist.Name,
			"id_catagory":    artist.IdCatagory,
			"id_genre":       artist.IdGenre,
			"phone_number":   artist.PhoneNumber,
			"address":        artist.Address,
			"price":          artist.Price,
			"description":    artist.Description,
			"account_number": artist.AccountNumber,
			"avatar":         artist.Avatar,
			"video_artist":   artist.VideoArtist,
			"not_available":  notAvailable,
			"hire_history":   history,
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get artist by id", responseArtist))
	}
}

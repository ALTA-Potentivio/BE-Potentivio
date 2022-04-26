package genre

import (
	"net/http"
	"potentivio-app/delivery/helper"
	"potentivio-app/entities"
	_genreUseCase "potentivio-app/usecase/genre"

	"github.com/labstack/echo/v4"
)

type GenreHandler struct {
	genreUseCase _genreUseCase.GenreUseCaseInterface
}

func NewGenreHandler(genreUseCase _genreUseCase.GenreUseCaseInterface) GenreHandler {
	return GenreHandler{
		genreUseCase: genreUseCase,
	}
}

func (gh *GenreHandler) CreateGenreHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newGenre entities.Genre
		err := c.Bind(&newGenre)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error()))
		}
		_, error := gh.genreUseCase.CreateGenre(newGenre)
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(error.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success to create genre"))
	}
}

func (gh *GenreHandler) GetAllGenreHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		genre, err := gh.genreUseCase.GetAllGenre()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}

		responseGenre := []map[string]interface{}{}
		for i := 0; i < len(genre); i++ {
			response := map[string]interface{}{
				"id":         genre[i].ID,
				"name_genre": genre[i].NameGenre,
			}
			responseGenre = append(responseGenre, response)
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success to get all genre", responseGenre))
	}
}

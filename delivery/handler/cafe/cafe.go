package cafe

import (
	"net/http"
	"potentivio-app/delivery/helper"
	_middlewares "potentivio-app/delivery/middlewares"
	_cafeUseCase "potentivio-app/usecase/cafe"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CafeHandler struct {
	cafeUseCase _cafeUseCase.CafeUseCaseInterface
}

func NewCafeHandler(cafe _cafeUseCase.CafeUseCaseInterface) *CafeHandler {
	return &CafeHandler{
		cafeUseCase: cafe,
	}
}
func (ch *CafeHandler) GetCafeByIdHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)

		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("failed to extract token"))
		}

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to convert id param"))
		}
		if idToken != id {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("Unauthorized"))
		}

		cafes, rows, err := ch.cafeUseCase.GetCafeById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("succes to get detail cafe", cafes))
	}
}
func (ch *CafeHandler) GetCafeProfileHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)

		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("failed to extract token"))
		}

		var id = idToken
		if idToken != id {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("Unauthorized"))
		}
		cafes, rows, err := ch.cafeUseCase.GetCafeById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("succes to get detail cafe", cafes))
	}
}

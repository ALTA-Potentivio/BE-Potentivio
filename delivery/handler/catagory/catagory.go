package catagory

import (
	"net/http"
	"potentivio-app/delivery/helper"
	"potentivio-app/entities"
	_catagoryUseCase "potentivio-app/usecase/catagory"

	"github.com/labstack/echo/v4"
)

type CatagoryHandler struct {
	catagoryUseCase _catagoryUseCase.CatagoryUseCaseInterface
}

func NewCatagoryHandler(catagoryUseCase _catagoryUseCase.CatagoryUseCaseInterface) CatagoryHandler {
	return CatagoryHandler{
		catagoryUseCase: catagoryUseCase,
	}
}

func (ch *CatagoryHandler) CreateCatagoryHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newCatagory entities.Catagory
		err := c.Bind(&newCatagory)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error()))
		}
		_, error := ch.catagoryUseCase.CreateCatagory(newCatagory)
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(error.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success to create catagory"))
	}
}

func (ch *CatagoryHandler) GetAllCatagoryHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		catagory, err := ch.catagoryUseCase.GetAllCatagory()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}

		responseCatagories := []map[string]interface{}{}
		for i := 0; i < len(catagory); i++ {
			response := map[string]interface{}{
				"id":            catagory[i].ID,
				"name_catagory": catagory[i].NameCatagory,
			}
			responseCatagories = append(responseCatagories, response)
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success to get all catagory", responseCatagories))
	}
}

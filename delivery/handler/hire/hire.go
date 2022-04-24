package hire

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"potentivio-app/delivery/helper"
	"potentivio-app/delivery/middlewares"
	"potentivio-app/entities"
	"potentivio-app/usecase/hire"
	"strconv"
	"time"
)

type HireHandler struct {
	hireUseCase hire.HireUseCaseInterface
}

func NewHireHandler(hireUseCase hire.HireUseCaseInterface) *HireHandler {
	return &HireHandler{
		hireUseCase: hireUseCase,
	}

}

func (hh *HireHandler) CreateHire() echo.HandlerFunc {
	return func(c echo.Context) error {
		var hireDate HireRequest
		var hire entities.Hire

		c.Bind(&hireDate)
		layoutFormat := "2006-01-02"
		value := hireDate.Date
		date, _ := time.Parse(layoutFormat, value)

		var id, _ = strconv.Atoi(c.Param("id"))
		var CafeID, _ = middlewares.ExtractToken(c)

		hire.IdArtist = uint(id)
		hire.IdCafe = uint(CafeID)
		hire.Date = date

		err := hh.hireUseCase.CreateHire(hire)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success hire artist"))
	}
}

func (hh *HireHandler) GetHireByIdArtis() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id_artist"))
		id, _ = middlewares.ExtractToken(c)

		hires, err := hh.hireUseCase.GetHireByIdArtis(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get hire by id", hires))
	}
}

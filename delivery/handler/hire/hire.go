package hire

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"potentivio-app/delivery/helper"
	"potentivio-app/delivery/middlewares"
	"potentivio-app/entities"
	"potentivio-app/usecase/hire"
	"strconv"
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
		var hire entities.Hire
		c.Bind(&hire)
		var id, _ = strconv.Atoi(c.Param("id"))
		var CafeID, _ = middlewares.ExtractToken(c)

		hire.IdArtist = uint(id)
		hire.IdCafe = uint(CafeID)

		err := hh.hireUseCase.CreateHire(hire)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success hire artist"))
	}
}

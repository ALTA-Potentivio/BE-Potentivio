package hire

import (
	"fmt"
	"net/http"
	"os"
	"potentivio-app/delivery/helper"
	"potentivio-app/delivery/middlewares"
	"potentivio-app/entities"
	"potentivio-app/usecase/hire"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
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
		hires, err := hh.hireUseCase.GetHireByIdArtist(id)
		results := []HireResponse{}

		for i := 0; i < len(hires); i++ {
			result := HireResponse{
				Id:           int(hires[i].ID),
				CafeName:     hires[i].Cafe.Name,
				Comment:      hires[i].Comment,
				Date:         fmt.Sprint(hires[i].Date),
				StatusArtist: hires[i].StatusArtist,
			}

			results = append(results, result)
		}

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get hire by id", results))
	}
}

func (hh *HireHandler) GetHireByIdCafe() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id_cafe"))

		id, _ = middlewares.ExtractToken(c)
		hires, err := hh.hireUseCase.GetHireByIdCafe(id)
		results := []HireCafeResponse{}

		for i := 0; i < len(hires); i++ {
			result := HireCafeResponse{
				Id:         int(hires[i].ID),
				ArtisName:  hires[i].Artist.Name,
				Date:       fmt.Sprint(hires[i].Date),
				StatusCafe: hires[i].StatusCafe,
				PaymentUrl: hires[i].PaymentUrl,
			}

			results = append(results, result)
		}

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get hire by id", results))
	}
}

func (hh *HireHandler) AcceptHire() echo.HandlerFunc {
	return func(c echo.Context) error {
		var hire entities.Hire
		var id, _ = strconv.Atoi(c.Param("id"))

		IdArtist, _ := middlewares.ExtractToken(c)
		hire.IdArtist = uint(IdArtist)
		hire.ID = uint(id)
		err := hh.hireUseCase.AcceptHire(hire)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success to accept"))
	}

}

func (hh *HireHandler) CancelHireByCafe() echo.HandlerFunc {
	return func(c echo.Context) error {
		var hire entities.Hire
		var id, _ = strconv.Atoi(c.Param("id"))

		c.Bind(&hire)
		IdCafe, _ := middlewares.ExtractToken(c)
		hire.IdCafe = uint(IdCafe)
		hire.ID = uint(id)
		err := hh.hireUseCase.CancelHireByCafe(hire)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("canceled"))
	}

}

func (hh *HireHandler) RejectHire() echo.HandlerFunc {
	return func(c echo.Context) error {
		var hire entities.Hire
		var id, _ = strconv.Atoi(c.Param("id"))

		IdArtist, _ := middlewares.ExtractToken(c)
		hire.IdArtist = uint(IdArtist)
		hire.ID = uint(id)
		err := hh.hireUseCase.Rejecthire(hire)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("rejected"))
	}

}

func (hh *HireHandler) CancelHireByArtis() echo.HandlerFunc {
	return func(c echo.Context) error {
		var hires entities.Hire
		var id, _ = strconv.Atoi(c.Param("id"))

		idArtist, _ := middlewares.ExtractToken(c)
		hires.ID = uint(id)
		hires.IdArtist = uint(idArtist)
		err := hh.hireUseCase.CancelHireByArtis(hires)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("canceled"))

	}
}

func (hh *HireHandler) Rating() echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, errconv := strconv.Atoi(idStr)
		if errconv != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}

		var hire entities.Hire
		c.Bind(&hire)
		hire.ID = uint(id)
		err := hh.hireUseCase.Rating(hire)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("succes to give rating"))
	}

}

func (hh *HireHandler) CallBack() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		headers := req.Header

		callBackToken := headers.Get("X-Callback-Token")

		if callBackToken != os.Getenv("CALLBACK_KEY") {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		var callback CallBackRequest
		c.Bind(&callback)

		callBackData := entities.Hire{
			Invoice:      callback.Invoice,
			StatusArtist: callback.Status,
			StatusCafe:   callback.Status,
		}

		err := hh.hireUseCase.CallBack(callBackData)

		return err
	}

}

func (hh *HireHandler) Done() echo.HandlerFunc {
	return func(c echo.Context) error {
		var hires entities.Hire
		var id, _ = strconv.Atoi(c.Param("id"))

		idCafe, _ := middlewares.ExtractToken(c)
		hires.ID = uint(id)
		hires.IdCafe = uint(idCafe)
		err := hh.hireUseCase.Done(hires)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("done"))

	}
}

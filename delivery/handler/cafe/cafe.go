package cafe

import (
	"net/http"
	"potentivio-app/delivery/helper"
	_middlewares "potentivio-app/delivery/middlewares"
	_entities "potentivio-app/entities"
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

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to convert id param"))
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

func (ch *CafeHandler) PostCafeHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		var cafe _entities.Cafe
		err := c.Bind(&cafe)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error()))
		}

		error := ch.cafeUseCase.PostCafe(cafe)
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(error.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("succes to create cafe-owner"))
	}
}
func (ch *CafeHandler) GetAllCafeHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		filters := map[string]string{}
		if c.QueryParam("name") != "" {
			filters["name"] = c.QueryParam("name")
		}

		cafes, error := ch.cafeUseCase.GetAllCafe(filters)
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(error.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("succes to get all cafe ownner", cafes))
	}
}
func (ch *CafeHandler) DeleteCafeHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)

		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("failed to extract token"))
		}

		error := ch.cafeUseCase.DeleteCafe(idToken)
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(error.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("succes to delete cafe ownner"))
	}
}

func (ch *CafeHandler) UpdateCafeHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		var updateCafe _entities.Cafe
		updateCafe.ID = uint(idToken)
		errBind := c.Bind(&updateCafe)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(errBind.Error()))
		}

		fileData, fileInfo, err_binding_image := c.Request().FormFile("avatar")
		if err_binding_image != http.ErrMissingFile {
			if err_binding_image != nil {
				return c.JSON(http.StatusBadRequest, helper.ResponseFailed("bind image error"))
			}

			_, err_check_extension := helper.CheckFileExtension(fileInfo.Filename)
			if err_check_extension != nil {
				return c.JSON(http.StatusBadRequest, helper.ResponseFailed("file extension error"))
			}

			err_check_size := helper.CheckFileSize(fileInfo.Size)
			if err_check_size != nil {
				return c.JSON(http.StatusBadRequest, helper.ResponseFailed("file size error"))
			}

			fileName := "foto_profile_" + strconv.Itoa(idToken)

			var err_upload_photo error
			theUrl, err_upload_photo := helper.UploadImage("foto_profile_cafe", fileName, fileData)
			if err_upload_photo != nil {
				return c.JSON(http.StatusBadRequest, helper.ResponseFailed("upload image failed"))
			}

			updateCafe.Avatar = &theUrl
		}

		err := ch.cafeUseCase.UpdateCafe(updateCafe)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success to update cafe"))
	}
}

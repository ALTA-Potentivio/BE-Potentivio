package imageCafe

import (
	"net/http"
	"potentivio-app/delivery/helper"
	_middlewares "potentivio-app/delivery/middlewares"
	"potentivio-app/entities"
	_imageCafeUseCase "potentivio-app/usecase/imageCafe"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ImageCafeHandler struct {
	imageCafeUseCase _imageCafeUseCase.ImageCafeUseCaseInterface
}

func NewImageCafeHandler(imageCafeUseCase _imageCafeUseCase.ImageCafeUseCaseInterface) *ImageCafeHandler {
	return &ImageCafeHandler{
		imageCafeUseCase: imageCafeUseCase,
	}
}

func (ich *ImageCafeHandler) CreateImageCafeHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		//mendapatkan id dari token yang login
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		var newImageCafe entities.ImageCafe
		err := c.Bind(&newImageCafe)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error()))
		}

		// memasukkan id cafe dengan id token yang login
		newImageCafe.IdCafe = uint(idToken)

		idname, err := ich.imageCafeUseCase.GetImageIDLast()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}

		// prosess binding image
		fileData, fileInfo, err_binding_image := c.Request().FormFile("image_url")
		if err_binding_image != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("bind image error"))
		}

		// check file extension
		_, err_check_extension := helper.CheckFileExtension(fileInfo.Filename)
		if err_check_extension != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("file extension error"))
		}

		// check file size
		err_check_size := helper.CheckFileSize(fileInfo.Size)
		if err_check_size != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("file size error"))
		}

		// memberikan nama file
		fileName := "image_cafe_" + strconv.Itoa(idToken) + "_" + strconv.Itoa(int(idname))

		// upload foto profile
		var err_upload_photo error
		theUrl, err_upload_photo := helper.UploadImage("image_cafe", fileName, fileData)
		if err_upload_photo != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("upload image failed"))
		}

		// create foto profile artist
		newImageCafe.ImageUrl = theUrl

		_, error := ich.imageCafeUseCase.CreateImageCafe(newImageCafe)
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(error.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success to upload image cafe"))
	}
}

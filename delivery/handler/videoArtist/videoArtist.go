package videoArtist

import (
	"net/http"
	"potentivio-app/delivery/helper"
	_middlewares "potentivio-app/delivery/middlewares"
	_entities "potentivio-app/entities"
	_videoUseCase "potentivio-app/usecase/videoArtist"
	"strconv"

	"github.com/labstack/echo/v4"
)

type VideoHandler struct {
	videoUseCase _videoUseCase.VideoUseCaseInterface
}

func NewVideoHandler(video _videoUseCase.VideoUseCaseInterface) *VideoHandler {
	return &VideoHandler{
		videoUseCase: video,
	}
}
func (vh *VideoHandler) PostVideoHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		var video _entities.VideoArtist
		err := c.Bind(&video)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error()))
		}

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("failed to extract token"))
		}

		name, errPrice := _middlewares.ExtractTokenName(c)
		if errPrice != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("failed to extract token"))
		}

		error := vh.videoUseCase.PostVideo(video, idToken, name)
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(error.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("succes to create video-owner"))
	}
}

func (vh *VideoHandler) DeleteVideoHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to convert id param"))
		}

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("failed to extract token"))
		}

		name, errPrice := _middlewares.ExtractTokenName(c)
		if errPrice != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("failed to extract token"))
		}

		error := vh.videoUseCase.DeleteVideo(id, idToken, name)
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(error.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("succes to delete video ownner"))
	}
}

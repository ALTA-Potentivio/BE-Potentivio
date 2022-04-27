package notification

import (
	"net/http"
	"potentivio-app/delivery/helper"
	_middlewares "potentivio-app/delivery/middlewares"
	"potentivio-app/entities"
	_notifUseCase "potentivio-app/usecase/notification"
	"strconv"

	"github.com/labstack/echo/v4"
)

type NotifHandler struct {
	notifUseCase _notifUseCase.NotifUseCaseInterface
}

func NewNotifHandler(notif _notifUseCase.NotifUseCaseInterface) *NotifHandler {
	return &NotifHandler{
		notifUseCase: notif,
	}
}
func (nh *NotifHandler) CreateNotifHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		//mendapatkan id dari token yang login
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to convert id param"))
		}

		var notif entities.Notification

		_, errNotif := nh.notifUseCase.CreateNotif(notif, uint(idToken), uint(id))
		if errNotif != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("succes to create notification"))
	}
}

func (nh *NotifHandler) GetAllNotifByIdCafe() echo.HandlerFunc {
	return func(c echo.Context) error {
		//mendapatkan id dari token yang login
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		notif, rows, err := nh.notifUseCase.GetAllNotifByIdCafe(uint(idToken))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}

		responseNotif := []map[string]interface{}{}
		for i := 0; i < len(notif); i++ {
			response := map[string]interface{}{
				"id":            notif[i].ID,
				"id_artist":     notif[i].Artist.ID,
				"artist_name":   notif[i].Artist.Name,
				"price":         notif[i].Artist.Price,
				"avatar":        notif[i].Artist.Avatar,
				"name_catagory": notif[i].Artist.Catagory.NameCatagory,
				"name_genre":    notif[i].Artist.Genre.NameGenre,
			}
			responseNotif = append(responseNotif, response)
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success to get all notification", responseNotif))
	}
}

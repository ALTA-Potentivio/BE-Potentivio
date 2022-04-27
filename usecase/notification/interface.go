package notification

import (
	_entities "potentivio-app/entities"
)

type NotifUseCaseInterface interface {
	CreateNotif(notif _entities.Notification, idToken uint, idCafe uint) (_entities.Notification, error)
	GetAllNotifByIdCafe(idToken uint) ([]_entities.Notification, uint, error)
}

package notification

import (
	_entities "potentivio-app/entities"
)

type NotifRepositoryInterface interface {
	CreateNotif(notif _entities.Notification) (_entities.Notification, error)
	GetAllNotifByIdCafe(idToken uint) ([]_entities.Notification, uint, error)
	GetAllNotifByIdArtist(idToken uint) ([]_entities.Notification, uint, error)
	GetNotifById(id uint) (_entities.Notification, int, error)
	DeleteNotif(id uint) (uint, error)
}

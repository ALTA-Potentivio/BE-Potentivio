package notification

import (
	_entities "potentivio-app/entities"
	_notifRepository "potentivio-app/repository/notification"
)

type NotifUseCase struct {
	notifRepository _notifRepository.NotifRepositoryInterface
}

func NewNotifUseCase(notifRepo _notifRepository.NotifRepositoryInterface) NotifUseCaseInterface {
	return &NotifUseCase{
		notifRepository: notifRepo,
	}
}

func (nuc *NotifUseCase) CreateNotif(notif _entities.Notification, idToken uint, idCafe uint) (_entities.Notification, error) {
	notif.IdArtist = idToken
	notif.IdCafe = idCafe
	createNotif, err := nuc.notifRepository.CreateNotif(notif)
	return createNotif, err
}

func (nuc *NotifUseCase) GetAllNotifByIdCafe(idToken uint) ([]_entities.Notification, uint, error) {
	notif, rows, err := nuc.notifRepository.GetAllNotifByIdCafe(idToken)
	return notif, rows, err
}

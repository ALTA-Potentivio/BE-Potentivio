package notification

import (
	"errors"
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
	getNotif, _, _ := nuc.notifRepository.GetAllNotifByIdArtist(idToken)

	for i := 0; i < len(getNotif); i++ {
		if getNotif[i].IdArtist == idToken && getNotif[i].IdCafe == idCafe {
			return notif, errors.New("already sent a notification")
		}
	}

	notif.IdArtist = idToken
	notif.IdCafe = idCafe
	createNotif, err := nuc.notifRepository.CreateNotif(notif)
	return createNotif, err
}

func (nuc *NotifUseCase) GetAllNotifByIdCafe(idToken uint) ([]_entities.Notification, uint, error) {
	notif, rows, err := nuc.notifRepository.GetAllNotifByIdCafe(idToken)
	return notif, rows, err
}

func (nuc *NotifUseCase) GetAllNotifByIdArtist(idToken uint) ([]_entities.Notification, uint, error) {
	notif, rows, err := nuc.notifRepository.GetAllNotifByIdCafe(idToken)
	return notif, rows, err
}

func (nuc *NotifUseCase) GetNotifById(id uint) (_entities.Notification, int, error) {
	notif, rows, err := nuc.notifRepository.GetNotifById(id)
	return notif, rows, err
}

func (nuc *NotifUseCase) DeleteNotif(idToken uint, id uint) (uint, error) {
	notif, _, _ := nuc.notifRepository.GetNotifById(id)
	if notif.IdCafe != idToken {
		return 1, errors.New("unauthorized")
	}
	rows, err := nuc.notifRepository.DeleteNotif(id)
	return rows, err
}

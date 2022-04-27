package notification

import (
	_entities "potentivio-app/entities"

	"gorm.io/gorm"
)

type NotifRepository struct {
	database *gorm.DB
}

func NewNotifRepository(db *gorm.DB) *NotifRepository {
	return &NotifRepository{
		database: db,
	}
}

func (nr *NotifRepository) CreateNotif(notif _entities.Notification) (_entities.Notification, error) {
	tx := nr.database.Save(&notif)
	if tx.Error != nil {
		return notif, tx.Error
	}
	return notif, nil
}

func (nr *NotifRepository) GetAllNotifByIdCafe(idToken uint) ([]_entities.Notification, uint, error) {
	var notif []_entities.Notification
	tx := nr.database.Order("created_at desc").Where("id_cafe = ?", idToken).Preload("Artist.Catagory").Preload("Artist.Genre").Find(&notif)
	if tx.Error != nil {
		return notif, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return notif, 0, nil
	}
	return notif, uint(tx.RowsAffected), nil
}

func (nr *NotifRepository) GetAllNotifByIdArtist(idToken uint) ([]_entities.Notification, uint, error) {
	var notif []_entities.Notification
	tx := nr.database.Order("created_at desc").Where("id_artist = ?", idToken).Preload("Artist.Catagory").Preload("Artist.Genre").Find(&notif)
	if tx.Error != nil {
		return notif, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return notif, 0, nil
	}
	return notif, uint(tx.RowsAffected), nil
}

func (nr *NotifRepository) GetNotifById(id uint) (_entities.Notification, int, error) {
	var notif _entities.Notification
	tx := nr.database.Find(&notif, id)
	if tx.Error != nil {
		return notif, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return notif, 0, nil
	}
	return notif, int(tx.RowsAffected), nil
}

func (nr *NotifRepository) DeleteNotif(id uint) (uint, error) {
	var notif _entities.Notification
	tx := nr.database.Delete(&notif, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, nil
	}
	return uint(tx.RowsAffected), nil
}

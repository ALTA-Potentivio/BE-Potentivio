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
	tx := nr.database.Order("created_at desc").Where("id_cafe = ?", idToken).Preload("Artist.Catagory.Genre").Find(&notif)
	if tx.Error != nil {
		return notif, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return notif, 0, nil
	}
	return notif, uint(tx.RowsAffected), nil
}

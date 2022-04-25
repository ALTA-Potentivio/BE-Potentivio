package videoArtist

import (
	"errors"
	"fmt"
	_entities "potentivio-app/entities"

	"gorm.io/gorm"
)

type VideoRepository struct {
	database *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepository {
	return &VideoRepository{
		database: db,
	}
}

func (cr *VideoRepository) PostVideo(video _entities.VideoArtist) error {
	tx := cr.database.Save(&video)
	if tx.Error != nil {
		return tx.Error
	}
	return tx.Error
}

func (cr *VideoRepository) DeleteVideo(id int, idToken int) error {

	var video _entities.VideoArtist
	txFind := cr.database.Find(&video, id)
	if txFind.RowsAffected == 0 {
		return fmt.Errorf("data not found")
	}

	if video.IdArtist != uint(idToken) {
		return errors.New("unautorized")
	}

	tx := cr.database.Where("id = ?", id).Delete(&video)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

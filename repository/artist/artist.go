package artist

import (
	_entities "potentivio-app/entities"

	"gorm.io/gorm"
)

type ArtistRepository struct {
	database *gorm.DB
}

func NewArtistRepository(db *gorm.DB) *ArtistRepository {
	return &ArtistRepository{
		database: db,
	}
}

func (ar *ArtistRepository) CreateArtist(artist _entities.Artist) (_entities.Artist, error) {
	tx := ar.database.Save(&artist)
	if tx.Error != nil {
		return artist, tx.Error
	}
	return artist, nil
}

func (ar *ArtistRepository) GetArtistById(id uint) (_entities.Artist, []_entities.Hire, []_entities.Hire, int, error) {
	var artists _entities.Artist
	var hireNotAvailable []_entities.Hire
	var hireHistory []_entities.Hire

	txNotAvailable := ar.database.Where("StatusCafe = ?", "accepted").Find(&hireNotAvailable)
	if txNotAvailable.Error != nil {
		return artists, hireNotAvailable, hireNotAvailable, 0, txNotAvailable.Error
	}

	txHireHistory := ar.database.Preload("Cafe").Where("StatusCafe = ?", "done").Find(&hireHistory)
	if txHireHistory.Error != nil {
		return artists, hireNotAvailable, hireHistory, 0, txHireHistory.Error
	}

	tx := ar.database.Preload("VideoArtist").Find(&artists, id)
	if tx.Error != nil {
		return artists, hireNotAvailable, hireHistory, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return artists, hireNotAvailable, hireHistory, 0, nil
	}
	return artists, hireNotAvailable, hireHistory, int(tx.RowsAffected), nil
}

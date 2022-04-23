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

package genre

import (
	_entities "potentivio-app/entities"

	"gorm.io/gorm"
)

type GenreRepository struct {
	DB *gorm.DB
}

func NewGenreRepository(db *gorm.DB) *GenreRepository {
	return &GenreRepository{
		DB: db,
	}
}

func (gr *GenreRepository) CreateGenre(genre _entities.Genre) (_entities.Genre, error) {
	tx := gr.DB.Save(&genre)
	if tx.Error != nil {
		return genre, tx.Error
	}
	return genre, nil
}

func (gr *GenreRepository) GetAllGenre() ([]_entities.Genre, error) {
	var genre []_entities.Genre
	tx := gr.DB.Find(&genre)
	if tx.Error != nil {
		return genre, tx.Error
	}
	return genre, nil
}

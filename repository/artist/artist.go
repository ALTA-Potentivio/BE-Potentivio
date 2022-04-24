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

func (ar *ArtistRepository) GetAllArtist(filters_catagory_genre map[string]int, filters_price map[string]string, filters_address map[string]string) ([]_entities.Artist, uint, error) {
	var artists []_entities.Artist

	builder := ar.database.Order("Name ASC")

	for key, value := range filters_price {
		if value == "asc" {
			builder = ar.database.Order(key + " " + value)
		}
		if value == "desc" {
			builder = ar.database.Order(key + " " + value)
		}
	}

	for key, value := range filters_catagory_genre {
		builder.Where(key+" = ?", value)
	}

	for key, value := range filters_address {
		builder.Where(key+" LIKE ?", "%"+value+"%")
	}

	tx := builder.Preload("Catagory").Preload("Genre").Find(&artists)
	if tx.Error != nil {
		return artists, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return artists, 0, nil
	}
	return artists, uint(tx.RowsAffected), nil
}

func (ar *ArtistRepository) GetProfileArtist(idToken uint) (_entities.Artist, uint, error) {
	var artist _entities.Artist
	tx := ar.database.Where("ID = ?", idToken).Preload("VideoArtist").Find(&artist)
	if tx.Error != nil {
		return artist, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return artist, 0, nil
	}
	return artist, uint(tx.RowsAffected), nil
}

func (ar *ArtistRepository) GetArtistById(id uint) (_entities.Artist, []_entities.Hire, []_entities.Hire, int, error) {
	var artists _entities.Artist
	var hireNotAvailable []_entities.Hire
	var hireHistory []_entities.Hire

	ar.database.Where("StatusCafe = ?", "accepted").Find(&hireNotAvailable)
	ar.database.Preload("Cafe").Where("StatusCafe = ?", "done").Find(&hireHistory)

	tx := ar.database.Preload("VideoArtist").Find(&artists, id)
	if tx.Error != nil {
		return artists, hireNotAvailable, hireHistory, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return artists, hireNotAvailable, hireHistory, 0, nil
	}
	return artists, hireNotAvailable, hireHistory, int(tx.RowsAffected), nil
}

func (ar *ArtistRepository) UpdateArtist(updateArtist _entities.Artist) (_entities.Artist, uint, error) {
	tx := ar.database.Save(&updateArtist)
	if tx.Error != nil {
		return updateArtist, 0, tx.Error
	}
	return updateArtist, uint(tx.RowsAffected), nil
}

func (ar *ArtistRepository) DeleteArtist(id uint) (uint, error) {
	var artist _entities.Artist
	tx := ar.database.Delete(&artist, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, tx.Error
	}
	return uint(tx.RowsAffected), nil
}

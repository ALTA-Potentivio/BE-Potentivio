package hire

import (
	"gorm.io/gorm"
	"potentivio-app/entities"
)

type HireRepository struct {
	database *gorm.DB
}

func NewHireRepository(db *gorm.DB) *HireRepository {
	return &HireRepository{
		database: db,
	}
}

func (hr *HireRepository) CheckHire(hire entities.Hire) error {
	tx := hr.database.Where("id_artist = ? and date = ? ", hire.IdArtist, hire.Date).First(&hire)

	return tx.Error
}

func (hr *HireRepository) CreateHire(hire entities.Hire) error {
	tx := hr.database.Create(&hire)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (hr *HireRepository) GetHireById(id int) (entities.Hire, error) {
	var hire entities.Hire
	tx := hr.database.Where("id = ?", id).First(&hire)
	if tx.Error != nil {
		return hire, tx.Error
	}
	return hire, nil
}

func (hr *HireRepository) GetHireByIdArtist(IdArtist int) ([]entities.Hire, error) {
	var hire []entities.Hire
	tx := hr.database.Where("id_artist = ?", IdArtist).Preload("Cafe").Find(&hire)
	if tx.Error != nil {
		return hire, tx.Error
	}
	return hire, nil
}

func (hr *HireRepository) GetHireByIdCafe(IdCafe int) ([]entities.Hire, error) {
	var hire []entities.Hire
	tx := hr.database.Where("id_cafe = ?", IdCafe).Preload("Artist").Find(&hire)
	if tx.Error != nil {
		return hire, tx.Error
	}
	return hire, nil
}

func (hr *HireRepository) AcceptHire(hire entities.Hire) error {
	tx := hr.database.Where("id = ?", hire.ID).Updates(&hire)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (hr *HireRepository) UpdateHire(id int, hire entities.Hire) error {
	tx := hr.database.Where("id = ?", id).Updates(&hire)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (hr *HireRepository) DeleteHire(hire entities.Hire) error {
	tx := hr.database.Where("id = ?", hire.ID).Delete(&hire)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (hr *HireRepository) Rating(rating entities.Rating) error {

	tx := hr.database.Create(&rating)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
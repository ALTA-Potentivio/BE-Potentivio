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
	tx := hr.database.Where("id_artis = ? and date = ? ", hire.IdArtist, hire.Date).First(&hire)

	return tx.Error
}

func (hr *HireRepository) CreateHire(hire entities.Hire) error {
	tx := hr.database.Create(&hire)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (hr *HireRepository) AcceptHire(hire entities.Hire) error {
	tx := hr.database.Where("id = ?", hire.ID).Updates(&hire)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

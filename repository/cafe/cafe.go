package cafe

import (
	_entities "potentivio-app/entities"

	"gorm.io/gorm"
)

type CafeRepository struct {
	database *gorm.DB
}

func NewCafeRepository(db *gorm.DB) *CafeRepository {
	return &CafeRepository{
		database: db,
	}
}

func (cr *CafeRepository) GetCafeById(id int) (_entities.Cafe, int, error) {
	var cafe _entities.Cafe

	tx := cr.database.Find(&cafe, id)
	if tx.Error != nil {
		return cafe, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return cafe, 0, nil
	}

	return cafe, int(tx.RowsAffected), nil
}

package cafe

import (
	"fmt"
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
func (cr *CafeRepository) PostCafe(cafe _entities.Cafe) error {
	tx := cr.database.Save(&cafe)
	if tx.Error != nil {
		return tx.Error
	}
	return tx.Error
}
func (cr *CafeRepository) GetAllCafe() ([]_entities.Cafe, error) {
	var cafes []_entities.Cafe

	tx := cr.database.Find(&cafes)
	if tx.Error != nil {
		return cafes, tx.Error
	}

	return cafes, nil
}

func (cr *CafeRepository) DeleteCafe(id int) error {
	var cafe _entities.Cafe
	txFind := cr.database.Find(&cafe, id)
	if txFind.RowsAffected == 0 {
		return fmt.Errorf("data not found")
	}
	tx := cr.database.Where("id = ?", id).Delete(&cafe)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func (cr *CafeRepository) UpdateCafe(updateCafe _entities.Cafe, idToken int) (uint, error) {
	tx := cr.database.Where("ID = ?", idToken).Updates(&updateCafe)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return uint(tx.RowsAffected), nil
}

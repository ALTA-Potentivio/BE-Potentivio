package imageCafe

import (
	"errors"
	_entities "potentivio-app/entities"

	"gorm.io/gorm"
)

type ImageCafeRepository struct {
	database *gorm.DB
}

func NewImageCafeRepository(db *gorm.DB) *ImageCafeRepository {
	return &ImageCafeRepository{
		database: db,
	}
}

func (icr *ImageCafeRepository) CreateImageCafe(imageCafe _entities.ImageCafe) (_entities.ImageCafe, error) {
	tx := icr.database.Save(&imageCafe)
	if tx.Error != nil {
		return imageCafe, tx.Error
	}
	return imageCafe, nil
}

func (icr *ImageCafeRepository) GetImageIDLast() (uint, error) {
	var imageCafe _entities.ImageCafe
	icr.database.Last(&imageCafe)

	var id_image_cafe uint
	if imageCafe.ID == 0 {
		id_image_cafe = 1
	}
	if imageCafe.ID != 0 {
		id_image_cafe = imageCafe.ID + 1
	}
	return id_image_cafe, nil
}

func (icr *ImageCafeRepository) DeleteImageCafe(id uint, idToken uint) (uint, error) {
	var imageCafe _entities.ImageCafe

	txFind := icr.database.Find(&imageCafe, id)
	if txFind.Error != nil {
		return 0, txFind.Error
	}
	if txFind.RowsAffected == 0 {
		return 0, txFind.Error
	}
	if idToken != imageCafe.IdCafe {
		return 0, errors.New("unauthorized")
	}

	tx := icr.database.Delete(&imageCafe, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, tx.Error
	}
	return uint(tx.RowsAffected), nil
}

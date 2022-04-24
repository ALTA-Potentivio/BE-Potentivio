package imageCafe

import (
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

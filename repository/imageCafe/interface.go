package imageCafe

import (
	_entities "potentivio-app/entities"
)

type ImageCafeRepositoryInterface interface {
	CreateImageCafe(imageCafe _entities.ImageCafe) (_entities.ImageCafe, error)
	GetImageIDLast() (uint, error)
}

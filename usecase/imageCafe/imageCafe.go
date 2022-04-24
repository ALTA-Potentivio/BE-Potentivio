package imageCafe

import (
	"errors"
	_entities "potentivio-app/entities"
	_imageCafeRepository "potentivio-app/repository/imageCafe"
)

type ImageCafeUseCase struct {
	imageCafeRepository _imageCafeRepository.ImageCafeRepositoryInterface
}

func NewImageCafeUseCase(imageCafeRepo _imageCafeRepository.ImageCafeRepositoryInterface) ImageCafeUseCaseInterface {
	return &ImageCafeUseCase{
		imageCafeRepository: imageCafeRepo,
	}
}

func (icuc *ImageCafeUseCase) CreateImageCafe(imageCafe _entities.ImageCafe) (_entities.ImageCafe, error) {

	if imageCafe.ImageUrl == "" {
		return imageCafe, errors.New("image_url is required")
	}

	createImageCafe, err := icuc.imageCafeRepository.CreateImageCafe(imageCafe)
	return createImageCafe, err
}

func (icuc *ImageCafeUseCase) GetImageIDLast() (uint, error) {
	idname, err := icuc.imageCafeRepository.GetImageIDLast()
	return idname, err
}

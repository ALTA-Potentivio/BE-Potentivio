package imageCafe

import (
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
	createImageCafe, err := icuc.imageCafeRepository.CreateImageCafe(imageCafe)
	return createImageCafe, err
}

func (icuc *ImageCafeUseCase) GetImageIDLast() (uint, error) {
	idname, err := icuc.imageCafeRepository.GetImageIDLast()
	return idname, err
}

func (icuc *ImageCafeUseCase) DeleteImageCafe(id uint, idToken uint) (uint, error) {
	rows, err := icuc.imageCafeRepository.DeleteImageCafe(id, idToken)
	return rows, err
}

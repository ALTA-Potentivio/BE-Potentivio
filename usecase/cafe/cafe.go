package cafe

import (
	_entities "potentivio-app/entities"
	_cafeRepository "potentivio-app/repository/cafe"
)

type CafeUseCase struct {
	cafeRepository _cafeRepository.CafeRepositoryInterface
}

func NewCafeUseCase(cafeRepo _cafeRepository.CafeRepositoryInterface) CafeUseCaseInterface {
	return &CafeUseCase{
		cafeRepository: cafeRepo,
	}
}
func (cuc *CafeUseCase) GetCafeById(id int) (_entities.Cafe, int, error) {
	cafe, rows, err := cuc.cafeRepository.GetCafeById(id)

	var GetCafe _entities.GetCafe
	GetCafe.AccountNumber = cafe.AccountNumber
	GetCafe.Address = cafe.Address
	GetCafe.Avatar = cafe.Avatar
	GetCafe.Description = cafe.Description
	GetCafe.Email = cafe.Email
	GetCafe.ID = cafe.ID
	GetCafe.ImageCafe = cafe.ImageCafe
	GetCafe.Latitude = cafe.Latitude
	GetCafe.Longitude = cafe.Longitude
	GetCafe.Name = cafe.Name
	GetCafe.OpeningHours = cafe.OpeningHours
	GetCafe.Owner = cafe.Owner
	GetCafe.PhoneNumber = cafe.PhoneNumber

	return cafe, rows, err
}

package cafe

import (
	"errors"
	"potentivio-app/delivery/helper"
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
func (cuc *CafeUseCase) GetCafeById(id int) (_entities.GetCafe, int, error) {
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

	return GetCafe, rows, err
}

func (cuc *CafeUseCase) PostCafe(cafe _entities.Cafe) error {
	password, errHash := helper.HashPassword(cafe.Password)
	if errHash != nil {
		return errors.New("error hashing password")
	}
	cafe.Password = password
	if cafe.Name == "" {
		return errors.New("name is required")
	}
	if cafe.Email == "" {
		return errors.New("email is required")
	}
	if cafe.Password == "" {
		return errors.New("password is required")
	}
	if cafe.Address == "" {
		return errors.New("address is required")
	}
	error := cuc.cafeRepository.PostCafe(cafe)
	return error
}

func (cuc *CafeUseCase) GetAllCafe(filters map[string]string) ([]_entities.GetAllCafe, error) {
	cafe, error := cuc.cafeRepository.GetAllCafe(filters)
	var GetCafe _entities.GetAllCafe
	var GetCafes []_entities.GetAllCafe

	for i := 0; i < len(cafe); i++ {

		GetCafe.Address = cafe[i].Address
		GetCafe.Avatar = cafe[i].Avatar
		GetCafe.Description = cafe[i].Description
		GetCafe.ID = cafe[i].ID
		GetCafe.Name = cafe[i].Name
		GetCafes = append(GetCafes, GetCafe)

	}
	return GetCafes, error
}
func (cuc *CafeUseCase) DeleteCafe(id int) error {
	error := cuc.cafeRepository.DeleteCafe(id)
	return error
}

func (cuc *CafeUseCase) UpdateCafe(updateCafe _entities.Cafe) error {
	updateErr := cuc.cafeRepository.UpdateCafe(updateCafe)
	return updateErr
}

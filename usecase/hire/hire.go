package hire

import (
	"errors"
	"fmt"
	"potentivio-app/entities"
	"potentivio-app/repository/artist"
	"potentivio-app/repository/cafe"
	"potentivio-app/repository/hire"
)

type HireUseCase struct {
	HireRepository   hire.HireRepositoryInterface
	ArtistRepository artist.ArtistRepositoryInterface
	CafeRepository   cafe.CafeRepositoryInterface
}

func NewHireUseCase(hireRepo hire.HireRepositoryInterface, artistRepo artist.ArtistRepositoryInterface, cafeRepo cafe.CafeRepositoryInterface) HireUseCaseInterface {
	return &HireUseCase{
		HireRepository:   hireRepo,
		ArtistRepository: artistRepo,
		CafeRepository:   cafeRepo,
	}
}

func (huc *HireUseCase) CreateHire(hire entities.Hire) error {

	artistData, _, _, _, err := huc.ArtistRepository.GetArtistById(hire.IdArtist)
	if err != nil {
		return errors.New("Artist not found")

	}

	cafeData, _, _ := huc.CafeRepository.GetCafeById(int(hire.IdCafe))
	err = huc.HireRepository.CheckHire(hire)
	fmt.Println(err)
	if err == nil {
		return errors.New("Artis not Available")
	}

	hire.Price = artistData.Price
	hire.AccountNumberArtist = artistData.AccountNumber
	hire.AccountNumberCafe = cafeData.AccountNumber
	
	err = huc.HireRepository.CreateHire(hire)

	return err

}

func (huc *HireUseCase) GetHireByIdArtis(IdArtist int) ([]entities.Hire, error) {
	hires, err := huc.HireRepository.GetHireByIdArtis(IdArtist)
	return hires, err
}

package artist

import (
	"errors"
	"potentivio-app/delivery/helper"
	_entities "potentivio-app/entities"
	_artistRepository "potentivio-app/repository/artist"
)

type ArtistUseCase struct {
	artistRepository _artistRepository.ArtistRepositoryInterface
}

func NewArtistUseCase(artistRepo _artistRepository.ArtistRepositoryInterface) ArtistUseCaseInterface {
	return &ArtistUseCase{
		artistRepository: artistRepo,
	}
}

func (auc *ArtistUseCase) CreateArtist(artist _entities.Artist) (_entities.Artist, error) {
	password, _ := helper.HashPassword(artist.Password)
	artist.Password = password
	//validasi saat registrasi
	if artist.Name == "" {
		return artist, errors.New("name is required")
	}
	if artist.Email == "" {
		return artist, errors.New("email is required")
	}
	if artist.Password == "" {
		return artist, errors.New("password is required")
	}
	if artist.Address == "" {
		return artist, errors.New("address is required")
	}

	createArtist, err := auc.artistRepository.CreateArtist(artist)
	return createArtist, err
}

func (auc *ArtistUseCase) GetAllArtist(filters_catagory_genre map[string]int, filters_price map[string]string, filters_address map[string]string) ([]_entities.Artist, uint, error) {
	artists, rows, err := auc.artistRepository.GetAllArtist(filters_catagory_genre, filters_price, filters_address)
	return artists, rows, err
}

func (auc *ArtistUseCase) GetArtistById(id uint) (_entities.Artist, []_entities.Hire, []_entities.Hire, int, error) {
	artist, hireNotAvailable, hireHistory, rows, err := auc.artistRepository.GetArtistById(id)
	return artist, hireNotAvailable, hireHistory, rows, err
}

func (auc *ArtistUseCase) GetProfileArtist(idToken uint) (_entities.Artist, uint, error) {
	artist, rows, err := auc.artistRepository.GetProfileArtist(idToken)
	return artist, rows, err
}

func (auc *ArtistUseCase) UpdateArtist(updateArtist _entities.Artist, idToken uint) (_entities.Artist, uint, error) {

	updateArtist, updateRows, updateErr := auc.artistRepository.UpdateArtist(updateArtist, idToken)
	return updateArtist, updateRows, updateErr
}

func (auc *ArtistUseCase) DeleteArtist(id uint) (uint, error) {
	rows, err := auc.artistRepository.DeleteArtist(id)
	return rows, err
}

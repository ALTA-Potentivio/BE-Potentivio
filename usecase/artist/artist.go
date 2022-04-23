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

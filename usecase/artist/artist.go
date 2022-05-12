package artist

import (
	"errors"
	"potentivio-app/delivery/helper"
	_entities "potentivio-app/entities"
	_artistRepository "potentivio-app/repository/artist"
	_hireRepository "potentivio-app/repository/hire"
)

type ArtistUseCase struct {
	artistRepository _artistRepository.ArtistRepositoryInterface
	hireRepository   _hireRepository.HireRepositoryInterface
}

func NewArtistUseCase(artistRepo _artistRepository.ArtistRepositoryInterface, hireRepo _hireRepository.HireRepositoryInterface) ArtistUseCaseInterface {
	return &ArtistUseCase{
		artistRepository: artistRepo,
		hireRepository:   hireRepo,
	}
}

func (auc *ArtistUseCase) CreateArtist(artist _entities.Artist) (_entities.Artist, error) {
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

	password, _ := helper.HashPassword(artist.Password)
	artist.Password = password

	createArtist, err := auc.artistRepository.CreateArtist(artist)
	return createArtist, err
}

func (auc *ArtistUseCase) GetAllArtist(filtersCatagoryGenre map[string]int, filtersPrice map[string]string, filtersAddress map[string]string) ([]_entities.Artist, uint, error) {
	artists, rows, err := auc.artistRepository.GetAllArtist(filtersCatagoryGenre, filtersPrice, filtersAddress)
	return artists, rows, err
}

func (auc *ArtistUseCase) GetArtistById(id uint) (_entities.Artist, []_entities.Hire, int, error) {
	hires, _ := auc.hireRepository.GetHireByIdArtist(int(id))

	artist, rows, err := auc.artistRepository.GetArtistById(id)
	return artist, hires, rows, err
}

func (auc *ArtistUseCase) GetProfileArtist(idToken uint) (_entities.Artist, uint, error) {
	artist, rows, err := auc.artistRepository.GetProfileArtist(idToken)
	return artist, rows, err
}

func (auc *ArtistUseCase) UpdateArtist(updateArtist _entities.Artist, idToken uint) (_entities.Artist, uint, error) {
	updateArtist, updateRows, updateErr := auc.artistRepository.UpdateArtist(updateArtist, idToken)
	return updateArtist, updateRows, updateErr
}

func (auc *ArtistUseCase) DeleteArtist(idToken uint) (uint, error) {
	rows, err := auc.artistRepository.DeleteArtist(idToken)
	return rows, err
}

func (auc *ArtistUseCase) CountRating(idArtist uint) (uint, float32, error) {
	total, rating, err := auc.artistRepository.CountRating(idArtist)
	return total, rating, err
}

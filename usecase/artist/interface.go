package artist

import (
	_entities "potentivio-app/entities"
)

type ArtistUseCaseInterface interface {
	CreateArtist(artist _entities.Artist) (_entities.Artist, error)
	GetAllArtist(filters_catagory_genre map[string]int, filters_price map[string]string, filters_address map[string]string) ([]_entities.Artist, uint, error)
	GetProfileArtist(idToken uint) (_entities.Artist, uint, error)
	GetArtistById(id uint) (_entities.Artist, error)
	UpdateArtist(updateArtist _entities.Artist, idToken uint) (_entities.Artist, uint, error)
	DeleteArtist(id uint) (uint, error)
}

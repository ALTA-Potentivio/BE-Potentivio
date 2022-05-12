package artist

import (
	_entities "potentivio-app/entities"
)

type ArtistRepositoryInterface interface {
	CreateArtist(artist _entities.Artist) (_entities.Artist, error)
	GetAllArtist(filtersCatagoryGenre map[string]int, filtersPrice map[string]string, filtersAddress map[string]string) ([]_entities.Artist, uint, error)
	GetProfileArtist(idToken uint) (_entities.Artist, uint, error)
	GetArtistById(id uint) (_entities.Artist, int, error)
	UpdateArtist(updateArtist _entities.Artist, idToken uint) (_entities.Artist, uint, error)
	DeleteArtist(idToken uint) (uint, error)
	GetArtistByIdForHire(id uint) (_entities.Artist, error)
	CountRating(idArtist uint) (uint, float32, error)
}

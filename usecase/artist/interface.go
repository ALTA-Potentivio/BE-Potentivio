package artist

import (
	_entities "potentivio-app/entities"
)

type ArtistUseCaseInterface interface {
	CreateArtist(artist _entities.Artist) (_entities.Artist, error)
	GetAllArtist() ([]_entities.Artist, uint, error)
	GetArtistById(id uint) (_entities.Artist, []_entities.Hire, []_entities.Hire, int, error)
}

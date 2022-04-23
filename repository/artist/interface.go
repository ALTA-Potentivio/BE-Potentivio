package artist

import (
	_entities "potentivio-app/entities"
)

type ArtistRepositoryInterface interface {
	CreateArtist(artist _entities.Artist) (_entities.Artist, error)
}

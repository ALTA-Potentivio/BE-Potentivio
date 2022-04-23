package artist

import (
	_entities "potentivio-app/entities"
)

type ArtistUseCaseInterface interface {
	CreateArtist(artist _entities.Artist) (_entities.Artist, error)
}

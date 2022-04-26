package genre

import (
	_entities "potentivio-app/entities"
)

type GenreRepositoryInterface interface {
	CreateGenre(genre _entities.Genre) (_entities.Genre, error)
	GetAllGenre() ([]_entities.Genre, error)
}

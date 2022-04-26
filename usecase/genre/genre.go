package genre

import (
	_entities "potentivio-app/entities"
	_genreRepository "potentivio-app/repository/genre"
)

type GenreUseCase struct {
	genreRepository _genreRepository.GenreRepositoryInterface
}

func NewGenreUseCase(genreRepo _genreRepository.GenreRepositoryInterface) GenreUseCaseInterface {
	return &GenreUseCase{
		genreRepository: genreRepo,
	}
}

func (guc *GenreUseCase) CreateGenre(genre _entities.Genre) (_entities.Genre, error) {
	createGenre, err := guc.genreRepository.CreateGenre(genre)
	return createGenre, err
}

func (guc *GenreUseCase) GetAllGenre() ([]_entities.Genre, error) {
	genre, err := guc.genreRepository.GetAllGenre()
	return genre, err
}

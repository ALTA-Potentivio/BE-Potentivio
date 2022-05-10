package genre

import (
	"fmt"
	_entities "potentivio-app/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGenre(t *testing.T) {
	t.Run("TestCreateGenreSuccess", func(t *testing.T) {
		genreUseCase := NewGenreUseCase(mockGenreRepository{})
		genre, err := genreUseCase.CreateGenre(_entities.Genre{NameGenre: "pop"})
		assert.Nil(t, err)
		assert.Equal(t, "pop", genre.NameGenre)
	})

	t.Run("TestCreateGenreError", func(t *testing.T) {
		genreUseCase := NewGenreUseCase(mockGenreRepositoryError{})
		genre, err := genreUseCase.CreateGenre(_entities.Genre{NameGenre: "pop"})
		assert.NotNil(t, err)
		assert.Equal(t, "", genre.NameGenre)
	})
}

func TestGetAllGenre(t *testing.T) {
	t.Run("TestGetAllGenreSuccess", func(t *testing.T) {
		genreUseCase := NewGenreUseCase(mockGenreRepository{})
		genre, err := genreUseCase.GetAllGenre()
		assert.Nil(t, err)
		assert.NotNil(t, genre)
	})

	t.Run("TestGetAllGenreError", func(t *testing.T) {
		genreUseCase := NewGenreUseCase(mockGenreRepositoryError{})
		genre, err := genreUseCase.GetAllGenre()
		assert.NotNil(t, err)
		assert.Nil(t, genre)
	})
}

// === mock success ===
type mockGenreRepository struct{}

func (m mockGenreRepository) GetAllGenre() ([]_entities.Genre, error) {
	return []_entities.Genre{
		{NameGenre: "pop"},
	}, nil
}

func (m mockGenreRepository) CreateGenre(genre _entities.Genre) (_entities.Genre, error) {
	return _entities.Genre{NameGenre: "pop"}, nil
}

// === mock error ===

type mockGenreRepositoryError struct{}

func (m mockGenreRepositoryError) GetAllGenre() ([]_entities.Genre, error) {
	return nil, fmt.Errorf("error get all genre")
}

func (m mockGenreRepositoryError) CreateGenre(genre _entities.Genre) (_entities.Genre, error) {
	return _entities.Genre{}, fmt.Errorf("error create genre")
}

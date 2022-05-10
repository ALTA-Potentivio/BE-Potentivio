package cafe

import (
	"errors"
	"fmt"
	_entities "potentivio-app/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostCafe(t *testing.T) {
	t.Run("TestPostCafeSuccess", func(t *testing.T) {
		cafeUseCase := NewCafeUseCase(mockCafeRepository{})
		err := cafeUseCase.PostCafe(_entities.Cafe{Name: "cafe semesta", Owner: "Satria", Email: "satria@mail.com", Address: "jakarta", Password: "pas123"})
		assert.Nil(t, err)
	})

	t.Run("TestPostCafeError", func(t *testing.T) {
		cafeUseCase := NewCafeUseCase(mockCafeRepositoryError{})
		err := cafeUseCase.PostCafe(_entities.Cafe{})
		if assert.Error(t, err) {
			assert.Equal(t, errors.New("name is required"), err)
		}
		assert.ErrorContains(t, err, "error hashing password")
		assert.NotNil(t, err)
	})
}

func TestGetAllCafe(t *testing.T) {
	t.Run("TestGetAllCafeSuccess", func(t *testing.T) {
		cafeUseCase := NewCafeUseCase(mockCafeRepository{})
		cafes, err := cafeUseCase.GetAllCafe(map[string]string{"Name": "Satria"})
		assert.Equal(t, "cafe semesta", cafes[0].Name)
		assert.Nil(t, err)
	})

	t.Run("TestGetAllCafeError", func(t *testing.T) {
		cafeUseCase := NewCafeUseCase(mockCafeRepositoryError{})
		cafes, err := cafeUseCase.GetAllCafe(map[string]string{"Name": "Satria"})
		assert.Nil(t, cafes)
		assert.NotNil(t, err)
	})
}

func TestGetCafeById(t *testing.T) {
	t.Run("TestGetCafeByIdSuccess", func(t *testing.T) {
		cafeUseCase := NewCafeUseCase(mockCafeRepository{})
		cafe, rows, err := cafeUseCase.GetCafeById(1)
		assert.Equal(t, 1, rows)
		assert.Equal(t, "Satria", cafe.Owner)
		assert.Nil(t, err)
	})

	t.Run("TestGetCafeByIdError", func(t *testing.T) {
		cafeUseCase := NewCafeUseCase(mockCafeRepositoryError{})
		cafe, rows, err := cafeUseCase.GetCafeById(1)
		assert.Equal(t, _entities.GetCafe{}, cafe)
		assert.Equal(t, 0, rows)
		assert.NotNil(t, err)
	})
}

func TestDeleteCafe(t *testing.T) {
	t.Run("TestDeleteCafeSuccess", func(t *testing.T) {
		cafeUseCase := NewCafeUseCase(mockCafeRepository{})
		err := cafeUseCase.DeleteCafe(1)
		assert.Nil(t, err)
	})

	t.Run("TestDeleteCafeError", func(t *testing.T) {
		cafeUseCase := NewCafeUseCase(mockCafeRepositoryError{})
		err := cafeUseCase.DeleteCafe(1)
		assert.NotNil(t, err)
	})
}

func TestUpdateCafe(t *testing.T) {
	t.Run("TestUpdateSuccess", func(t *testing.T) {
		cafeUseCase := NewCafeUseCase(mockCafeRepository{})
		err := cafeUseCase.UpdateCafe(_entities.Cafe{Name: "cafe putra", Owner: "Putra", Email: "puta@mail.com"})
		assert.Nil(t, err)
	})

	t.Run("TestGetCafeByIdError", func(t *testing.T) {
		cafeUseCase := NewCafeUseCase(mockCafeRepositoryError{})
		err := cafeUseCase.UpdateCafe(_entities.Cafe{Name: "cafe putra", Owner: "Putra", Email: "puta@mail.com"})
		assert.NotNil(t, err)
	})
}

// === mock success ===
type mockCafeRepository struct{}

func (m mockCafeRepository) PostCafe(cafe _entities.Cafe) error {
	return nil
}

func (m mockCafeRepository) GetAllCafe(filters map[string]string) ([]_entities.Cafe, error) {
	return []_entities.Cafe{
		{Name: "cafe semesta", Owner: "Satria", Email: "satria@mail.com"},
	}, nil
}

func (m mockCafeRepository) GetCafeById(id int) (_entities.Cafe, int, error) {
	return _entities.Cafe{
		Name: "cafe semesta", Owner: "Satria", Email: "satria@mail.com",
	}, 1, nil
}

func (m mockCafeRepository) DeleteCafe(id int) error {
	return nil
}

func (m mockCafeRepository) UpdateCafe(updateCafe _entities.Cafe) error {
	return nil
}

func (m mockCafeRepository) GetCafeByIdForHire(id uint) (_entities.Cafe, error) {
	return _entities.Cafe{
		Name: "cafe semesta", Owner: "Satria", Email: "satria@mail.com",
	}, nil
}

// === mock error ===

type mockCafeRepositoryError struct{}

func (m mockCafeRepositoryError) PostCafe(cafe _entities.Cafe) error {
	return fmt.Errorf("failed")
}

func (m mockCafeRepositoryError) GetAllCafe(filters map[string]string) ([]_entities.Cafe, error) {
	return []_entities.Cafe{}, fmt.Errorf("failed")
}

func (m mockCafeRepositoryError) GetCafeById(id int) (_entities.Cafe, int, error) {
	return _entities.Cafe{}, 0, fmt.Errorf("failed")
}

func (m mockCafeRepositoryError) DeleteCafe(id int) error {
	return fmt.Errorf("failed")
}

func (m mockCafeRepositoryError) UpdateCafe(updateCafe _entities.Cafe) error {
	return fmt.Errorf("failed")
}
func (m mockCafeRepositoryError) GetCafeByIdForHire(id uint) (_entities.Cafe, error) {
	return _entities.Cafe{}, fmt.Errorf("failed")
}

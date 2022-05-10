package catagory

import (
	"fmt"
	_entities "potentivio-app/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCatagory(t *testing.T) {
	t.Run("TestCreateCatagorySuccess", func(t *testing.T) {
		catagoryUseCase := NewCatagoryUseCase(mockCategoryRepository{})
		catagory, err := catagoryUseCase.CreateCatagory(_entities.Catagory{NameCatagory: "solo"})
		assert.Nil(t, err)
		assert.Equal(t, "solo", catagory.NameCatagory)
	})

	t.Run("TestCreateCatgoryError", func(t *testing.T) {
		catagoryUseCase := NewCatagoryUseCase(mockCategoryRepositoryError{})
		catagory, err := catagoryUseCase.CreateCatagory(_entities.Catagory{NameCatagory: "solo"})
		assert.NotNil(t, err)
		assert.Equal(t, "", catagory.NameCatagory)
	})
}

func TestGetAllCatagory(t *testing.T) {
	t.Run("TestGetAllCatagorySuccess", func(t *testing.T) {
		catagoryUseCase := NewCatagoryUseCase(mockCategoryRepository{})
		catagory, err := catagoryUseCase.GetAllCatagory()
		assert.Nil(t, err)
		assert.NotNil(t, catagory)
	})

	t.Run("TestGetAllCatagoryError", func(t *testing.T) {
		catagoryUseCase := NewCatagoryUseCase(mockCategoryRepositoryError{})
		catagory, err := catagoryUseCase.GetAllCatagory()
		assert.NotNil(t, err)
		assert.Nil(t, catagory)
	})
}

// === mock success ===
type mockCategoryRepository struct{}

func (m mockCategoryRepository) GetAllCatagory() ([]_entities.Catagory, error) {
	return []_entities.Catagory{
		{NameCatagory: "solo"},
	}, nil
}

func (m mockCategoryRepository) CreateCatagory(catagory _entities.Catagory) (_entities.Catagory, error) {
	return _entities.Catagory{NameCatagory: "solo"}, nil
}

// === mock error ===

type mockCategoryRepositoryError struct{}

func (m mockCategoryRepositoryError) GetAllCatagory() ([]_entities.Catagory, error) {
	return nil, fmt.Errorf("error get all catagory")
}

func (m mockCategoryRepositoryError) CreateCatagory(catagory _entities.Catagory) (_entities.Catagory, error) {
	return _entities.Catagory{}, fmt.Errorf("error create catagory")
}

package imageCafe

import (
	"fmt"
	_entities "potentivio-app/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateImageCafe(t *testing.T) {
	t.Run("TestCreateImageCafeSuccess", func(t *testing.T) {
		imageCafeUseCase := NewImageCafeUseCase(mockImageCafeRepository{})
		createImageCafe, err := imageCafeUseCase.CreateImageCafe(_entities.ImageCafe{IdCafe: 1, ImageUrl: "cafe.png"})
		assert.Nil(t, err)
		assert.Equal(t, "cafe.png", createImageCafe.ImageUrl)
		assert.Equal(t, uint(1), createImageCafe.IdCafe)
	})

	t.Run("TestCreateImageCafeError", func(t *testing.T) {
		imageCafeUseCase := NewImageCafeUseCase(mockImageCafeRepositoryError{})
		createImageCafe, err := imageCafeUseCase.CreateImageCafe(_entities.ImageCafe{IdCafe: 1, ImageUrl: "cafe.png"})
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), createImageCafe.IdCafe)
		assert.Equal(t, "", createImageCafe.ImageUrl)
	})
}

func TestDeleteImageCafe(t *testing.T) {
	t.Run("TestDeleteImageCafeSuccess", func(t *testing.T) {
		imageCafeUseCase := NewImageCafeUseCase(mockImageCafeRepository{})
		rows, err := imageCafeUseCase.DeleteImageCafe(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, uint(1), rows)
	})

	t.Run("TestDeleteImageCafeError", func(t *testing.T) {
		imageCafeUseCase := NewImageCafeUseCase(mockImageCafeRepositoryError{})
		rows, err := imageCafeUseCase.DeleteImageCafe(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), rows)
	})
}

func TestGetImageIDLast(t *testing.T) {
	t.Run("TestGetImageIDLastSuccess", func(t *testing.T) {
		imageCafeUseCase := NewImageCafeUseCase(mockImageCafeRepository{})
		idName, err := imageCafeUseCase.GetImageIDLast()
		assert.Nil(t, err)
		assert.Equal(t, uint(1), idName)
	})

	t.Run("TestGetImageIDLastError", func(t *testing.T) {
		imageCafeUseCase := NewImageCafeUseCase(mockImageCafeRepositoryError{})
		idName, err := imageCafeUseCase.GetImageIDLast()
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), idName)
	})
}

// === mock success ===
type mockImageCafeRepository struct{}

func (m mockImageCafeRepository) CreateImageCafe(imageCafe _entities.ImageCafe) (_entities.ImageCafe, error) {
	return imageCafe, nil
}

func (m mockImageCafeRepository) DeleteImageCafe(id uint, idToken uint) (uint, error) {
	return 1, nil
}

func (m mockImageCafeRepository) GetImageIDLast() (uint, error) {
	return 1, nil
}

// === mock error ===
type mockImageCafeRepositoryError struct{}

func (m mockImageCafeRepositoryError) CreateImageCafe(imageCafe _entities.ImageCafe) (_entities.ImageCafe, error) {
	return _entities.ImageCafe{}, fmt.Errorf("error create image cafe")
}

func (m mockImageCafeRepositoryError) DeleteImageCafe(id uint, idToken uint) (uint, error) {
	return 0, fmt.Errorf("error delete image cafe")
}

func (m mockImageCafeRepositoryError) GetImageIDLast() (uint, error) {
	return 0, fmt.Errorf("error get id last image cafe")
}

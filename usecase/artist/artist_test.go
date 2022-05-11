package artist

import (
	"fmt"
	_entities "potentivio-app/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateArtist(t *testing.T) {
	t.Run("TestCreateArtistSuccess", func(t *testing.T) {
		artistUseCase := NewArtistUseCase(mockArtistRepository{}, mockHireRepository{})
		createArtist, err := artistUseCase.CreateArtist(_entities.Artist{Name: "usamah", Email: "usamah@gmail.com", Password: "usamah", Address: "Kota Bogor"})
		assert.Nil(t, err)
		assert.Equal(t, "usamah", createArtist.Name)
	})

	t.Run("TestCreateArtistSuccessValidation-1", func(t *testing.T) {
		artistUseCase := NewArtistUseCase(mockArtistRepository{}, mockHireRepository{})
		createArtist, err := artistUseCase.CreateArtist(_entities.Artist{Email: "usamah@gmail.com", Password: "usamah", Address: "Kota Bogor"})
		assert.NotNil(t, err)
		assert.Equal(t, "", createArtist.Name)
	})

	t.Run("TestCreateArtistSuccessValidation-2", func(t *testing.T) {
		artistUseCase := NewArtistUseCase(mockArtistRepository{}, mockHireRepository{})
		createArtist, err := artistUseCase.CreateArtist(_entities.Artist{Name: "usamah", Password: "usamah", Address: "Kota Bogor"})
		assert.NotNil(t, err)
		assert.Equal(t, "usamah", createArtist.Name)
	})

	t.Run("TestCreateArtistSuccessValidation-3", func(t *testing.T) {
		artistUseCase := NewArtistUseCase(mockArtistRepository{}, mockHireRepository{})
		createArtist, err := artistUseCase.CreateArtist(_entities.Artist{Name: "usamah", Email: "usamah@gmail.com", Password: "usamah"})
		assert.NotNil(t, err)
		assert.Equal(t, "usamah", createArtist.Name)
	})

	t.Run("TestCreateArtistSuccessValidation-4", func(t *testing.T) {
		artistUseCase := NewArtistUseCase(mockArtistRepository{}, mockHireRepository{})
		createArtist, err := artistUseCase.CreateArtist(_entities.Artist{Name: "usamah", Email: "usamah@gmail.com", Address: "Kota Bogor"})
		assert.NotNil(t, err)
		assert.Equal(t, "usamah", createArtist.Name)
	})

	t.Run("TestCreateArtistError", func(t *testing.T) {
		artistUseCase := NewArtistUseCase(mockArtistRepositoryError{}, mockHireRepository{})
		createArtist, err := artistUseCase.CreateArtist(_entities.Artist{Name: "usamah", Email: "usamah@gmail.com", Password: "usamah", Address: "Kota Bogor"})
		assert.NotNil(t, err)
		assert.Equal(t, "", createArtist.Name)
	})
}

func TestGetAllArtist(t *testing.T) {
	t.Run("TestGetAllArtistSuccess", func(t *testing.T) {
		artistUseCase := NewArtistUseCase(mockArtistRepository{}, mockHireRepository{})
		artist, rows, err := artistUseCase.GetAllArtist(map[string]int{}, map[string]string{}, map[string]string{})
		assert.Nil(t, err)
		assert.NotNil(t, artist)
		assert.Equal(t, uint(1), rows)
	})

	t.Run("TestGetAllArtistError", func(t *testing.T) {
		artistUseCase := NewArtistUseCase(mockArtistRepositoryError{}, mockHireRepository{})
		artist, rows, err := artistUseCase.GetAllArtist(map[string]int{}, map[string]string{}, map[string]string{})
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), rows)
		assert.Nil(t, artist)
	})
}

func TestGetProfileArtist(t *testing.T) {
	t.Run("TestGetProfileArtistSuccess", func(t *testing.T) {
		artistUseCase := NewArtistUseCase(mockArtistRepository{}, mockHireRepository{})
		artist, rows, err := artistUseCase.GetProfileArtist(1)
		assert.Nil(t, err)
		assert.Equal(t, "usamah", artist.Name)
		assert.Equal(t, uint(1), rows)
	})

	t.Run("TestGetProfileArtistError", func(t *testing.T) {
		artistUseCase := NewArtistUseCase(mockArtistRepositoryError{}, mockHireRepository{})
		artist, rows, err := artistUseCase.GetProfileArtist(1)
		assert.NotNil(t, err)
		assert.Equal(t, "", artist.Name)
		assert.Equal(t, uint(0), rows)
	})
}

func TestGetArtistById(t *testing.T) {
	t.Run("TestGetArtistByIdSuccess", func(t *testing.T) {
		artistUseCase := NewArtistUseCase(mockArtistRepository{}, mockHireRepository{})
		artist, hire, rows, err := artistUseCase.GetArtistById(1)
		assert.Nil(t, err)
		assert.Equal(t, "usamah", artist.Name)
		assert.NotNil(t, hire)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestGetArtistByIdError", func(t *testing.T) {
		artistUseCase := NewArtistUseCase(mockArtistRepositoryError{}, mockHireRepository{})
		artist, hire, rows, err := artistUseCase.GetArtistById(1)
		assert.NotNil(t, err)
		assert.Equal(t, "", artist.Name)
		assert.NotNil(t, hire)
		assert.Equal(t, 0, rows)
	})
}

func TestUpdateArtist(t *testing.T) {
	t.Run("TestUpdateArtistSuccess", func(t *testing.T) {
		artistUseCase := NewArtistUseCase(mockArtistRepository{}, mockHireRepository{})
		artist, rows, err := artistUseCase.UpdateArtist(_entities.Artist{Name: "usamah", Email: "usamah@gmail.com", Password: "usamah", Address: "Kota Bogor"}, 1)
		assert.Nil(t, err)
		assert.Equal(t, uint(1), rows)
		assert.Equal(t, "usamah", artist.Name)
	})

	t.Run("TestUpdateArtistError", func(t *testing.T) {
		artistUseCase := NewArtistUseCase(mockArtistRepositoryError{}, mockHireRepository{})
		artist, rows, err := artistUseCase.UpdateArtist(_entities.Artist{Name: "abdurrahman"}, 1)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), rows)
		assert.Equal(t, "usamah", artist.Name)
	})
}

func TestDeleteArtist(t *testing.T) {
	t.Run("TestDeleteArtistSuccess", func(t *testing.T) {
		artistUseCase := NewArtistUseCase(mockArtistRepository{}, mockHireRepository{})
		rows, err := artistUseCase.DeleteArtist(1)
		assert.Nil(t, err)
		assert.Equal(t, uint(1), rows)
	})

	t.Run("TestDeleteArtistError", func(t *testing.T) {
		artistUseCase := NewArtistUseCase(mockArtistRepositoryError{}, mockHireRepository{})
		rows, err := artistUseCase.DeleteArtist(1)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), rows)
	})
}

func TestCountRating(t *testing.T) {
	t.Run("TestCountRatingSuccess", func(t *testing.T) {
		artistUseCase := NewArtistUseCase(mockArtistRepository{}, mockHireRepository{})
		total, rating, err := artistUseCase.CountRating(1)
		assert.Nil(t, err)
		assert.Equal(t, uint(5), total)
		assert.Equal(t, float32(5), rating)
	})

	t.Run("TestCountRatingError", func(t *testing.T) {
		artistUseCase := NewArtistUseCase(mockArtistRepositoryError{}, mockHireRepository{})
		total, rating, err := artistUseCase.CountRating(1)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), total)
		assert.Equal(t, float32(0), rating)
	})
}

// === mock success ===
type mockArtistRepository struct{}

func (m mockArtistRepository) CreateArtist(artist _entities.Artist) (_entities.Artist, error) {
	return artist, nil
}

func (m mockArtistRepository) GetAllArtist(filters_catagory_genre map[string]int, filters_price map[string]string, filters_address map[string]string) ([]_entities.Artist, uint, error) {
	artist := []_entities.Artist{{Name: "usamah", Email: "usamah@gmail.com", Password: "usamah", Address: "Kota Bogor"}}
	return artist, 1, nil
}

func (m mockArtistRepository) GetProfileArtist(idToken uint) (_entities.Artist, uint, error) {
	artist := _entities.Artist{Name: "usamah", Email: "usamah@gmail.com", Password: "usamah", Address: "Kota Bogor"}
	return artist, 1, nil
}

func (m mockArtistRepository) GetArtistById(id uint) (_entities.Artist, int, error) {
	artist := _entities.Artist{Name: "usamah", Email: "usamah@gmail.com", Password: "usamah", Address: "Kota Bogor"}
	return artist, 1, nil
}

func (m mockArtistRepository) UpdateArtist(updateArtist _entities.Artist, idToken uint) (_entities.Artist, uint, error) {
	return updateArtist, 1, nil
}

func (m mockArtistRepository) DeleteArtist(idToken uint) (uint, error) {
	return 1, nil
}

func (m mockArtistRepository) CountRating(idArtist uint) (uint, float32, error) {
	return 5, 5, nil
}

func (m mockArtistRepository) GetArtistByIdForHire(id uint) (_entities.Artist, error) {
	artist := _entities.Artist{Name: "usamah", Email: "usamah@gmail.com", Password: "usamah", Address: "Kota Bogor"}
	return artist, nil
}

// === mock success ===
type mockHireRepository struct{}

func (m mockHireRepository) AcceptHire(hire _entities.Hire) error {
	return nil
}

func (m mockHireRepository) CheckHire(hire _entities.Hire) _entities.Hire {
	return _entities.Hire{}
}

func (m mockHireRepository) CallBack(hire _entities.Hire) error {
	return nil
}

func (m mockHireRepository) CreateHire(hire _entities.Hire) error {
	return nil
}

func (m mockHireRepository) GetHireByIdArtist(IdArtist int) ([]_entities.Hire, error) {
	return []_entities.Hire{}, nil
}

func (m mockHireRepository) GetHireById(id int) (_entities.Hire, error) {
	return _entities.Hire{}, nil
}

func (m mockHireRepository) GetHireByIdCafe(IdCafe int) ([]_entities.Hire, error) {
	return []_entities.Hire{}, nil
}

func (m mockHireRepository) UpdateHire(id int, hire _entities.Hire) error {
	return nil
}

func (m mockHireRepository) DeleteHire(hire _entities.Hire) error {
	return nil
}

func (m mockHireRepository) Rating(rating _entities.Rating) error {
	return nil
}

// === mock error ===
type mockArtistRepositoryError struct{}

func (m mockArtistRepositoryError) CreateArtist(artist _entities.Artist) (_entities.Artist, error) {
	return _entities.Artist{}, fmt.Errorf("error create artist")
}

func (m mockArtistRepositoryError) GetAllArtist(filters_catagory_genre map[string]int, filters_price map[string]string, filters_address map[string]string) ([]_entities.Artist, uint, error) {
	return nil, 0, fmt.Errorf("error get all artist")
}

func (m mockArtistRepositoryError) GetProfileArtist(idToken uint) (_entities.Artist, uint, error) {
	return _entities.Artist{}, 0, fmt.Errorf("error get profile artist")
}

func (m mockArtistRepositoryError) GetArtistById(id uint) (_entities.Artist, int, error) {
	return _entities.Artist{}, 0, fmt.Errorf("error get artist by id")
}

func (m mockArtistRepositoryError) UpdateArtist(updateArtist _entities.Artist, idToken uint) (_entities.Artist, uint, error) {
	artist := _entities.Artist{Name: "usamah", Email: "usamah@gmail.com", Password: "usamah", Address: "Kota Bogor"}
	return artist, 0, fmt.Errorf("error update artist")
}

func (m mockArtistRepositoryError) DeleteArtist(idToken uint) (uint, error) {
	return 0, fmt.Errorf("error delete artist")
}

func (m mockArtistRepositoryError) CountRating(idArtist uint) (uint, float32, error) {
	return 0, 0, fmt.Errorf("error count rating")
}

func (m mockArtistRepositoryError) GetArtistByIdForHire(id uint) (_entities.Artist, error) {
	return _entities.Artist{}, fmt.Errorf("error get artist by id for hire")
}

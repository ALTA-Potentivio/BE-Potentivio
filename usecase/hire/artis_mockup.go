package hire

import (
	"github.com/stretchr/testify/mock"
	"potentivio-app/entities"
)

type MockupArtisRepository struct {
	mock.Mock
}

func (m *MockupArtisRepository) CreateArtist(artist entities.Artist) (entities.Artist, error) {
	res := m.Called(artist)

	return res.Get(0).(entities.Artist), res.Error(0)
}

func (m *MockupArtisRepository) GetAllArtist(filters_catagory_genre map[string]int, filters_price map[string]string, filters_address map[string]string) ([]entities.Artist, uint, error) {
	res := m.Called(filters_catagory_genre, filters_price, filters_address)

	return res.Get(0).([]entities.Artist), res.Get(1).(uint), res.Error(2)
}

func (m *MockupArtisRepository) GetProfileArtist(idToken uint) (entities.Artist, uint, error) {
	res := m.Called(idToken)

	return res.Get(0).(entities.Artist), res.Get(1).(uint), res.Error(2)
}

func (m *MockupArtisRepository) GetArtistById(id uint) (entities.Artist, int, error) {
	res := m.Called(id)

	return res.Get(0).(entities.Artist), res.Get(1).(int), res.Error(2)
}

func (m *MockupArtisRepository) UpdateArtist(updateArtist entities.Artist, idToken uint) (entities.Artist, uint, error) {
	res := m.Called(updateArtist, idToken)

	return res.Get(0).(entities.Artist), res.Get(1).(uint), res.Error(2)
}

func (m *MockupArtisRepository) DeleteArtist(idToken uint) (uint, error) {
	res := m.Called(idToken)

	return res.Get(0).(uint), res.Error(1)
}

func (m *MockupArtisRepository) GetArtistByIdForHire(id uint) (entities.Artist, error) {
	res := m.Called(id)

	return res.Get(0).(entities.Artist), res.Error(1)
}

func (m *MockupArtisRepository) CountRating(idArtist uint) (uint, float32, error) {
	res := m.Called(idArtist)

	return res.Get(0).(uint), res.Get(1).(float32), res.Error(2)
}

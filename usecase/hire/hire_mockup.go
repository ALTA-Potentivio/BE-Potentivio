package hire

import (
	"github.com/stretchr/testify/mock"
	"potentivio-app/entities"
)

type MockupHireRepository struct {
	mock.Mock
}

func (m *MockupHireRepository) CreateHire(hire entities.Hire) error {
	res := m.Called(hire)

	return res.Error(0)
}

func (m *MockupHireRepository) CheckHire(hire entities.Hire) entities.Hire {
	res := m.Called(hire)

	return res.Get(0).(entities.Hire)
}

func (m *MockupHireRepository) AcceptHire(hire entities.Hire) error {
	res := m.Called(hire)

	return res.Error(0)
}

func (m *MockupHireRepository) GetHireByIdArtist(IdArtist int) ([]entities.Hire, error) {
	res := m.Called(IdArtist)

	return res.Get(0).([]entities.Hire), res.Error(1)
}

func (m *MockupHireRepository) GetHireById(id int) (entities.Hire, error) {
	res := m.Called(id)

	return res.Get(0).(entities.Hire), res.Error(1)
}

func (m *MockupHireRepository) GetHireByIdCafe(IdCafe int) ([]entities.Hire, error) {
	res := m.Called(IdCafe)

	return res.Get(0).([]entities.Hire), res.Error(1)
}

func (m *MockupHireRepository) UpdateHire(id int, hire entities.Hire) error {
	res := m.Called(id, hire)

	return res.Error(0)
}

func (m *MockupHireRepository) DeleteHire(hire entities.Hire) error {
	res := m.Called(hire)

	return res.Error(0)
}

func (m *MockupHireRepository) Rating(rating entities.Rating) error {
	res := m.Called(rating)

	return res.Error(0)
}

func (m *MockupHireRepository) CallBack(hire entities.Hire) error {
	res := m.Called(hire)

	return res.Error(0)
}

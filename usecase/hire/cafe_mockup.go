package hire

import (
	"github.com/stretchr/testify/mock"
	"potentivio-app/entities"
)

type MockupCafeRepository struct {
	mock.Mock
}

func (m *MockupCafeRepository) GetCafeById(id int) (entities.Cafe, int, error) {
	res := m.Called(id)

	return res.Get(0).(entities.Cafe), res.Get(1).(int), res.Error(2)
}

func (m *MockupCafeRepository) PostCafe(cafe entities.Cafe) error {
	res := m.Called(cafe)

	return res.Error(0)
}

func (m *MockupCafeRepository) GetAllCafe(filters map[string]string) ([]entities.Cafe, error) {
	res := m.Called(filters)

	return res.Get(0).([]entities.Cafe), res.Error(1)
}

func (m *MockupCafeRepository) DeleteCafe(id int) error {
	res := m.Called(id)

	return res.Error(0)
}

func (m *MockupCafeRepository) UpdateCafe(updateCafe entities.Cafe) error {
	res := m.Called(updateCafe)

	return res.Error(0)
}

func (m *MockupCafeRepository) GetCafeByIdForHire(id uint) (entities.Cafe, error) {
	res := m.Called(id)

	return res.Get(0).(entities.Cafe), res.Error(1)
}

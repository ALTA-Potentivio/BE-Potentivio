package notification

import (
	"fmt"
	_entities "potentivio-app/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNotif(t *testing.T) {
	t.Run("TestCreateNotifSuccess", func(t *testing.T) {
		notifUseCase := NewNotifUseCase(mockNotifRepository{})
		notif, err := notifUseCase.CreateNotif(_entities.Notification{IdArtist: 1, IdCafe: 1}, 1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, uint(1), notif.IdArtist)
		assert.Equal(t, uint(1), notif.IdCafe)
	})

	t.Run("TestCreateNotifSuccess", func(t *testing.T) {
		notifUseCase := NewNotifUseCase(mockNotifRepository{})
		notif, err := notifUseCase.CreateNotif(_entities.Notification{IdArtist: 2, IdCafe: 2}, 2, 2)
		assert.Nil(t, err)
		assert.Equal(t, uint(2), notif.IdArtist)
		assert.Equal(t, uint(2), notif.IdCafe)
	})

	t.Run("TestCreateNotifError", func(t *testing.T) {
		notifUseCase := NewNotifUseCase(mockNotifRepositoryError{})
		notif, err := notifUseCase.CreateNotif(_entities.Notification{IdArtist: 1, IdCafe: 1}, 1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), notif.IdArtist)
		assert.Equal(t, uint(0), notif.IdCafe)
	})
}

func TestGetAllNotifByIdCafe(t *testing.T) {
	t.Run("TestGetAllNotifByIdCafeSuccess", func(t *testing.T) {
		notifUseCase := NewNotifUseCase(mockNotifRepository{})
		notif, rows, err := notifUseCase.GetAllNotifByIdCafe(1)
		assert.Nil(t, err)
		assert.NotNil(t, notif)
		assert.Equal(t, uint(1), rows)
	})

	t.Run("TestGetAllNotifByIdCafeError", func(t *testing.T) {
		notifUseCase := NewNotifUseCase(mockNotifRepositoryError{})
		notif, rows, err := notifUseCase.GetAllNotifByIdCafe(1)
		assert.NotNil(t, err)
		assert.Nil(t, notif)
		assert.Equal(t, uint(0), rows)
	})
}

func TestGetAllNotifByIdArtist(t *testing.T) {
	t.Run("TestGetAllNotifByIdArtistSuccess", func(t *testing.T) {
		notifUseCase := NewNotifUseCase(mockNotifRepository{})
		notif, rows, err := notifUseCase.GetAllNotifByIdArtist(1)
		assert.Nil(t, err)
		assert.NotNil(t, notif)
		assert.Equal(t, uint(1), rows)
	})

	t.Run("TestGetAllNotifByIdArtistSuccess", func(t *testing.T) {
		notifUseCase := NewNotifUseCase(mockNotifRepositoryError{})
		notif, rows, err := notifUseCase.GetAllNotifByIdArtist(1)
		assert.NotNil(t, err)
		assert.Nil(t, notif)
		assert.Equal(t, uint(0), rows)
	})
}

func TestGetNotifById(t *testing.T) {
	t.Run("TestGetNotifByIdSuccess", func(t *testing.T) {
		notifUseCase := NewNotifUseCase(mockNotifRepository{})
		notif, rows, err := notifUseCase.GetNotifById(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, rows)
		assert.Equal(t, uint(1), notif.IdArtist)
		assert.Equal(t, uint(1), notif.IdCafe)
	})

	t.Run("TestGetNotifByIdError", func(t *testing.T) {
		notifUseCase := NewNotifUseCase(mockNotifRepositoryError{})
		notif, rows, err := notifUseCase.GetNotifById(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
		assert.Equal(t, uint(0), notif.IdArtist)
		assert.Equal(t, uint(0), notif.IdCafe)
	})
}

func TestDeleteNotif(t *testing.T) {
	t.Run("TestDeleteNotifSuccess", func(t *testing.T) {
		notifUseCase := NewNotifUseCase(mockNotifRepository{})
		rows, err := notifUseCase.DeleteNotif(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, uint(1), rows)
	})

	t.Run("TestDeleteNotifError", func(t *testing.T) {
		notifUseCase := NewNotifUseCase(mockNotifRepositoryError{})
		rows, err := notifUseCase.DeleteNotif(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, uint(1), rows)
	})
}

// === mock success ===
type mockNotifRepository struct{}

func (m mockNotifRepository) CreateNotif(notif _entities.Notification) (_entities.Notification, error) {
	return notif, nil
}

func (m mockNotifRepository) GetAllNotifByIdCafe(idToken uint) ([]_entities.Notification, uint, error) {
	notif := []_entities.Notification{{IdArtist: 1, IdCafe: 1}}
	return notif, 1, nil
}

func (m mockNotifRepository) GetAllNotifByIdArtist(idToken uint) ([]_entities.Notification, uint, error) {
	notif := []_entities.Notification{{IdArtist: 1, IdCafe: 1}}
	return notif, 1, nil
}

func (m mockNotifRepository) GetNotifById(id uint) (_entities.Notification, int, error) {
	notif := _entities.Notification{IdArtist: 1, IdCafe: 1}
	return notif, 1, nil
}

func (m mockNotifRepository) DeleteNotif(id uint) (uint, error) {
	return 1, nil
}

// === mock error ===
type mockNotifRepositoryError struct{}

func (m mockNotifRepositoryError) CreateNotif(notif _entities.Notification) (_entities.Notification, error) {
	return _entities.Notification{}, fmt.Errorf("error create notif")
}

func (m mockNotifRepositoryError) GetAllNotifByIdCafe(idToken uint) ([]_entities.Notification, uint, error) {
	return nil, 0, fmt.Errorf("error get all notif by id cafe")
}

func (m mockNotifRepositoryError) GetAllNotifByIdArtist(idToken uint) ([]_entities.Notification, uint, error) {
	return nil, 0, fmt.Errorf("error get all notif by id artist")
}

func (m mockNotifRepositoryError) GetNotifById(id uint) (_entities.Notification, int, error) {
	return _entities.Notification{}, 0, fmt.Errorf("error get notif by id")
}

func (m mockNotifRepositoryError) DeleteNotif(id uint) (uint, error) {
	return 1, fmt.Errorf("error delete notif")
}

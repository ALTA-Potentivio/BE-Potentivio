package videoArtist

import (
	"errors"
	"fmt"
	_entities "potentivio-app/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostVideo(t *testing.T) {
	t.Run("TestPostVideoSuccess", func(t *testing.T) {
		videoUseCase := NewVideoUseCase(mockVideoRepository{})
		err := videoUseCase.PostVideo(_entities.VideoArtist{IdArtist: 1, VideoUrl: "youtube.com"}, 1, "satria")
		assert.Nil(t, err)
	})

	t.Run("TestPostVideoError", func(t *testing.T) {
		videoUseCase := NewVideoUseCase(mockVideoRepositoryError{})
		err := videoUseCase.PostVideo(_entities.VideoArtist{IdArtist: 1, VideoUrl: "youtube.com"}, 1, "satria")
		if assert.Error(t, err) {
			assert.Equal(t, errors.New("failed"), err)
		}
		assert.NotNil(t, err)
	})

	t.Run("TestPostVideoErrorLinkRequires", func(t *testing.T) {
		videoUseCase := NewVideoUseCase(mockVideoRepositoryError{})
		err := videoUseCase.PostVideo(_entities.VideoArtist{IdArtist: 1}, 1, "satria")
		if assert.Error(t, err) {
			assert.Equal(t, errors.New("please insert link video"), err)
		}
		assert.NotNil(t, err)
	})
}

func TestDeleteVideo(t *testing.T) {
	t.Run("TestDeleteVideoSuccess", func(t *testing.T) {
		videoUseCase := NewVideoUseCase(mockVideoRepository{})
		err := videoUseCase.DeleteVideo(1, 1, "satria")
		assert.Nil(t, err)
	})

	t.Run("TestDeleteVideoError", func(t *testing.T) {
		videoUseCase := NewVideoUseCase(mockVideoRepositoryError{})
		err := videoUseCase.DeleteVideo(1, 1, "satria")
		if assert.Error(t, err) {
			assert.Equal(t, errors.New("failed"), err)
		}
		assert.NotNil(t, err)
	})

}

// === mock success ===
type mockVideoRepository struct{}

func (m mockVideoRepository) PostVideo(video _entities.VideoArtist, name string) error {
	return nil
}

func (m mockVideoRepository) DeleteVideo(id int, idToken int, name string) error {
	return nil
}

// === mock error ===

type mockVideoRepositoryError struct{}

func (m mockVideoRepositoryError) PostVideo(video _entities.VideoArtist, name string) error {
	return fmt.Errorf("failed")
}

func (m mockVideoRepositoryError) DeleteVideo(id int, idToken int, name string) error {
	return fmt.Errorf("failed")
}

package auth

import (
	"fmt"
	_entities "potentivio-app/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginArtis(t *testing.T) {
	t.Run("TestLoginArtistSuccess", func(t *testing.T) {
		authUseCase := NewAuthUseCase(mockAuthRepository{})
		token, idArtist, validation, err := authUseCase.LoginArtist("usamah@gmail.com", "usamah")
		assert.Nil(t, err)
		assert.Equal(t, "token", token)
		assert.Equal(t, uint(1), idArtist)
		assert.Equal(t, "not complete", validation)
	})

	t.Run("TestLoginArtistError", func(t *testing.T) {
		authUseCase := NewAuthUseCase(mockAuthRepositoryError{})
		token, idArtist, validation, err := authUseCase.LoginArtist("usamah@gmail.com", "abdurrahman")
		assert.NotNil(t, err)
		assert.Equal(t, "password incorrect", token)
		assert.Equal(t, uint(1), idArtist)
		assert.Equal(t, "not complete", validation)
	})
}

func TestLoginCafe(t *testing.T) {
	t.Run("TestLoginCafeSuccess", func(t *testing.T) {
		authUseCase := NewAuthUseCase(mockAuthRepository{})
		token, idCafe, validation, err := authUseCase.LoginCafe("cafesemesta@gmail.com", "cafesemesta")
		assert.Nil(t, err)
		assert.Equal(t, "token", token)
		assert.Equal(t, uint(1), idCafe)
		assert.Equal(t, "not complete", validation)
	})

	t.Run("TestLoginCafeError", func(t *testing.T) {
		authUseCase := NewAuthUseCase(mockAuthRepositoryError{})
		token, idCafe, validation, err := authUseCase.LoginCafe("cafesemesta@gmail.com", "cafe")
		assert.NotNil(t, err)
		assert.Equal(t, "password incorrect", token)
		assert.Equal(t, uint(1), idCafe)
		assert.Equal(t, "not complete", validation)
	})
}

// === mock success ===
type mockAuthRepository struct{}

func (m mockAuthRepository) LoginArtist(email string, password string) (string, uint, string, error) {
	artist := _entities.Artist{Name: "usamah", Email: "usamah@gmail.com", Password: "usamah", Address: "Kota Bogor"}

	if artist.Password != password {
		return "password incorrect", 1, "not complete", fmt.Errorf("password incorrect")
	}
	if artist.Description == nil {
		return "token", 1, "not complete", nil
	}
	return "token", 1, "complete", nil
}

func (m mockAuthRepository) LoginCafe(email string, password string) (string, uint, string, error) {
	cafe := _entities.Cafe{Name: "Cafe Semesta", Owner: "usamah", Email: "cafesemesta@gmail.com", Password: "cafesemesta", Address: "Kota Bogor"}

	if cafe.Password != password {
		return "password incorrect", 1, "not complete", fmt.Errorf("password incorrect")
	}
	if cafe.Description == nil {
		return "token", 1, "not complete", nil
	}
	return "token", 1, "complete", nil
}

// === mock error ===
type mockAuthRepositoryError struct{}

func (m mockAuthRepositoryError) LoginArtist(email string, password string) (string, uint, string, error) {
	artist := _entities.Artist{Name: "usamah", Email: "usamah@gmail.com", Password: "usamah", Address: "Kota Bogor"}

	if artist.Password != password {
		return "password incorrect", 1, "not complete", fmt.Errorf("password incorrect")
	}
	if artist.Description == nil {
		return "token", 1, "not complete", nil
	}
	return "token", 1, "complete", nil
}

func (m mockAuthRepositoryError) LoginCafe(email string, password string) (string, uint, string, error) {
	cafe := _entities.Cafe{Name: "Cafe Semesta", Owner: "usamah", Email: "cafesemesta@gmail.com", Password: "cafesemesta", Address: "Kota Bogor"}

	if cafe.Password != password {
		return "password incorrect", 1, "not complete", fmt.Errorf("password incorrect")
	}
	if cafe.Description == nil {
		return "token", 1, "not complete", nil
	}
	return "token", 1, "complete", nil
}

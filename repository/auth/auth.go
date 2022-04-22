package auth

import (
	"errors"
	"potentivio-app/delivery/helper"
	_middlewares "potentivio-app/delivery/middlewares"
	_entities "potentivio-app/entities"

	"gorm.io/gorm"
)

type AuthRepository struct {
	database *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		database: db,
	}
}

func (ar *AuthRepository) LoginArtist(email string, password string) (string, uint, error) {
	var artist _entities.Artist
	//mencari artist dengan menggunakan email
	tx := ar.database.Where("email = ?", email).Find(&artist)
	if tx.Error != nil {
		return "failed to login", artist.ID, tx.Error
	}

	//jika data artist dengan email tsb tidak ada
	if tx.RowsAffected == 0 {
		return "artist not found", artist.ID, errors.New("artist not found")
	}

	if helper.CheckPassHash(password, artist.Password) {
		return "password incorrect", artist.ID, errors.New("password incorrect")
	}

	//jika password sama
	token, err := _middlewares.CreateToken(int(artist.ID), artist.Name)
	if err != nil {
		return "create token failed", artist.ID, err
	}
	return token, artist.ID, nil
}

func (ar *AuthRepository) LoginCafe(email string, password string) (string, uint, error) {
	var cafe _entities.Cafe
	//mencari cafe dengan menggunakan email
	tx := ar.database.Where("email = ?", email).Find(&cafe)
	if tx.Error != nil {
		return "failed to login", cafe.ID, tx.Error
	}

	//jika data cafe dengan email tsb tidak ada
	if tx.RowsAffected == 0 {
		return "cafe not found", cafe.ID, errors.New("cafe not found")
	}

	if helper.CheckPassHash(password, cafe.Password) {
		return "password incorrect", cafe.ID, errors.New("password incorrect")
	}

	//jika password sama
	token, err := _middlewares.CreateToken(int(cafe.ID), cafe.Name)
	if err != nil {
		return "create token failed", cafe.ID, err
	}
	return token, cafe.ID, nil
}

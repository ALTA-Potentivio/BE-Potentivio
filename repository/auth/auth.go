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

func (ar *AuthRepository) LoginArtist(email string, password string) (string, uint, string, error) {
	var artist _entities.Artist
	//mencari artist dengan menggunakan email
	tx := ar.database.Where("email = ?", email).Find(&artist)
	if tx.Error != nil {
		return "failed to login", artist.ID, "not complete", tx.Error
	}

	//jika data artist dengan email tsb tidak ada
	if tx.RowsAffected == 0 {
		return "artist not found", artist.ID, "not complete", errors.New("artist not found")
	}

	if !helper.CheckPassHash(password, artist.Password) {
		return "password incorrect", artist.ID, "not complete", errors.New("password incorrect")
	}

	//jika password sama
	token, err := _middlewares.CreateToken(int(artist.ID), artist.Name)
	if err != nil {
		return "create token failed", artist.ID, "not complete", err
	}
	if artist.Description == nil {
		return token, artist.ID, "not complete", nil
	}
	return token, artist.ID, "complete", nil
}

func (ar *AuthRepository) LoginCafe(email string, password string) (string, uint, string, error) {
	var cafe _entities.Cafe
	//mencari cafe dengan menggunakan email
	tx := ar.database.Where("email = ?", email).Find(&cafe)
	if tx.Error != nil {
		return "failed to login", cafe.ID, "not complete", tx.Error
	}

	//jika data cafe dengan email tsb tidak ada
	if tx.RowsAffected == 0 {
		return "cafe not found", cafe.ID, "not complete", errors.New("cafe not found")
	}

	if !helper.CheckPassHash(password, cafe.Password) {
		return "password incorrect", cafe.ID, "not complete", errors.New("password incorrect")
	}

	//jika password sama
	token, err := _middlewares.CreateToken(int(cafe.ID), cafe.Name)
	if err != nil {
		return "create token failed", cafe.ID, "not complete", err
	}
	if cafe.Description == nil {
		return token, cafe.ID, "not complete", nil
	}
	return token, cafe.ID, "complete", nil
}

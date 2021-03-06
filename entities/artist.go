package entities

import "gorm.io/gorm"

type Artist struct {
	gorm.Model
	IdCatagory    *uint          `json:"id_catagory" form:"id_catagory"`
	IdGenre       *uint          `json:"id_genre" form:"id_genre"`
	Name          string         `gorm:"not null" json:"artist_name" form:"artist_name"`
	Email         string         `gorm:"unique;not null" json:"email" form:"email"`
	Password      string         `gorm:"not null" json:"password" form:"password"`
	Address       string         `gorm:"not null" json:"address" form:"address"`
	PhoneNumber   *string        `json:"phone_number" form:"phone_number"`
	Price         *float64       `json:"price" form:"price"`
	Description   *string        `json:"description" form:"description"`
	AccountNumber *string        `json:"account_number" form:"account_number"`
	Avatar        *string        `json:"avatar" form:"avatar"`
	Rating        float32        `json:"rating" form:"rating"`
	TotalRate     uint           `json:"total_rate" form:"total_rate"`
	Catagory      Catagory       `gorm:"foreignKey:IdCatagory;references:ID"`
	Genre         Genre          `gorm:"foreignKey:IdGenre;references:ID"`
	VideoArtist   []VideoArtist  `gorm:"foreignKey:IdArtist;references:ID"`
	Hire          []Hire         `gorm:"foreignKey:IdArtist;references:ID"`
	Notification  []Notification `gorm:"foreignKey:IdArtist;references:ID"`
	Ratings       Rating         `gorm:"foreignKey:IdArtist;references:ID"`
}

type Rating struct {
	gorm.Model
	IdArtist uint `gorm:"not null" json:"id_artist" form:"id_artist"`
	Rating   uint `json:"rating" form:"rating"`
}

package entities

import "gorm.io/gorm"

type ImageArtist struct {
	gorm.Model
	IdArtist uint   `gorm:"not null" json:"id_artist" form:"id_artist"`
	ImageUrl string `gorm:"not null" json:"image_url" form:"image_url"`
}

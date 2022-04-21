package entities

import "gorm.io/gorm"

type VideoArtist struct {
	gorm.Model
	IdArtist uint   `gorm:"not null" json:"id_artist" form:"id_artist"`
	VideoUrl string `gorm:"not null" json:"video_url" form:"video_url"`
}

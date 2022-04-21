package entities

import (
	"time"

	"gorm.io/gorm"
)

type Hire struct {
	gorm.Model
	IdArtist uint      `gorm:"not null" json:"id_artist" form:"id_artist"`
	IdCafe   uint      `gorm:"not null" json:"id_cafe" form:"id_cafe"`
	Date     time.Time `gorm:"not null" json:"date" form:"date"`
	Status   string    `gorm:"not null;default:waiting" json:"status" form:"status"`
	Comment  string    `gorm:"not null" json:"comment" form:"comment"`
	Rating   uint      `gorm:"not null" json:"rating" form:"rating"`
	Artist   Artist    `gorm:"foreignKey:IdArtist;references:ID"`
	Cafe     Cafe      `gorm:"foreignKey:IdCafe;references:ID"`
}

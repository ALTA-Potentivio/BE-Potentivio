package entities

import (
	"time"

	"gorm.io/gorm"
)

type Hire struct {
	gorm.Model
	IdArtist            uint      `gorm:"not null" json:"id_artist" form:"id_artist"`
	IdCafe              uint      `gorm:"not null" json:"id_cafe" form:"id_cafe"`
	Invoice             string    `gorm:"not null" json:"invoice" form:"invoice"`
	Date                time.Time `gorm:"not null" json:"date" form:"date"`
	StatusArtist        *string   `json:"status_artist" form:"status_artist"`
	StatusCafe          string    `gorm:"not null;default:waiting" json:"status_cafe" form:"status_cafe"`
	Comment             string    `gorm:"not null" json:"comment" form:"comment"`
	Rating              uint      `gorm:"not null" json:"rating" form:"rating"`
	Price               uint      `gorm:"not null" json:"price" form:"price"`
	AccountNumberArtist string    `gorm:"not null" json:"account_number_artist" form:"account_number_artist"`
	AccountNumberCafe   string    `gorm:"not null" json:"account_number_cafe" form:"account_number_cafe"`
	Artist              Artist    `gorm:"foreignKey:IdArtist;references:ID"`
	Cafe                Cafe      `gorm:"foreignKey:IdCafe;references:ID"`
}

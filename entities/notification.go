package entities

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	IdArtist uint   `gorm:"not null" json:"id_artist" form:"id_artist"`
	IdCafe   uint   `gorm:"not null" json:"id_cafe" form:"id_cafe"`
	Artist   Artist `gorm:"foreignKey:IdArtist;references:ID"`
}

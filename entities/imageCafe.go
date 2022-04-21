package entities

import "gorm.io/gorm"

type ImageCafe struct {
	gorm.Model
	IdCafe   uint   `gorm:"not null" json:"id_cafe" form:"id_cafe"`
	ImageUrl string `gorm:"not null" json:"image_url" form:"image_url"`
}

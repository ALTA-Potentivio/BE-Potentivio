package hire

import (
	"time"
)

type HireRequest struct {
	IdArtist            uint      `gorm:"not null" json:"id_artist" form:"id_artist"`
	IdCafe              uint      `gorm:"not null" json:"id_cafe" form:"id_cafe"`
	Invoice             string    `json:"invoice" form:"invoice"`
	Date                time.Time `gorm:"not null" json:"date" form:"date"`
	Price               float64   `gorm:"not null" json:"price" form:"price"`
	AccountNumberArtist *string   `gorm:"not null" json:"account_number_artist" form:"account_number_artist"`
	AccountNumberCafe   *string   `gorm:"not null" json:"account_number_cafe" form:"account_number_cafe"`
	PaymentUrl          *string   `json:"payment_url" form:"payment_url"`
}
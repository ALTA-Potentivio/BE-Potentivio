package entities

import (
	"time"

	"gorm.io/gorm"
)

type Cafe struct {
	gorm.Model
	Name          string         `gorm:"not null" json:"cafe_name" form:"cafe_name"`
	Owner         string         `gorm:"not null" json:"owner" form:"owner"`
	Email         string         `gorm:"unique;not null" json:"email" form:"email"`
	Password      string         `gorm:"not null" json:"password" form:"password"`
	Address       string         `gorm:"not null" json:"address" form:"address"`
	PhoneNumber   *string        `json:"phone_number" form:"phone_number"`
	Description   *string        `json:"description" form:"description"`
	OpeningHours  *time.Time     `json:"opening_hours" form:"opening_hours"`
	AccountNumber *string        `json:"account_number" form:"account_number"`
	Avatar        *string        `json:"avatar" form:"avatar"`
	Longitude     *float64       `json:"longitude" form:"longitude"`
	Latitude      *float64       `json:"latitude" form:"latitude"`
	ImageCafe     []ImageCafe    `gorm:"foreignKey:IdCafe;references:ID"`
	Hire          []Hire         `gorm:"foreignKey:IdCafe;references:ID"`
	Notification  []Notification `gorm:"foreignKey:IdCafe;references:ID"`
}

type GetCafe struct {
	ID            uint        `json:"id" form:"id"`
	Name          string      `json:"cafe_name" form:"cafe_name"`
	Owner         string      `json:"owner" form:"owner"`
	Email         string      `json:"email" form:"email"`
	PhoneNumber   *string     `json:"phone_number" form:"phone_number"`
	Address       string      `gorm:"not null" json:"address" form:"address"`
	Description   *string     `json:"description" form:"description"`
	OpeningHours  *time.Time  `json:"opening_hours" form:"opening_hours"`
	Avatar        *string     `json:"avatar" form:"avatar"`
	AccountNumber *string     `json:"account_number" form:"account_number"`
	Longitude     *float64    `json:"longitude" form:"longitude"`
	Latitude      *float64    `json:"latitude" form:"latitude"`
	ImageCafe     []ImageCafe `gorm:"foreignKey:IdCafe;references:ID"`
}
type GetAllCafe struct {
	ID          uint    `json:"id" form:"id"`
	Name        string  `json:"cafe_name" form:"cafe_name"`
	Address     string  `gorm:"not null" json:"address" form:"address"`
	Description *string `json:"description" form:"description"`
	Avatar      *string `json:"avatar" form:"avatar"`
}

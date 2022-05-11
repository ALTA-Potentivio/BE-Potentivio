package entities

type Catagory struct {
	ID           uint     `gorm:"primarykey"`
	NameCatagory string   `gorm:"unique;not null" json:"name_catagory" form:"name_catagory"`
	Artist       []Artist `gorm:"foreignKey:IdCatagory;references:ID"`
}

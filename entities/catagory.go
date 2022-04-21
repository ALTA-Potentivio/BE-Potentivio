package entities

type Catagory struct {
	ID           uint     `gorm:"primarykey"`
	NameCatagory string   `gorm:"not null" json:"catagory_name" form:"catagory_name"`
	Artist       []Artist `gorm:"foreignKey:IdCatagory;references:ID"`
}

package entities

type Genre struct {
	ID        uint     `gorm:"primarykey"`
	NameGenre string   `gorm:"not null" json:"name_genre" form:"name_genre"`
	Artist    []Artist `gorm:"foreignKey:IdGenre;references:ID"`
}

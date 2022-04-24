package hire

type HireRequest struct {
	Date string `gorm:"not null" json:"date" form:"date"`
}

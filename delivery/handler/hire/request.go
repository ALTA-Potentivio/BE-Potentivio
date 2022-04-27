package hire

type HireRequest struct {
	Date string `gorm:"not null" json:"date" form:"date"`
}

type CallBackRequest struct {
	Invoice string `json:"external_id"`
	Status  string `json:"status"`
}

package hire

type HireResponse struct {
	Id           int    `json:"id"`
	CafeName     string `json:"cafe_name"`
	Date         string `json:"date"`
	StatusArtist string `json:"status_artist"`
}

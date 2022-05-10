package hire

type HireResponse struct {
	Id           int    `json:"id"`
	CafeName     string `json:"cafe_name"`
	Comment      string `json:"comment"`
	Date         string `json:"date"`
	StatusArtist string `json:"status_artist"`
}

type HireCafeResponse struct {
	Id         int    `json:"id"`
	ArtisName  string `json:"artis_name"`
	Date       string `json:"date"`
	StatusCafe string `json:"status_cafe"`
	PaymentUrl string `json:"paymentUrl"`
}

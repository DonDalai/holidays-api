package models

type Holiday struct {
	Date        string `json:"date"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Inalienable bool   `json:"inalienable"`
	Extra       string `json:"extra"`
}

type HolidayResponse struct {
	Status string    `json:"status"`
	Data   []Holiday `json:"data"`
}

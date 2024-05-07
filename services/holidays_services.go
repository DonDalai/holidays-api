package services

import (
	"io"
	"net/http"
)

type HolidayService struct {
	APIEndpoint string
}

func NewHolidayService(endpoint string) *HolidayService {
	return &HolidayService{APIEndpoint: endpoint}
}

func (hs *HolidayService) FetchHolidays() ([]byte, error) {
	resp, err := http.Get(hs.APIEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

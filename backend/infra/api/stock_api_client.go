package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type APIResponse struct {
	Items    []StockAPIItem `json:"items"`
	NextPage string         `json:"next_page"`
}

type StockAPIItem struct {
	Ticker     string `json:"ticker"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Company    string `json:"company"`
	Action     string `json:"action"`
	Brokerage  string `json:"brokerage"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	Time       string `json:"time"`
}

type StockAPIClient interface {
	FetchPage(nextPage string) (*APIResponse, error)
}

type DefaultStockAPIClient struct{}

func NewStockAPIClient() StockAPIClient {
	return &DefaultStockAPIClient{}
}

func (client *DefaultStockAPIClient) FetchPage(nextPage string) (*APIResponse, error) {
	url := os.Getenv("API_URL")
	if nextPage != "" {
		url += "?next_page=" + nextPage
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+os.Getenv("API_AUTH_TOKEN"))
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	var result APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

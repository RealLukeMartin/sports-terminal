package scraper

import (
	"encoding/json"
	"io"
	"net/http"
)

type TypeData struct {
	Type string `json:"type"`
}

type ScoreboardData struct {
	Events []struct {
		GameID    string `json:"id"`
		ShortName string `json:"shortName"`
		// Status    string `json:"status[type][name]"`
		Status struct {
			Period int `json:"period"`
			Type   struct {
				Name string `json:"name"`
			} `json:"type"`
		} `json:"status"`
		// Add other fields as needed
	} `json:"events"`
}

func fetchJSON(url string) ([]byte, error) {
	resp, err := http.Get(url)

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

func parseJSON(data []byte) (*ScoreboardData, error) {
	var scoreboard ScoreboardData
	err := json.Unmarshal(data, &scoreboard)
	if err != nil {
		return nil, err
	}
	return &scoreboard, nil
}

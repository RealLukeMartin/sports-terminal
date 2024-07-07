package scraper

import (
	"encoding/json"
)

type TypeData struct {
	Type string `json:"type"`
}

type MlbScoreboardData struct {
	Events []struct {
		GameID       string `json:"id"`
		ShortName    string `json:"shortName"`
		Competitions []struct {
			Id          string `json:"id"`
			Competitors []struct {
				Id         string `json:"id"`
				HomeAway   string `json:"homeAway"`
				Order      int    `json:"order"`
				Score      string `json:"score"`
				Linescores []struct {
					Value float32 `json:"value"`
				} `json:"linescores,omitempty"`
				Statistics []struct {
					Name         string `json:"name"`
					DisplayValue string `json:"displayValue"`
				} `json:"statistics,omitempty"`
				Team struct {
					Id       string `json:"id"`
					Location string `json:"location"`
					Name     string `json:"name"`
				} `json:"team"`
			} `json:"competitors"`
		} `json:"competitions"`
		Status struct {
			Period int `json:"period"`
			Type   struct {
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"type"`
		} `json:"status"`
	} `json:"events"`
}

func parseJSON(data []byte) (*MlbScoreboardData, error) {
	var scoreboard MlbScoreboardData
	err := json.Unmarshal(data, &scoreboard)
	if err != nil {
		return nil, err
	}
	return &scoreboard, nil
}

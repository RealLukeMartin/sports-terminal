package scraper

import (
	"encoding/json"
)

type MlbStatistic struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	DisplayValue string `json:"displayValue"`
}

type MlbTeam struct {
	Id           string `json:"id"`
	Abbreviation string `json:"abbreviation"`
	Location     string `json:"location"`
	Name         string `json:"name"`
}

type MlbScoreboardData struct {
	Events []struct {
		GameID       string `json:"id"`
		Name         string `json:"name"`
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
				Statistics []MlbStatistic `json:"statistics,omitempty"`
				Team       MlbTeam        `json:"team"`
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

func parseMlbJSON(data []byte) (*MlbScoreboardData, error) {
	var scoreboard MlbScoreboardData
	err := json.Unmarshal(data, &scoreboard)
	if err != nil {
		return nil, err
	}
	return &scoreboard, nil
}

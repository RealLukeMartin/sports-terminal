package scraper

import (
	"fmt"
	"log"
)

func GetMlbGames() {
	url := "https://site.api.espn.com/apis/site/v2/sports/baseball/mlb/scoreboard"
	jsonData, err := fetchJSON(url)
	if err != nil {
		log.Fatalf("Error fetching JSON: %v", err)
	}

	scoreboard, err := parseJSON(jsonData)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	fmt.Printf("%+v\n", scoreboard)
}

package boxScores

import (
	"github.com/RealLukeMartin/sports-terminal/internal/components"
	"github.com/RealLukeMartin/sports-terminal/scraper"
	tea "github.com/charmbracelet/bubbletea"
)

type MlbTeamsBoxScoreStats = components.MlbTeamsBoxScoreStats
type Inning = components.Inning

type MlbBoxScoreModel struct {
	homeStats MlbTeamsBoxScoreStats
	awayStats MlbTeamsBoxScoreStats
	innings   []Inning
}

var mlbGames = scraper.GetMlbGames()

func getTeamStats(statistics []scraper.MlbStatistic) MlbTeamsBoxScoreStats {
	var stats = MlbTeamsBoxScoreStats{}
	for _, stat := range statistics {
		if stat.Abbreviation == "R" {
			stats.Runs = stat.DisplayValue
		}
		if stat.Abbreviation == "H" {
			stats.Hits = stat.DisplayValue
		}
		if stat.Abbreviation == "E" {
			stats.Errors = stat.DisplayValue
		}
	}
	return stats
}

func processCompetitors(mlbGames *scraper.MlbScoreboardData) (MlbTeamsBoxScoreStats, MlbTeamsBoxScoreStats, []Inning) {
	var firstGame = mlbGames.Events[0]
	var firstCompetition = firstGame.Competitions[0]

	var innings = []Inning{}
	var homeInnings = []int{}
	var awayInnings = []int{}
	var homeStats = MlbTeamsBoxScoreStats{}
	var awayStats = MlbTeamsBoxScoreStats{}

	for _, competitor := range firstCompetition.Competitors {

		if competitor.HomeAway == "home" {
			homeStats.Name = competitor.Team.Abbreviation
			var stats = getTeamStats(competitor.Statistics)

			homeStats.Runs = stats.Runs
			homeStats.Hits = stats.Hits
			homeStats.Errors = stats.Errors

			for _, teamInnings := range competitor.Linescores {
				var singleInningScore = int(teamInnings.Value)
				homeInnings = append(homeInnings, singleInningScore)
			}
		}

		if competitor.HomeAway == "away" {
			awayStats.Name = competitor.Team.Abbreviation
			var stats = getTeamStats(competitor.Statistics)

			awayStats.Runs = stats.Runs
			awayStats.Hits = stats.Hits
			awayStats.Errors = stats.Errors

			for _, teamInnings := range competitor.Linescores {
				var singleInningScore = int(teamInnings.Value)
				awayInnings = append(awayInnings, singleInningScore)
			}
		}
	}

	for i := range homeInnings {
		innings = append(innings, Inning{HomeScore: homeInnings[i], AwayScore: awayInnings[i]})
	}

	return homeStats, awayStats, innings
}

func MlbBoxScoreInitialModel() MlbBoxScoreModel {
	homeStats, awayStats, innings := processCompetitors(mlbGames)
	return MlbBoxScoreModel{
		homeStats: homeStats,
		awayStats: awayStats,
		innings:   innings,
	}
}

func (m MlbBoxScoreModel) Init() tea.Cmd {
	return nil
}

func (m MlbBoxScoreModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m MlbBoxScoreModel) View() string {
	homeStats, awayStats, innings := processCompetitors(mlbGames)
	return components.MlbBoxScoreTable(innings, homeStats, awayStats)
}

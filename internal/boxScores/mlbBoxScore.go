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
var gameIndexRef *int

func MlbBoxScoreInitialModel(gameIndex int) MlbBoxScoreModel {
	homeStats, awayStats, innings := processMlbTeams(mlbGames, gameIndex)
	gameIndexRef = &gameIndex
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
	homeStats, awayStats, innings := processMlbTeams(mlbGames, *gameIndexRef)
	return components.MlbBoxScoreTable(innings, homeStats, awayStats)
}

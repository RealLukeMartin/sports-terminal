package boxScores

import (
	"github.com/RealLukeMartin/sports-terminal/internal/components"
	tea "github.com/charmbracelet/bubbletea"
)

type MlbTeamsBoxScoreStats = components.MlbTeamsBoxScoreStats
type Inning = components.Inning

type MlbBoxScoreModel struct {
	homeStats MlbTeamsBoxScoreStats
	awayStats MlbTeamsBoxScoreStats
	innings   []Inning
}

var innings = []Inning{
	{HomeScore: 0, AwayScore: 0},
	{HomeScore: 0, AwayScore: 0},
	{HomeScore: 0, AwayScore: 0},
	{HomeScore: 0, AwayScore: 0},
	{HomeScore: 0, AwayScore: 0},
	{HomeScore: 0, AwayScore: 0},
	{HomeScore: 0, AwayScore: 0},
	{HomeScore: 0, AwayScore: 0},
	{HomeScore: 0, AwayScore: 0},
}

var homeStats = MlbTeamsBoxScoreStats{
	Name:   "HOU",
	Runs:   0,
	Hits:   0,
	Errors: 0,
}

var awayStats = MlbTeamsBoxScoreStats{
	Name:   "SEA",
	Runs:   0,
	Hits:   0,
	Errors: 0,
}

func MlbBoxScoreInitialModel() MlbBoxScoreModel {
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
	//return components.MlbBoxScoreTable(m.innings, m.homeStats, m.awayStats)
	return components.MlbBoxScoreTable(innings, homeStats, awayStats)
}

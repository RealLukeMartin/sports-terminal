package boxScores

import tea "github.com/charmbracelet/bubbletea"

type mlbTeamsBoxScoreStats struct {
	name   string
	runs   int
	hits   int
	errors int
}

type inning struct {
	homeScore int
	awayScore int
}

type MlbBoxScoreModel struct {
	homeStats   mlbTeamsBoxScoreStats
	awayStats   mlbTeamsBoxScoreStats
	homeInnings []inning
	awayInnings []inning
}

func MlbBoxScoreInitialModel(homeStats mlbTeamsBoxScoreStats, awayStats mlbTeamsBoxScoreStats, homeInnings []inning, awayInnings []inning) MlbBoxScoreModel {
	return MlbBoxScoreModel{
		homeStats:   homeStats,
		awayStats:   awayStats,
		homeInnings: homeInnings,
		awayInnings: awayInnings,
	}
}

func (m MlbBoxScoreModel) Init() tea.Cmd {
	return nil
}

func (m MlbBoxScoreModel) Update() (tea.Msg, tea.Cmd) {
	return m, nil
}

func (m MlbBoxScoreModel) View() string {
	return "Box Scores Coming Soon..."
}

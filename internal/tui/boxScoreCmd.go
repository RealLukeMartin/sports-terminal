package tui

import (
	"github.com/RealLukeMartin/sports-terminal/internal/boxScores"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *tuiModel) boxScoreCmd(msg tea.Msg) tea.Cmd {
	// Update the leagues menu
	newBoxScore, newBoxScoreCmd := m.boxScore.Update(msg)

	// Get the boxScore model from the updated model
	boxScoreModel, ok := newBoxScore.(boxScores.MlbBoxScoreModel)

	if !ok {
		panic("could not perform assertion on LeaguesModel")
	}

	// Update the leagues model in state
	m.boxScore = boxScoreModel

	// set the new cmd we will be returning
	return newBoxScoreCmd
}

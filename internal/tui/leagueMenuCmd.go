package tui

import (
	"github.com/RealLukeMartin/sports-terminal/internal/menus"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *tuiModel) leagueMenuCmd(msg tea.Msg) tea.Cmd {
	// Update the leagues menu
	newLeagueMenu, newLeagueCmd := m.leagues.Update(msg)

	// Get the league model from the updated model
	leaguesModel, ok := newLeagueMenu.(menus.LeaguesModel)

	if !ok {
		panic("could not perform assertion on LeaguesModel")
	}

	// Update the leagues model in state
	m.leagues = leaguesModel

	// set the new cmd we will be returning
	return newLeagueCmd
}

package tui

import (
	"github.com/RealLukeMartin/sports-terminal/internal/menus"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *tuiModel) leagueGamesCmd(msg tea.Msg) tea.Cmd {
	// Update the leagues menu
	newLeagueGamesMenu, newLeagueGamesCmd := m.leagues.Update(msg)

	// Get the league model from the updated model
	leagueGamesModel, ok := newLeagueGamesMenu.(menus.LeaguesModel)

	if !ok {
		panic("could not perform assertion on LeaguesModel")
	}

	// Update the leagues model in state
	m.leagues = leagueGamesModel

	// set the new cmd we will be returning
	return newLeagueGamesCmd
}

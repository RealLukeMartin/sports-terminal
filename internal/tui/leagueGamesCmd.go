package tui

import (
	"github.com/RealLukeMartin/sports-terminal/internal/menus"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *tuiModel) leagueGamesCmd(msg tea.Msg) tea.Cmd {
	// Update the league Games menu
	newLeagueGamesMenu, newLeagueGamesCmd := m.leagueGames.Update(msg)

	// Get the league games model from the updated model
	leagueGamesModel, ok := newLeagueGamesMenu.(menus.LeagueGamesModel)

	if !ok {
		panic("could not perform assertion on LeagueGamesModel")
	}

	// Update the league games model in state
	m.leagueGames = leagueGamesModel

	// set the new cmd we will be returning
	return newLeagueGamesCmd
}

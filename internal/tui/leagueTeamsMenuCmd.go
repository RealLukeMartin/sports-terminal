package tui

import (
	"github.com/RealLukeMartin/sports-terminal/internal/menus"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *tuiModel) leagueTeamsMenuCmd(msg tea.Msg) tea.Cmd {
	// Update the leagueTeams menu, get the new model and cmd
	newLeagueTeamsMenu, newLeagueTeamsCmd := m.leagueTeams.Update(msg)

	// Get the leagueTeams model from the updated model
	leagueTeamsModel, ok := newLeagueTeamsMenu.(menus.LeagueTeamsModel)

	if !ok {
		panic("could not perform assertion on LeagueTeamsModel")
	}

	// Update the leagueTeams model in state
	m.leagueTeams = leagueTeamsModel

	// set the new cmd we will be returning
	return newLeagueTeamsCmd
}

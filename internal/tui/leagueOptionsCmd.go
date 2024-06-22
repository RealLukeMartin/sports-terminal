package tui

import (
	"github.com/RealLukeMartin/sports-terminal/internal/menus"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *tuiModel) leagueOptionsMenuCmd(msg tea.Msg) tea.Cmd {
	// Update the leagueOptions menu, get the new model and cmd
	newLeagueOptionsMenu, newLeagueOptionsCmd := m.leagueOptions.Update(msg)

	// Get the leagueOptions model from the updated model
	leagueOptionsModel, ok := newLeagueOptionsMenu.(menus.LeagueOptionsModel)

	if !ok {
		panic("could not perform assertion on LeagueOptionsModel")
	}

	// Update the leagueOptions model in state
	m.leagueOptions = leagueOptionsModel

	// set the new cmd we will be returning
	return newLeagueOptionsCmd
}

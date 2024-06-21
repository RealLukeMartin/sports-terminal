package tui

import (
	"github.com/RealLukeMartin/sports-terminal/internal/menus"
	tea "github.com/charmbracelet/bubbletea"
)

type currentPageState int

const (
	leaguesMenu currentPageState = iota
	leagueTeamsMenu
)

type tuiModel struct {
	currentPage   currentPageState
	leaguesCursor int
	leagues       tea.Model
	leagueTeams   tea.Model
}

func TuiInitialModel() tuiModel {
	return tuiModel{
		currentPage: leaguesMenu,
		leagues:     menus.LeaguesInitialModel(),
		leagueTeams: menus.LeagueTeamsInitialModel(0),
	}
}

func (m tuiModel) Init() tea.Cmd {
	return nil
}

func (m tuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	// Did we get key press?
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter", " ":
			if m.currentPage == leaguesMenu {
				m.leagueTeams = menus.LeagueTeamsInitialModel(m.leaguesCursor)
				m.currentPage = leagueTeamsMenu
			}
		case "esc":
			if m.currentPage == leagueTeamsMenu {
				m.currentPage = leaguesMenu
			}

		case "up", "w", "k":
			if m.currentPage == leaguesMenu {
				if m.leaguesCursor > 0 {
					m.leaguesCursor--
				}
			}

		case "down", "s", "j":
			if m.currentPage == leaguesMenu {
				if m.leaguesCursor < 3 { // TODO: Make this dynamic
					m.leaguesCursor++
				}
			}
		}

		switch m.currentPage {
		case leaguesMenu:
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
			cmd = newLeagueCmd

		case leagueTeamsMenu:
			// Update the mlb menu, get the new model and cmd
			newLeagueTeamsMenu, newLeagueTeamsCmd := m.leagueTeams.Update(msg)

			// Get the mlb model from the updated model
			leagueTeamsModel, ok := newLeagueTeamsMenu.(menus.LeagueTeamsModel)

			if !ok {
				panic("could not perform assertion on MlbModel")
			}

			// Update the mlb model in state
			m.leagueTeams = leagueTeamsModel

			// set the new cmd we will be returning
			cmd = newLeagueTeamsCmd
		}

	}
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m tuiModel) View() string {

	switch m.currentPage {
	case leaguesMenu: // Leagues
		return m.leagues.View()
	case leagueTeamsMenu: // League Teams
		return m.leagueTeams.View()
	}

	return "Oh no..."
}

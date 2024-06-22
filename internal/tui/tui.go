package tui

import (
	"github.com/RealLukeMartin/sports-terminal/internal/data"
	"github.com/RealLukeMartin/sports-terminal/internal/menus"
	"github.com/RealLukeMartin/sports-terminal/internal/styles"
	tea "github.com/charmbracelet/bubbletea"
)

type currentPageState int

const (
	leaguesMenu currentPageState = iota
	leagueOptionsMenu
	leagueTeamsMenu
)

type tuiModel struct {
	currentPage   currentPageState
	leaguesCursor int
	leagues       tea.Model
	leagueOptions tea.Model
	leagueTeams   tea.Model
}

func TuiInitialModel() tuiModel {
	return tuiModel{
		currentPage:   leaguesMenu,
		leagues:       menus.LeaguesInitialModel(),
		leagueOptions: menus.LeagueOptionsInitialModel(),
		leagueTeams:   menus.LeagueTeamsInitialModel(0),
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
			if m.currentPage == leagueOptionsMenu {
				m.currentPage = leagueTeamsMenu
			}

			if m.currentPage == leaguesMenu {
				// Set the league teams model to the selected league
				m.leagueTeams = menus.LeagueTeamsInitialModel(m.leaguesCursor)

				// Set the current page to the league options menu
				m.currentPage = leagueOptionsMenu
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
				if m.leaguesCursor < len(data.LeagueOptions)-1 {
					m.leaguesCursor++
				}
			}
		}

		switch m.currentPage {
		case leaguesMenu:
			cmd = m.leagueMenuCmd(msg)

		case leagueOptionsMenu:
			cmd = m.leagueOptionsMenuCmd(msg)

		case leagueTeamsMenu:
			cmd = m.leagueTeamsMenuCmd(msg)
		}

	}
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m tuiModel) View() string {

	switch m.currentPage {
	case leaguesMenu: // Leagues
		return styles.DefaultStyle.Render(m.leagues.View())
	case leagueOptionsMenu: // League Options
		return styles.DefaultStyle.Render(m.leagueOptions.View())
	case leagueTeamsMenu: // League Teams
		return styles.DefaultStyle.Render(m.leagueTeams.View())
	}

	return "Oh no..."
}

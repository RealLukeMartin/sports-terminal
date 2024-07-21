package tui

import (
	"github.com/RealLukeMartin/sports-terminal/internal/boxScores"
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
	boxScore
	leagueGames
)

type tuiModel struct {
	currentPage         currentPageState
	leaguesCursor       int
	leagueOptionsCursor int
	leagueGamesCursor   int
	leagues             tea.Model
	leagueOptions       tea.Model
	leagueTeams         tea.Model
	leagueGames         tea.Model
	boxScore            tea.Model
}

func TuiInitialModel() tuiModel {
	return tuiModel{
		currentPage:   leaguesMenu,
		leagues:       menus.LeaguesInitialModel(),
		leagueOptions: menus.LeagueOptionsInitialModel(),
		leagueTeams:   menus.LeagueTeamsInitialModel(0),
		leagueGames:   menus.LeagueGamesInitialModel(),
		boxScore:      boxScores.MlbBoxScoreInitialModel(0),
	}
}

func (m tuiModel) Init() tea.Cmd {
	return nil
}

func (m tuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	// TODO: Refactor all this
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter", " ":
			if m.currentPage == leagueGames {
				// Set the box score model to the selected game
				m.boxScore = boxScores.MlbBoxScoreInitialModel(m.leagueGamesCursor)
				m.currentPage = boxScore
			}

			if m.currentPage == leagueOptionsMenu {
				switch m.leagueOptionsCursor {
				case 0:
					m.currentPage = leagueTeamsMenu

				case 1:
					m.currentPage = leagueGames
				}
			}

			if m.currentPage == leaguesMenu {
				// Set the league teams model to the selected league
				m.leagueTeams = menus.LeagueTeamsInitialModel(m.leaguesCursor)

				// Set the current page to the league options menu
				m.currentPage = leagueOptionsMenu
			}

		case "esc":
			if m.currentPage == leagueOptionsMenu {
				m.currentPage = leaguesMenu
			}

			if m.currentPage == leagueTeamsMenu {
				m.currentPage = leagueOptionsMenu
			}

			if m.currentPage == leagueGames {
				m.currentPage = leagueOptionsMenu
			}

			if m.currentPage == boxScore {
				m.currentPage = leagueGames
			}

		case "up", "w", "k":
			if m.currentPage == leaguesMenu {
				if m.leaguesCursor > 0 {
					m.leaguesCursor--
				}
			}

			if m.currentPage == leagueOptionsMenu {
				if m.leagueOptionsCursor > 0 {
					m.leagueOptionsCursor--
				}
			}

			if m.currentPage == leagueGames {
				if m.leagueGamesCursor > 0 {
					m.leagueGamesCursor--
				}
			}

		case "down", "s", "j":
			if m.currentPage == leaguesMenu {
				if m.leaguesCursor < len(data.LeagueOptions)-1 {
					m.leaguesCursor++
				}
			}

			if m.currentPage == leagueOptionsMenu {
				if m.leagueOptionsCursor < 3 { // TODO: Clean up this hard coded value
					m.leagueOptionsCursor++
				}
			}

			if m.currentPage == leagueGames {
				// if m.leagueGamesCursor < len(data.LeagueOptions)-1 {
				m.leagueGamesCursor++
				//}
			}
		}

		switch m.currentPage {
		case leaguesMenu:
			cmd = m.leagueMenuCmd(msg)

		case leagueOptionsMenu:
			cmd = m.leagueOptionsMenuCmd(msg)

		case leagueTeamsMenu:
			cmd = m.leagueTeamsMenuCmd(msg)

		case boxScore:
			cmd = m.boxScoreCmd(msg)

		case leagueGames:
			cmd = m.leagueGamesCmd(msg)
		}

	}
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m tuiModel) View() string {

	footer := "\nPress 'Esc' to go back, 'q' to quit."
	switch m.currentPage {
	case leaguesMenu: // Leagues
		return styles.DefaultStyle.Render(m.leagues.View())
	case leagueOptionsMenu: // League Options
		return styles.DefaultStyle.Render(m.leagueOptions.View()) + footer
	case leagueTeamsMenu: // League Teams
		return styles.DefaultStyle.Render(m.leagueTeams.View()) + footer
	case boxScore: // Box Score
		return m.boxScore.View() + footer
	case leagueGames: // League Games
		return styles.DefaultStyle.Render(m.leagueGames.View() + footer)
	}

	return "Oh no..."
}

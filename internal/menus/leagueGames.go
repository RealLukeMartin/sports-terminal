package menus

import tea "github.com/charmbracelet/bubbletea"

type LeagueGamesModel struct {
	games []string
}

func LeagueGamesInitialModel() LeagueGamesModel {
	return LeagueGamesModel{
		games: []string{},
	}
}

func (m LeagueGamesModel) Init() tea.Cmd {
	return nil
}

func (m LeagueGamesModel) Update() (tea.Msg, tea.Cmd) {
	return m, nil
}

func (m LeagueGamesModel) View() string {
	return "Games Coming Soon..."
}

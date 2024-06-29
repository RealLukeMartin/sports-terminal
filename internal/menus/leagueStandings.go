package menus

import tea "github.com/charmbracelet/bubbletea"

type LeagueStandingsModel struct {
	conferences []string
	divisions   []string
	teams       []string
}

func LeageStandingsInitialModel(conferences []string, divisons []string, teams []string) LeagueStandingsModel {
	return LeagueStandingsModel{
		conferences: conferences,
		divisions:   divisons,
		teams:       teams,
	}
}

func (m LeagueStandingsModel) Init() tea.Cmd {
	return nil
}

func (m LeagueStandingsModel) Update() (tea.Msg, tea.Cmd) {
	return m, nil
}

func (m LeagueStandingsModel) View() string {
	return "Standings Coming Soon..."
}

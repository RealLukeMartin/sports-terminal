package menus

import (
	"fmt"

	"github.com/RealLukeMartin/sports-terminal/internal/data"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/exp/maps"
)

// const (
// 	mlb leagueSelection = iota
// 	nfl
// 	nba
// 	nhl
// )

type LeagueTeamsModel struct {
	teams    []string
	cursor   int
	selected map[int]struct{}
}

func LeagueTeamsInitialModel(leagueSelection int) LeagueTeamsModel {
	var teamsData []string
	switch leagueSelection {
	case 0:
		teamsData = maps.Keys(data.MlbTeams)
	case 1:
		teamsData = maps.Keys(data.NflTeams)
	case 2:
		teamsData = maps.Keys(data.NbaTeams)
	case 3:
		teamsData = maps.Keys(data.NhlTeams)
	default:
		panic("Invalid league selection")
	}

	teams := teamsData
	return LeagueTeamsModel{
		teams:    teams,
		selected: make(map[int]struct{}),
	}
}

func (m LeagueTeamsModel) Init() tea.Cmd {
	return nil
}

func (m LeagueTeamsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "up", "w", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "s", "j":
			if m.cursor < len(m.teams)-1 {
				m.cursor++
			}

		case "enter", " ":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m LeagueTeamsModel) View() string {
	prompt := "Select Team:\n\n"

	for i, team := range m.teams {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		prompt += fmt.Sprintf("%s %s\n", cursor, team)
	}

	prompt += "\nPress 'Esc' to go back to leagues menu, 'q' to quit."

	return prompt
}

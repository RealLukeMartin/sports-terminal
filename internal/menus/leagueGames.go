package menus

import (
	"github.com/RealLukeMartin/sports-terminal/scraper"
	tea "github.com/charmbracelet/bubbletea"
)

type LeagueGamesModel struct {
	cursor int
	games  []string
}

func LeagueGamesInitialModel() LeagueGamesModel {
	var games = scraper.GetMlbGames()
	gamesList := []string{}

	for _, game := range games.Events {
		gamesList = append(gamesList, game.Name)
	}

	return LeagueGamesModel{
		games: gamesList,
	}
}

func (m LeagueGamesModel) Init() tea.Cmd {
	return nil
}

func (m LeagueGamesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "w", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "a", "j":
			if m.cursor < len(m.games) {
				m.cursor++
			}
		}
	}

	return m, nil
}

func (m LeagueGamesModel) View() string {
	prompt := "Select Game:\n\n"
	for index, option := range m.games {
		cursor := " "
		if m.cursor == index {
			cursor = ">"
		}
		prompt += cursor + " " + option + "\n"
	}

	return prompt
}

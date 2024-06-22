package menus

import (
	"github.com/RealLukeMartin/sports-terminal/internal/data"
	tea "github.com/charmbracelet/bubbletea"
)

type LeagueOptionsModel struct {
	cursor  int
	options []string
}

func LeagueOptionsInitialModel() LeagueOptionsModel {
	return LeagueOptionsModel{
		options: data.LeagueOptions,
	}
}

func (m LeagueOptionsModel) Init() tea.Cmd {
	return nil
}

func (m LeagueOptionsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "w", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "a", "j":
			if m.cursor < len(m.options) {
				m.cursor++
			}
		}
	}

	return m, nil
}

func (m LeagueOptionsModel) View() string {
	prompt := "Select Option:\n\n"
	for index, option := range m.options {
		cursor := " "
		if m.cursor == index {
			cursor = ">"
		}
		prompt += cursor + " " + option + "\n"
	}

	return prompt
}

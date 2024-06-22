package menus

import (
	"fmt"

	"github.com/RealLukeMartin/sports-terminal/internal/data"
	tea "github.com/charmbracelet/bubbletea"
)

type LeaguesModel struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func LeaguesInitialModel() LeaguesModel {
	return LeaguesModel{
		choices:  data.Leagues,
		selected: make(map[int]struct{}),
	}
}

func (m LeaguesModel) Init() tea.Cmd {
	return nil
}

func (m LeaguesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Did we get key press?
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "w", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "s", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		}
	}

	return m, nil
}

func (m LeaguesModel) View() string {

	prompt := "Select Sports League:\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		// Render the row
		prompt += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	prompt += "\nPress q to quit.\n"

	return prompt
}

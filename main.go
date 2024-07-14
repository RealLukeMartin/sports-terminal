package main

import (
	"fmt"

	"github.com/RealLukeMartin/sports-terminal/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(tui.TuiInitialModel())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting program:", err)
	}
}

// func main() {
// 	scraper.GetMlbGames()

// }

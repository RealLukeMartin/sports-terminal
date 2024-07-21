package components

import (
	"strconv"

	"github.com/charmbracelet/lipgloss"
	table "github.com/charmbracelet/lipgloss/table"
)

type Inning struct {
	HomeScore int
	AwayScore int
}

type MlbTeamsBoxScoreStats struct {
	Name   string
	Runs   string
	Hits   string
	Errors string
}

func mlbBoxScoreRowParser(stats MlbTeamsBoxScoreStats, innings []string) []string {
	// Add empty innings if less than 9
	// print(len(innings))
	inningsLength := len(innings)
	if inningsLength < 9 {
		for i := inningsLength; i < 9; i++ {
			innings = append(innings, "-")
		}
	}
	combinedStats := append([]string{stats.Name}, innings...)

	combinedStats = append(combinedStats, stats.Runs)
	combinedStats = append(combinedStats, stats.Hits)
	combinedStats = append(combinedStats, stats.Errors)

	return combinedStats
}

func MlbBoxScoreTable(
	innings []Inning,
	homeStats MlbTeamsBoxScoreStats,
	awayStats MlbTeamsBoxScoreStats,
) string {
	var awayInnings []string
	var homeInnings []string

	for _, inning := range innings {
		awayInnings = append(awayInnings, strconv.Itoa(inning.AwayScore))
		homeInnings = append(homeInnings, strconv.Itoa(inning.HomeScore))
	}

	homeRow := mlbBoxScoreRowParser(homeStats, homeInnings)
	awayRow := mlbBoxScoreRowParser(awayStats, awayInnings)

	mlbBoxScoreTable := table.New().
		Border(lipgloss.NormalBorder()).
		Headers(" ", "1", "2", "3", "4", "5", "6", "7", "8", "9", "R", "H", "E").
		Row(awayRow...).
		Row("---").
		Row(homeRow...).
		StyleFunc(func(row, col int) lipgloss.Style {
			var baseStyle lipgloss.Style
			if row == 0 {
				baseStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFCC00"))
			} else {
				baseStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
			}
			// baseStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
			// if col == 9 {
			// 	baseStyle = baseStyle.MarginRight(2).BorderRightBackground(lipgloss.Color("#000000"))
			// }

			return baseStyle.Width(4).Align(lipgloss.Center)

		})

	return mlbBoxScoreTable.Render()
}

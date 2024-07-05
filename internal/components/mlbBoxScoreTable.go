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
	Runs   int
	Hits   int
	Errors int
}

func mlbBoxScoreRowParser(stats MlbTeamsBoxScoreStats, innings []string) []string {
	combinedStats := append([]string{stats.Name}, innings...)

	combinedStats = append(combinedStats, strconv.Itoa(stats.Runs))
	combinedStats = append(combinedStats, strconv.Itoa(stats.Hits))
	combinedStats = append(combinedStats, strconv.Itoa(stats.Errors))

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
		Headers("Team", "1", "2", "3", "4", "5", "6", "7", "8", "9", "R", "H", "E").
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

			return baseStyle.Width(5).Align(lipgloss.Center)

		})

	return mlbBoxScoreTable.Render()
}

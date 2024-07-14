package boxScores

import "github.com/RealLukeMartin/sports-terminal/scraper"

func getMlbTeamStats(statistics []scraper.MlbStatistic) MlbTeamsBoxScoreStats {
	var stats = MlbTeamsBoxScoreStats{}
	for _, stat := range statistics {
		if stat.Abbreviation == "R" {
			stats.Runs = stat.DisplayValue
		}
		if stat.Abbreviation == "H" {
			stats.Hits = stat.DisplayValue
		}
		if stat.Abbreviation == "E" {
			stats.Errors = stat.DisplayValue
		}
	}
	return stats
}

func processMlbTeams(mlbGames *scraper.MlbScoreboardData) (MlbTeamsBoxScoreStats, MlbTeamsBoxScoreStats, []Inning) {
	var firstGame = mlbGames.Events[0]
	var firstCompetition = firstGame.Competitions[0]

	var innings = []Inning{}
	var homeInnings = []int{}
	var awayInnings = []int{}
	var homeStats = MlbTeamsBoxScoreStats{}
	var awayStats = MlbTeamsBoxScoreStats{}

	for _, competitor := range firstCompetition.Competitors {

		if competitor.HomeAway == "home" {
			homeStats.Name = competitor.Team.Abbreviation
			var stats = getMlbTeamStats(competitor.Statistics)

			homeStats.Runs = stats.Runs
			homeStats.Hits = stats.Hits
			homeStats.Errors = stats.Errors

			for _, teamInnings := range competitor.Linescores {
				var singleInningScore = int(teamInnings.Value)
				homeInnings = append(homeInnings, singleInningScore)
			}
		}

		if competitor.HomeAway == "away" {
			awayStats.Name = competitor.Team.Abbreviation
			var stats = getMlbTeamStats(competitor.Statistics)

			awayStats.Runs = stats.Runs
			awayStats.Hits = stats.Hits
			awayStats.Errors = stats.Errors

			for _, teamInnings := range competitor.Linescores {
				var singleInningScore = int(teamInnings.Value)
				awayInnings = append(awayInnings, singleInningScore)
			}
		}
	}

	for i := range homeInnings {
		innings = append(innings, Inning{HomeScore: homeInnings[i], AwayScore: awayInnings[i]})
	}

	return homeStats, awayStats, innings
}

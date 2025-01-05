package riothelper

import "github.com/adamGrgic/riotgosdk/riotmodels"

func GetLeagueEntryIds(leagueEntries []riotmodels.LeagueEntryDto) []string {
	ids := make([]string, 0, len(leagueEntries)) // Initialize a slice with enough capacity
	for _, obj := range leagueEntries {
		ids = append(ids, obj.LeagueId)
	}
	return ids
}

package endpoint

import (
	"fmt"

	"github.com/adamGrgic/riotgosdk/publicconstants/gamemodes"
	"github.com/adamGrgic/riotgosdk/publicconstants/leagues"
)

func GetLeagueEntriesEndpoint(queue gamemodes.GameMode, tier leagues.League, division string, pageNumber int) string {
	endpoint := fmt.Sprintf("/lol/league/v4/entries/%s/%s/%s?page=%o", queue, tier, division, pageNumber)

	return endpoint
}

func GetLeagueEntryEndpoint(leagueId string) string {
	endpoint := fmt.Sprintf("/lol/league/v4/leagues/%s", leagueId)

	return endpoint
}

func GetRecentMatchesEndpoint(puuid string, start string, count string) string {
	endpoint := fmt.Sprintf("/lol/match/v5/matches/by-puuid/%s/ids?start=%s&count=%s", puuid, start, count)

	return endpoint
}

func GetAccountFromPuuidEndpoint(puuid string) string {
	endpoint := fmt.Sprintf("/riot/account/v1/accounts/by-puuid/%s", puuid)

	return endpoint
}

func GetAccountFromGameNameEndpoint(gameName string, tagLine string) string {
	endpoint := fmt.Sprintf("/riot/account/v1/accounts/by-riot-id/%s/%s", gameName, tagLine)

	return endpoint
}

package client

import (
	"fmt"

	"github.com/adamGrgic/riotgosdk/publicconstants/gamemodes"
	"github.com/adamGrgic/riotgosdk/publicconstants/leagues"
)

func GetLeagueEntriesEndpoint(queue gamemodes.GameMode, tier leagues.League, division string, pageNumber int) string {
	endpoint := fmt.Sprintf("/lol/league/v4/entries/%s/%s/%s?page=%o", queue, tier, division, pageNumber)

	return endpoint
}

func GetLeagueManifestEndpoint(leagueId string) string {
	endpoint := fmt.Sprintf("/lol/league/v4/leagues/%s", leagueId)

	return endpoint
}

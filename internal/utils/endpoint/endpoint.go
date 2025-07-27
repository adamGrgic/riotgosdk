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

func GetLeagueEntryFromPuuidEndpoint(puuid string) string {
	endpoint := fmt.Sprintf("/lol/league/v4/entries/by-puuid/%s", puuid)

	return endpoint
}

func GetLeagueEntryFromSummonerIdEndpoint(summonerId string) string {
	endpoint := fmt.Sprintf("/lol/league/v4/entries/by-puuid/%s", summonerId)

	return endpoint
}

func GetChallengerLeagueEntriesFromQueueIdEndpoint(queueId string) string {
	endpoint := fmt.Sprintf("/lol/league/v4/challengerleagues/by-queue/%s", queueId)

	return endpoint
}

func GetGrandmasterLeagueEntriesFromQueueIdEndpoint(queueId string) string {
	endpoint := fmt.Sprintf("/lol/league/v4/grandmasterleagues/by-queue/%s", queueId)

	return endpoint
}

func GetMasterLeagueEntriesFromQueueIdEndpoint(queueId string) string {
	endpoint := fmt.Sprintf("/lol/league/v4/masterleagues/by-queue/%s", queueId)

	return endpoint
}

func GetRecentMatchesEndpoint(puuid string, start string, count string) string {
	endpoint := fmt.Sprintf("/lol/match/v5/matches/by-puuid/%s/ids?start=%s&count=%s", puuid, start, count)

	return endpoint
}

func GetMatchDataEndpoint(matchId string) string {
	endpoint := fmt.Sprintf("/lol/match/v5/matches/%s", matchId)

	return endpoint
}

func GetMatchTimelineDataEndpoint(matchId string) string {
	endpoint := fmt.Sprintf("/lol/match/v5/matches/%s/timeline", matchId)

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

// Summoner API endpoints
func GetSummonerByAccountIdEndpoint(accountId string) string {
	endpoint := fmt.Sprintf("/lol/summoner/v4/summoners/by-account/%s", accountId)
	return endpoint
}

func GetSummonerByNameEndpoint(summonerName string) string {
	endpoint := fmt.Sprintf("/lol/summoner/v4/summoners/by-name/%s", summonerName)
	return endpoint
}

func GetSummonerByPuuidEndpoint(puuid string) string {
	endpoint := fmt.Sprintf("/lol/summoner/v4/summoners/by-puuid/%s", puuid)
	return endpoint
}

func GetSummonerByIdEndpoint(summonerId string) string {
	endpoint := fmt.Sprintf("/lol/summoner/v4/summoners/%s", summonerId)
	return endpoint
}

// Champion Mastery API endpoints
func GetChampionMasteryByPuuidEndpoint(puuid string) string {
	endpoint := fmt.Sprintf("/lol/champion-mastery/v4/champion-masteries/by-puuid/%s", puuid)
	return endpoint
}

func GetChampionMasteryByPuuidAndChampionIdEndpoint(puuid string, championId string) string {
	endpoint := fmt.Sprintf("/lol/champion-mastery/v4/champion-masteries/by-puuid/%s/by-champion/%s", puuid, championId)
	return endpoint
}

func GetChampionMasteryBySummonerIdEndpoint(summonerId string) string {
	endpoint := fmt.Sprintf("/lol/champion-mastery/v4/champion-masteries/by-summoner/%s", summonerId)
	return endpoint
}

func GetChampionMasteryBySummonerIdAndChampionIdEndpoint(summonerId string, championId string) string {
	endpoint := fmt.Sprintf("/lol/champion-mastery/v4/champion-masteries/by-summoner/%s/by-champion/%s", summonerId, championId)
	return endpoint
}

func GetChampionMasteryScoreBySummonerIdEndpoint(summonerId string) string {
	endpoint := fmt.Sprintf("/lol/champion-mastery/v4/scores/by-summoner/%s", summonerId)
	return endpoint
}

// Status API endpoints
func GetStatusEndpoint() string {
	endpoint := "/lol/status/v4/platform-data"
	return endpoint
}

// Champion API endpoints
func GetChampionRotationsEndpoint() string {
	endpoint := "/lol/platform/v3/champion-rotations"
	return endpoint
}

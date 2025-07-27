package riotclient

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"encoding/json"

	"github.com/adamGrgic/riotgosdk/internal/constants/protocol"
	"github.com/adamGrgic/riotgosdk/internal/utils/endpoint"
	"github.com/adamGrgic/riotgosdk/publicconstants/gamemodes"
	"github.com/adamGrgic/riotgosdk/publicconstants/leagues"
	"github.com/adamGrgic/riotgosdk/publicconstants/platformroutes"
	"github.com/adamGrgic/riotgosdk/publicconstants/regionalroutes"
	"github.com/adamGrgic/riotgosdk/riotmodels"
)

// func TestFunction() {
// 	testResponse := utils.FmtTextColor(constants.ApiTestOk, constants.Green)
// 	fmt.Println(testResponse)
// }

func GetLeagueEntry(apiKey string, leagueId string) (*riotmodels.LeagueListDto, error) {
	url := protocol.HTTPS +
		regionalroutes.AMERICAS +
		endpoint.GetLeagueEntryEndpoint(leagueId)

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var leagueResponse riotmodels.LeagueListDto
	if err := json.Unmarshal(body, &leagueResponse); err != nil {
		return nil, err
	}

	return &leagueResponse, nil
}

func GetLeagueEntries(apiKey string, gameMode gamemodes.GameMode, league leagues.League, division string) (*[]riotmodels.LeagueEntryDto, error) {
	url := protocol.HTTPS +
		regionalroutes.AMERICAS +
		endpoint.GetLeagueEntriesEndpoint(gameMode, league, division, 10)

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var leagueResponse []riotmodels.LeagueEntryDto
	if err := json.Unmarshal(body, &leagueResponse); err != nil {
		return nil, err
	}

	return &leagueResponse, nil
}

func GetLeagueEntryFromPuuid(apiKey string, puuid string) (*riotmodels.LeagueEntryDto, error) {
	url := protocol.HTTPS +
		regionalroutes.AMERICAS +
		endpoint.GetLeagueEntryFromPuuidEndpoint(puuid)

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var leagueEntry riotmodels.LeagueEntryDto
	if err := json.Unmarshal(body, &leagueEntry); err != nil {
		return nil, err
	}

	return &leagueEntry, nil
}

func GetLeagueEntryFromSummonerId(apiKey string, summonerId string) (*riotmodels.LeagueEntryDto, error) {
	url := protocol.HTTPS +
		regionalroutes.AMERICAS +
		endpoint.GetLeagueEntryFromSummonerIdEndpoint(summonerId)

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var leagueEntry riotmodels.LeagueEntryDto
	if err := json.Unmarshal(body, &leagueEntry); err != nil {
		return nil, err
	}

	return &leagueEntry, nil
}

func GetChallengerLeagueEntriesFromQueueId(apiKey string, queueId string) (*riotmodels.LeagueListDto, error) {
	url := protocol.HTTPS +
		regionalroutes.AMERICAS +
		endpoint.GetChallengerLeagueEntriesFromQueueIdEndpoint(queueId)

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var leagueList riotmodels.LeagueListDto
	if err := json.Unmarshal(body, &leagueList); err != nil {
		return nil, err
	}

	return &leagueList, nil
}

func GetGrandmasterLeagueEntriesFromQueueId(apiKey string, queueId string) (*[]riotmodels.LeagueEntryDto, error) {
	url := protocol.HTTPS +
		regionalroutes.AMERICAS +
		endpoint.GetGrandmasterLeagueEntriesFromQueueIdEndpoint(queueId)

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var leagueEntries []riotmodels.LeagueEntryDto
	if err := json.Unmarshal(body, &leagueEntries); err != nil {
		return nil, err
	}

	return &leagueEntries, nil
}

func GetMasterLeagueEntriesFromQueueId(apiKey string, queueId string) (*[]riotmodels.LeagueEntryDto, error) {
	url := protocol.HTTPS +
		regionalroutes.AMERICAS +
		endpoint.GetMasterLeagueEntriesFromQueueIdEndpoint(queueId)

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var leagueEntries []riotmodels.LeagueEntryDto
	if err := json.Unmarshal(body, &leagueEntries); err != nil {
		return nil, err
	}

	return &leagueEntries, nil
}

func GetRecentMatches(w http.ResponseWriter, apiKey string, puuid string, start string, count string) (*[]string, error) {
	url := protocol.HTTPS +
		platformroutes.AMERICAS +
		endpoint.GetRecentMatchesEndpoint(puuid, start, count)

	// w.Write([]byte(url))

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var matches []string
	if err := json.Unmarshal(body, &matches); err != nil {
		return nil, err
	}

	return &matches, nil
}

func GetMatchData(w http.ResponseWriter, apiKey string, matchId string) (*riotmodels.MatchDto, error) {
	url := protocol.HTTPS +
		platformroutes.AMERICAS +
		endpoint.GetMatchDataEndpoint(matchId)

	// w.Write([]byte(url))

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var match riotmodels.MatchDto
	if err := json.Unmarshal(body, &match); err != nil {
		return nil, err
	}

	return &match, nil
}

func GetMatchTimelineData(w http.ResponseWriter, apiKey string, matchId string) (*riotmodels.TimelineDto, error) {
	url := protocol.HTTPS +
		platformroutes.AMERICAS +
		endpoint.GetMatchTimelineDataEndpoint(matchId)

	// w.Write([]byte(url))

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var match riotmodels.TimelineDto
	if err := json.Unmarshal(body, &match); err != nil {
		return nil, err
	}

	return &match, nil
}

func GetAccountFromGameName(w http.ResponseWriter, apiKey string, gameName string, tagLine string) (*riotmodels.AccountDto, error) {
	url := protocol.HTTPS +
		platformroutes.AMERICAS +
		endpoint.GetAccountFromGameNameEndpoint(gameName, tagLine)

	// w.Write([]byte(url))

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var account riotmodels.AccountDto
	if err := json.Unmarshal(body, &account); err != nil {
		return nil, err
	}

	return &account, nil
}

func GetAccountFromPuuid(w http.ResponseWriter, apiKey string, puuid string) (*riotmodels.AccountDto, error) {
	url := protocol.HTTPS +
		platformroutes.AMERICAS +
		endpoint.GetAccountFromPuuidEndpoint(puuid)

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var account riotmodels.AccountDto
	if err := json.Unmarshal(body, &account); err != nil {
		return nil, err
	}

	return &account, nil
}

// Summoner API functions
func GetSummonerByAccountId(apiKey string, accountId string) (*riotmodels.SummonerDto, error) {
	url := protocol.HTTPS +
		platformroutes.NA1 +
		endpoint.GetSummonerByAccountIdEndpoint(accountId)

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var summoner riotmodels.SummonerDto
	if err := json.Unmarshal(body, &summoner); err != nil {
		return nil, err
	}

	return &summoner, nil
}

func GetSummonerByName(apiKey string, summonerName string) (*riotmodels.SummonerDto, error) {
	url := protocol.HTTPS +
		platformroutes.NA1 +
		endpoint.GetSummonerByNameEndpoint(summonerName)

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var summoner riotmodels.SummonerDto
	if err := json.Unmarshal(body, &summoner); err != nil {
		return nil, err
	}

	return &summoner, nil
}

func GetSummonerByPuuid(apiKey string, puuid string) (*riotmodels.SummonerDto, error) {
	url := protocol.HTTPS +
		platformroutes.NA1 +
		endpoint.GetSummonerByPuuidEndpoint(puuid)

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var summoner riotmodels.SummonerDto
	if err := json.Unmarshal(body, &summoner); err != nil {
		return nil, err
	}

	return &summoner, nil
}

func GetSummonerById(apiKey string, summonerId string) (*riotmodels.SummonerDto, error) {
	url := protocol.HTTPS +
		platformroutes.NA1 +
		endpoint.GetSummonerByIdEndpoint(summonerId)

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var summoner riotmodels.SummonerDto
	if err := json.Unmarshal(body, &summoner); err != nil {
		return nil, err
	}

	return &summoner, nil
}

// Champion Mastery API functions
func GetChampionMasteryByPuuid(apiKey string, puuid string) (*[]riotmodels.ChampionMasteryDto, error) {
	url := protocol.HTTPS +
		platformroutes.NA1 +
		endpoint.GetChampionMasteryByPuuidEndpoint(puuid)

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var championMasteries []riotmodels.ChampionMasteryDto
	if err := json.Unmarshal(body, &championMasteries); err != nil {
		return nil, err
	}

	return &championMasteries, nil
}

func GetChampionMasteryByPuuidAndChampionId(apiKey string, puuid string, championId string) (*riotmodels.ChampionMasteryDto, error) {
	url := protocol.HTTPS +
		platformroutes.NA1 +
		endpoint.GetChampionMasteryByPuuidAndChampionIdEndpoint(puuid, championId)

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var championMastery riotmodels.ChampionMasteryDto
	if err := json.Unmarshal(body, &championMastery); err != nil {
		return nil, err
	}

	return &championMastery, nil
}

func GetChampionMasteryBySummonerId(apiKey string, summonerId string) (*[]riotmodels.ChampionMasteryDto, error) {
	url := protocol.HTTPS +
		platformroutes.NA1 +
		endpoint.GetChampionMasteryBySummonerIdEndpoint(summonerId)

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var championMasteries []riotmodels.ChampionMasteryDto
	if err := json.Unmarshal(body, &championMasteries); err != nil {
		return nil, err
	}

	return &championMasteries, nil
}

func GetChampionMasteryBySummonerIdAndChampionId(apiKey string, summonerId string, championId string) (*riotmodels.ChampionMasteryDto, error) {
	url := protocol.HTTPS +
		platformroutes.NA1 +
		endpoint.GetChampionMasteryBySummonerIdAndChampionIdEndpoint(summonerId, championId)

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var championMastery riotmodels.ChampionMasteryDto
	if err := json.Unmarshal(body, &championMastery); err != nil {
		return nil, err
	}

	return &championMastery, nil
}

func GetChampionMasteryScoreBySummonerId(apiKey string, summonerId string) (*int, error) {
	url := protocol.HTTPS +
		platformroutes.NA1 +
		endpoint.GetChampionMasteryScoreBySummonerIdEndpoint(summonerId)

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var score int
	if err := json.Unmarshal(body, &score); err != nil {
		return nil, err
	}

	return &score, nil
}

// Status API functions
func GetStatus(apiKey string) (*riotmodels.StatusDto, error) {
	url := protocol.HTTPS +
		platformroutes.NA1 +
		endpoint.GetStatusEndpoint()

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var status riotmodels.StatusDto
	if err := json.Unmarshal(body, &status); err != nil {
		return nil, err
	}

	return &status, nil
}

// Champion API functions
func GetChampionRotations(apiKey string) (*riotmodels.ChampionRotationsDto, error) {
	url := protocol.HTTPS +
		platformroutes.NA1 +
		endpoint.GetChampionRotationsEndpoint()

	resp, err := executeApiRequest(url, apiKey)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var rotations riotmodels.ChampionRotationsDto
	if err := json.Unmarshal(body, &rotations); err != nil {
		return nil, err
	}

	return &rotations, nil
}

// todo: keep this in client.go and move the rest to another file (maybe)
func executeApiRequest(url string, apiKey string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("X-Riot-Token", apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		time.Sleep(120 * time.Second)
		return executeApiRequest(url, apiKey)
	}

	return resp, nil

}

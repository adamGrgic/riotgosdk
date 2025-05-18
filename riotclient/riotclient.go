package riotclient

import (
	"fmt"
	"io"
	"net/http"

	"encoding/json"

	"github.com/adamGrgic/riotgosdk/internal/constants/protocol"
	"github.com/adamGrgic/riotgosdk/internal/utils/endpoint"
	"github.com/adamGrgic/riotgosdk/publicconstants/gamemodes"
	"github.com/adamGrgic/riotgosdk/publicconstants/leagues"
	miscrouting "github.com/adamGrgic/riotgosdk/publicconstants/misc"
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

func GetRecentMatches(apiKey string, puuid string, start string, count string) (*[]string, error) {
	url := protocol.HTTPS +
		miscrouting.Americas +
		endpoint.GetRecentMatchesEndpoint(puuid, start, count)

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

func GetPuuid(apiKey string, gameName string, tagLine string) (*riotmodels.AccountDto, error) {
	url := protocol.HTTPS +
		miscrouting.Americas +
		endpoint.GetPuuidEndpoint(gameName, tagLine)

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

	return resp, nil

}

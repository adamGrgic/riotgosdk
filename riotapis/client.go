package riotapis

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/adamGrgic/riotgosdk/internal/constants"
	"github.com/adamGrgic/riotgosdk/internal/constants/protocol"
	client "github.com/adamGrgic/riotgosdk/internal/utils/clienthelper"
	utils "github.com/adamGrgic/riotgosdk/internal/utils/textformatter"
	"github.com/adamGrgic/riotgosdk/publicconstants/gamemodes"
	"github.com/adamGrgic/riotgosdk/publicconstants/leagues"
	"github.com/adamGrgic/riotgosdk/publicconstants/regionalroutes"
	"github.com/adamGrgic/riotgosdk/riotmodels"
)

func TestFunction() {
	testResponse := utils.FmtTextColor(constants.ApiTestOk, constants.Green)
	fmt.Println(testResponse)
}

func GetLeagueEntriesResponse(apiKey string, gameMode gamemodes.GameMode, league leagues.League, division string) string {
	url := protocol.HTTPS +
		regionalroutes.AMERICAS +
		client.GetLeagueEntriesEndpoint(gameMode, league, division, 10)

	return executeApiRequest(url, apiKey)
}

func GetLeagueEntriesAsDTO(apiKey string, gameMode gamemodes.GameMode, league leagues.League, division string) []riotmodels.LeagueEntryDto {
	url := protocol.HTTPS +
		regionalroutes.AMERICAS +
		client.GetLeagueEntriesEndpoint(gameMode, league, division, 10)

	var leagueEntries []riotmodels.LeagueEntryDto

	response := executeApiRequest(url, apiKey)

	err := json.Unmarshal([]byte(response), &leagueEntries)
	if err != nil {
		log.Println("Error deserializing JSON:", err)
	}

	return leagueEntries
}

func GetLeagueManifest(leagueId string, apiKey string) string {
	url := protocol.HTTPS + regionalroutes.AMERICAS + client.GetLeagueManifestEndpoint(leagueId)
	return executeApiRequest(url, apiKey)
}

func executeApiRequest(url string, apiKey string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	req.Header.Set("X-Riot-Token", apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return ""
	}

	return string(body)
}

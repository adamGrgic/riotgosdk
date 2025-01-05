package riotapis

import (
	"fmt"
	"io"
	"net/http"

	"github.com/adamGrgic/riotgosdk/internal/constants"
	"github.com/adamGrgic/riotgosdk/internal/constants/protocol"
	client "github.com/adamGrgic/riotgosdk/internal/utils/clienthelper"
	utils "github.com/adamGrgic/riotgosdk/internal/utils/textformatter"
	"github.com/adamGrgic/riotgosdk/publicconstants/gamemodes"
	"github.com/adamGrgic/riotgosdk/publicconstants/leagues"
	"github.com/adamGrgic/riotgosdk/publicconstants/regionalroutes"
	"github.com/adamGrgic/riotgosdk/riotmodels"
    "encoding/json"
)

func TestFunction() {
	testResponse := utils.FmtTextColor(constants.ApiTestOk, constants.Green)
	fmt.Println(testResponse)
}

func GetLeague(apiKey string, leagueId string) (*riotmodels.LeagueListDto, error) {
	url := protocol.HTTPS +
        regionalroutes.AMERICAS +
        client.GetLeagueManifestEndpoint(leagueId)

    // fmt.Println("DEBUG RIOTGO URL: %s",url)
    // fmt.Println(url)
    resp, err:= executeApiRequest(url, apiKey)
    if err != nil {
        fmt.Println("Error creating request: ", err)
        return nil,err
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
		client.GetLeagueEntriesEndpoint(gameMode, league, division, 10)

    resp, err:= executeApiRequest(url, apiKey)
    if err != nil {
        fmt.Println("Error creating request: ", err)
        return nil,err
    }


    defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil,err
	}

    var leagueResponse []riotmodels.LeagueEntryDto
    if err := json.Unmarshal(body, &leagueResponse); err != nil {
        return nil,err
    }

    return &leagueResponse, nil

}


// todo: keep this in client.go and move the rest to another file (maybe)
func executeApiRequest(url string, apiKey string, ) (*http.Response, error) {
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

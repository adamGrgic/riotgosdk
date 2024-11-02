package riotapis

import (
	"fmt"
	"io"
	"net/http"

	"github.com/adamGrgic/riotgosdk/internal/constants"
	"github.com/adamGrgic/riotgosdk/internal/constants/protocol"
	"github.com/adamGrgic/riotgosdk/internal/constants/regionalroutes"
	client "github.com/adamGrgic/riotgosdk/internal/utils/clienthelper"
	utils "github.com/adamGrgic/riotgosdk/internal/utils/textformatter"
)

type Client struct {
	APIKey  string
	BaseUrl string
}

func NewClient(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

func TestFunction() {
	testResponse := utils.FmtTextColor(constants.ApiTestOk, constants.Green)
	fmt.Println(testResponse)
}

func GetLeagueEntries(apiKey string) string {
	url := protocol.HTTPS + regionalroutes.AMERICAS + client.GetLeagueEntriesRoute("RANKED_SOLO_5x5", "BRONZE", "III", 10)
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

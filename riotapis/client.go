package riotapis

import (
	"github.com/adamGrgic/riotgosdk/internal/constants"
	"github.com/adamGrgic/riotgosdk/internal/utils"
)

type Client struct {
	APIKey string
}

func NewClient(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

func TestFunction() string {
	return utils.FmtTextColor(constants.ApiTestOk, constants.Green)
}

package riotapis

import "fmt"

type Client struct {
	APIKey string
}

func NewClient(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

func Test() {
	fmt.Println("Riot API SDK Test Ok")
}

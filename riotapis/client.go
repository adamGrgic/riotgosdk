package riotapis

import "fmt"

type Client struct {
	APIKey string
}

func NewClient(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

func Test() {
	fmt.Println("\033[32mRiot API SDK Response Ok\033[0m")
}

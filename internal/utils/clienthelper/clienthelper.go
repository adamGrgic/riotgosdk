package client

import "fmt"

func GetLeagueEntriesRoute(queue string, tier string, division string, pageNumber int) string {
	endpoint := fmt.Sprintf("/lol/league/v4/entries/%s/%s/%s?page=%o", queue, tier, division, pageNumber)

	return endpoint
}

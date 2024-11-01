package riotmodels

type LeagueListDto struct {
	LeagueId string
	Entries  []LeagueItemDto
	Tier     string
	Name     string
	Queue    string
}

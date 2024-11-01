package riotmodels

type TeamDto struct {
	Bans       []BanDto
	Objectives ObjectivesDto
	TeamId     int
	Win        bool
}

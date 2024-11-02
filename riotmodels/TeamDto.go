package riotmodels

type TeamDto struct {
	Bans       []BanDto      `json:"bans"`
	Objectives ObjectivesDto `json:"objectives"`
	TeamId     int           `json:"teamId"`
	Win        bool          `json:"win"`
}

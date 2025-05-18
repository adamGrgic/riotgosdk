package riotmodels

type ChampionMasteryDto struct {
	Puuid                        string                  `json:"puuid"`
	ChampionPointsUntilNextLevel string                  `json:"championPointsUntilNextLevel"`
	ChestGranted                 bool                    `json:"chestGranted"`
	ChampionId                   int                     `json:"championId"`
	LastPlayTime                 int                     `json:"lastPlayTime"`
	ChampionLevel                int                     `json:"championLevel"`
	ChampionPoints               int                     `json:"championPoints"`
	ChampionPointsSinceLastLevel int                     `json:"championPointsSinceLastLevel"`
	MarkRequiredForNextLevel     int                     `json:"markRequiredForNextLevel"`
	ChampionSeasonMilestone      int                     `json:"championSeasonMilestone"`
	NextSeasonMilestone          NextSeasonMilestonesDto `json:"nextSeasonMilestone"`
	TokensEarned                 int                     `json:"tokensEarned"`
	MilestoneGrades              []string                `json:"milestoneGrades"`
}

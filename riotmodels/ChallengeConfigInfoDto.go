package riotmodels

type ChallengeConfigInfoDto struct {
	Id             int                          `json:"id"`
	LocalizedNames map[string]map[string]string `json:"localizedNames"`
	StartTimestamp int                          `json:"startTimestamp"`
	EndTimestamp   int                          `json:"endTimestamp"`
	Leaderboard    bool                         `json:"leaderboard"`
	Thresholds     map[string]float64           `json:"thresholds"`
}

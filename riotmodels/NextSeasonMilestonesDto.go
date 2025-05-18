package riotmodels

type NextSeasonMilestonesDto struct {
	RequireGradeCounts map[string]interface{} `json:"requireGradeCounts"`
	RewardMarks        int                    `json:"rewardMarks"`
	Bonus              bool                   `json:"bonus"`
	RewardConfig       RewardConfigDto        `json:"rewardConfigDto"`
}

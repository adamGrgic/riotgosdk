package riotmodels

type ParticipantFrameDto struct {
	ChampionStats            ChampionStatsDto `json:"championStats"`
	CurrentGold              int              `json:"currentGold"`
	DamageStats              DamageStatsDto   `json:"damageStats"`
	GoldPerSecond            int              `json:"goldPerSecond"`
	JungleMinionsKilled      int              `json:"jungleMinionsKilled"`
	Level                    int              `json:"level"`
	MinionsKilled            int              `json:"minionsKilled"`
	ParticipantId            int              `json:"participantId"`
	Position                 PositionDto      `json:"position"`
	TimeEnemySpentControlled int              `json:"timeEnemySpentControlled"`
	TotalGold                int              `json:"totalGold"`
	Xp                       int              `json:"xp"`
}

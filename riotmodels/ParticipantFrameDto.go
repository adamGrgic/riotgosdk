package riotmodels

type ParticipantFrameDto struct {
	ChampionStats            ChampionStatsDto
	CurrentGold              int
	DamageStats              DamageStatsDto
	GoldPerSecond            int
	JungleMinionsKilled      int
	Level                    int
	MinionsKilled            int
	ParticipantId            int
	Position                 PositionDto
	TimeEnemySpentControlled int
	TotalGold                int
	Xp                       int
}

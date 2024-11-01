package riotmodels

type LeagueEntryDto struct {
	LeagueId     string
	SummonerId   string
	QueueType    string
	Tier         string
	Rank         string
	LeaguePoints int
	Wins         int
	Losses       int
	HotStreak    bool
	Veteran      bool
	FreshBlood   bool
	Inactive     bool
	MiniSeries   MiniSeriesDto
}

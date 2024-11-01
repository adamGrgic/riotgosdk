package riotmodels

type LeagueItemDto struct {
	FreshBlood   bool
	Wins         int
	MiniSeries   MiniSeriesDto
	Inactive     bool
	Veteran      bool
	HotStreak    bool
	Rank         string
	LeaguePoints int
	Losses       int
	SummonerId   string
}

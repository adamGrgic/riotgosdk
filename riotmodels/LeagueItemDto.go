package riotmodels

type LeagueItemDto struct {
	FreshBlood   bool          `json:"freshBlood"`
	Wins         int           `json:"wins"`
	MiniSeries   MiniSeriesDto `json:"miniSeries"`
	Inactive     bool          `json:"inactive"`
	Veteran      bool          `json:"veteran"`
	HotStreak    bool          `json:"hotStreak"`
	Rank         string        `json:"rank"`
	LeaguePoints int           `json:"leaguePoints"`
	Losses       int           `json:"losses"`
	SummonerId   string        `json:"summonerId"`
}

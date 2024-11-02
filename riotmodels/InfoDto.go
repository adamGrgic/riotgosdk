package riotmodels

type InfoDto struct {
	EndOfGameResult    string           `json:"endOfGameResult"`
	GameCreation       int64            `json:"gameCreation"`
	GameDuration       int64            `json:"gameDuration"`
	GameEndTimestamp   int64            `json:"gameEndTimestamp"`
	GameId             int64            `json:"gameId"`
	GameMode           string           `json:"gameMode"`
	GameName           string           `json:"gameName"`
	GameStartTimestamp int64            `json:"gameStartTimestamp"`
	GameType           string           `json:"gameType"`
	GameVersion        string           `json:"gameVersion"`
	MapId              int              `json:"mapId"`
	PlatformId         string           `json:"platformId"`
	Participants       []ParticipantDto `json:"participants"`
	Teams              TeamDto          `json:"teams"`
	QueueId            int              `json:"queueId"`
	TournamentCode     string           `json:"tournamentCode"`
}

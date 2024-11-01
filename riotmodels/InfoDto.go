package riotmodels

type InfoDto struct {
	EndOfGameResult    string
	GameCreation       int64
	GameDuration       int64
	GameEndTimestamp   int64
	GameId             int64
	GameMode           string
	GameName           string
	GameStartTimestamp int64
	GameType           string
	GameVersion        string
	MapId              int
	PlatformId         string
	Participants       []ParticipantDto
	Teams              TeamDto
	QueueId            int
	TournamentCode     string
}

package riotmodels

type InfoTimeLineDto struct {
	EndOfGameResult string
	FrameInterval   int64
	GameId          int64
	Participants    []ParticipantTimeLineDto
	Frames          []FramesTimeLineDto
}

package riotmodels

type InfoTimeLineDto struct {
	EndOfGameResult string                   `json:"endOfGameResult"`
	FrameInterval   int                      `json:"frameInterval"`
	GameId          int                      `json:"gameId"`
	Participants    []ParticipantTimeLineDto `json:"participants"`
	Frames          []FramesTimeLineDto      `json:"frames"`
}

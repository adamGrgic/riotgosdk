package riotmodels

type FramesTimeLineDto struct {
	Events            []EventsTimeLineDto  `json:"events"`
	ParticipantFrames ParticipantFramesDto `json:"particpantFrames"`
	Timestamp         int                  `json:"timestamp"`
}

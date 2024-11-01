package riotmodels

type FramesTimeLineDto struct {
	Events            []EventsTimeLineDto
	ParticipantFrames ParticipantFramesDto
	Timestamp         int
}

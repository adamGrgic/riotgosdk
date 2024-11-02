package riotmodels

type ParticipantFramesDto struct {
	Participants map[string]ParticipantFrameDto `json:"participants"`
}

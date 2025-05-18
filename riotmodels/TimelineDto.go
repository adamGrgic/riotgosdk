package riotmodels

type TimelineDto struct {
	MetaData MetaDataTimeLineDto `json:"metadata"`
	Info     InfoTimeLineDto     `json:"info"`
}

package riotmodels

type EventsTimeLineDto struct {
	Timestamp     int    `json:"timeStamp"`
	RealTimeStamp int    `json:"realTimeStamp"`
	Type          string `json:"type"`
}

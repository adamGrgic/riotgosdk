package riotmodels

type PlatformDataDto struct {
	Id           string      `json:"id"`
	Name         string      `json:"name"`
	Locales      []string    `json:"locales"`
	Maintenances []StatusDto `json:"maintenances"`
	Incidents    []StatusDto `json:"incidents"`
}

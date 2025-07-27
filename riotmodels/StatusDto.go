package riotmodels

type StatusDto struct {
	Id           string           `json:"id"`
	Name         string           `json:"name"`
	Locales      []string         `json:"locales"`
	Maintenances []MaintenanceDto `json:"maintenances"`
	Incidents    []IncidentDto    `json:"incidents"`
}

type MaintenanceDto struct {
	Id                int          `json:"id"`
	MaintenanceStatus string       `json:"maintenance_status"`
	IncidentSeverity  string       `json:"incident_severity"`
	Titles            []ContentDto `json:"titles"`
	Updates           []UpdateDto  `json:"updates"`
	CreatedAt         string       `json:"created_at"`
	ArchiveAt         string       `json:"archive_at"`
	UpdatedAt         string       `json:"updated_at"`
	Platforms         []string     `json:"platforms"`
}

type IncidentDto struct {
	Id               int          `json:"id"`
	IncidentSeverity string       `json:"incident_severity"`
	Titles           []ContentDto `json:"titles"`
	Updates          []UpdateDto  `json:"updates"`
	CreatedAt        string       `json:"created_at"`
	ArchiveAt        string       `json:"archive_at"`
	UpdatedAt        string       `json:"updated_at"`
	Platforms        []string     `json:"platforms"`
}

package riotmodels

type StatusDto struct {
	Id                int          `json:"id"`
	MaintenanceStatus string       `json:"maintenance_status"`
	IncidentSeverity  string       `json:"incident_severity"`
	Titles            []ContentDto `json:"titles"`
	Updates           []ContentDto `json:"updates"`
	CreatedAt         string       `json:"created_at"`
	ArchiveAt         string       `json:"archive_at"`
	UpdatedAt         string       `json:"updated_at"`
	Platforms         []string     `json:"platforms"`
}

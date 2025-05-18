package riotmodels

type UpdateDto struct {
	Id               int          `json:"id"`
	Author           string       `json:"author"`
	Publish          bool         `json:"publish"`
	PublishLocations []string     `json:"publish_locations"`
	Translations     []ContentDto `json:"translations"`
	CreatedAt        string       `json:"created_at"`
	UpdatedAt        string       `json:"updated_at"`
}

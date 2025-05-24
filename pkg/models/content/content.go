package content

import "github.com/jmoiron/sqlx/types"

// ContentItem represents a searchable content item in the system.
type ContentItem struct {
	ID          string          `json:"id" db:"id"`
	Name        string          `json:"name" db:"name"`
	Type        string          `json:"type" db:"type"` // e.g., Movie, Series, Episode
	Description string          `json:"description,omitempty" db:"description"`
	PrimaryImage string        `json:"primaryImage,omitempty" db:"primary_image"`
	MediaSources types.JSONText  `json:"mediaSources" db:"media_sources"` // Stored as JSON in DB
}

// MediaSource represents a source for playing a ContentItem, typically an HLS URL.
type MediaSource struct {
	ID       string `json:"id"`
	Path     string `json:"path"` // HLS URL
	Protocol string `json:"protocol"`
	Container string `json:"container"`
}
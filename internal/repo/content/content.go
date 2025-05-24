package contentrepo


import (
	"database/sql"

	"jellyfin-hls-server/pkg/models/content" // Replace your_module_name
)

// ContentRepo handles database operations for content items and media sources.
type ContentRepo struct {
	DB *sql.DB
}

// NewContentRepo creates a new ContentRepo.
func NewContentRepo(db *sql.DB) *ContentRepo {
	return &ContentRepo{DB: db}
}

// ListContentItems retrieves a list of content items from the database.
func (r *ContentRepo) ListContentItems(parentID, includeItemTypes, fields string) ([]content.ContentItem, error) {
	// Placeholder SQL query - replace with actual query
	rows, err := r.DB.Query(`SELECT id, name, type, description FROM content_items`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []content.ContentItem
	for rows.Next() {
		var item content.ContentItem
		if err := rows.Scan(&item.ID, &item.Name, &item.Type, &item.Description); err != nil {
			return nil, err
		}
		// You'll need to fetch media sources for each item here
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

// GetContentItemByID retrieves a single content item by its ID.
func (r *ContentRepo) GetContentItemByID(id string) (*content.ContentItem, error) {
	// Placeholder SQL query - replace with actual query
	row := r.DB.QueryRow(`SELECT id, name, type, description FROM content_items WHERE id = $1`, id)

	var item content.ContentItem
	if err := row.Scan(&item.ID, &item.Name, &item.Type, &item.Description); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Item not found
		}
		return nil, err
	}

	// Fetch media sources for the item here

	return &item, nil
}

// GetMediaSourcesByItemID retrieves media sources for a content item.
func (r *ContentRepo) GetMediaSourcesByItemID(itemID string) ([]content.MediaSource, error) {
	// Placeholder SQL query - replace with actual query
	rows, err := r.DB.Query(`SELECT id, path, protocol, container FROM media_sources WHERE content_item_id = $1`, itemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sources []content.MediaSource
	for rows.Next() {
		var source content.MediaSource
		if err := rows.Scan(&source.ID, &source.Path, &source.Protocol, &source.Container); err != nil {
			return nil, err
		}
		sources = append(sources, source)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sources, nil
}

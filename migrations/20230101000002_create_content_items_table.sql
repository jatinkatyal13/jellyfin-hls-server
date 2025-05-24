-- migrate:up
CREATE TABLE content_items (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL,
    description TEXT
);

CREATE TABLE media_sources (
    id UUID PRIMARY KEY,
    content_item_id UUID NOT NULL REFERENCES content_items(id) ON DELETE CASCADE,
    path TEXT NOT NULL,
    protocol VARCHAR(50) NOT NULL,
    container VARCHAR(50) NOT NULL
);

-- migrate:down
DROP TABLE media_sources;
DROP TABLE content_items;
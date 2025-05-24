-- migrate:up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    primary_image VARCHAR(255)
);

-- migrate:down
DROP TABLE users;
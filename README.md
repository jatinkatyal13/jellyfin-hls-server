## Application Overview

This application is a Go-based service that manages searchable content items. It utilizes a PostgreSQL database for storage and provides an API for interacting with content and user sessions.

## Key Components

*   **`internal/config/config.go`**:  Handles application configuration, loading settings from a `.env` file or environment variables using Viper.
*   **`internal/db/postgres.go`**:  Manages the PostgreSQL database connection.
*   **`internal/api/`**: Contains API handler functions for authentication, content management, and user sessions.
*   **`internal/repo/`**: Implements data access logic for content and users.
*   **`pkg/models/`**: Defines data models for content items and users.

## Data Models

*   **`ContentItem`**: Represents a searchable content item with fields like ID, Name, Type (e.g., Movie, Series, Episode), Description, PrimaryImage, and MediaSources (stored as JSON).
*   **`MediaSource`**: Defines a source for playing content, including ID, Path (HLS URL), Protocol, and Container.

## Configuration

The application uses a `.env` file to manage configuration settings.  The following environment variables are supported:

*   `DATABASE_URL`: The connection string for the PostgreSQL database.
*   `SERVER_PORT`: The port on which the server will listen.

## Getting Started

1.  **Environment Setup**: Customize your environment by modifying the `.idx/dev.nix` file to include the necessary tools and IDE extensions.
2.  **Configuration**: Create a `.env` file in the project root and set the required environment variables.
3.  **Database Migrations**: Run the database migrations located in the `migrations/` directory to set up the database schema.
4.  **Run the Application**: Execute the `main.go` file to start the application.

## Next steps

* Explore the [Firebase Studio documentation](/docs/studio).
* [Get started with Firebase Studio](https://studio.firebase.google.com/).

Send feedback
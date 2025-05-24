package usersrepo

import (
	"database/sql"

	"jellyfin-hls-server/pkg/models/users" // Replace your_module_name with your actual module name
)

// UserRepo handles database interactions for users.
type UserRepo struct {
	db *sql.DB
}

// NewUserRepo creates a new UserRepo.
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

// CreateUser inserts a new user into the database.
func (r *UserRepo) CreateUser(user *users.User) error {
	// Placeholder SQL
	query := `INSERT INTO users (id, name, primary_image) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, user.ID, user.Name, user.PrimaryImage)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByID retrieves a user by their ID.
func (r *UserRepo) GetUserByID(id string) (*users.User, error) {
	// Placeholder SQL
	query := `SELECT id, name, primary_image FROM users WHERE id = $1`
	row := r.db.QueryRow(query, id)

	user := &users.User{}
	err := row.Scan(&user.ID, &user.Name, &user.PrimaryImage)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}
	return user, nil
}

// Other user-related database functions would go here

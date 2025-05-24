package users

type User struct {
	ID           string `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	PrimaryImage string `json:"primaryImage,omitempty" db:"primary_image"`
}
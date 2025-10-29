package models

type User struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	LastName string   `json:"last_name"`
	Email    string   `json:"email"`
	Role     []string `json:"role"`
}

type UserCreate struct {
	Name     string   `json:"name"`
	LastName string   `json:"last_name"`
	Email    string   `json:"email"`
	Role     []string `json:"role"`
}

type UserUpdate struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	LastName string   `json:"last_name"`
	Email    string   `json:"email"`
	Role     []string `json:"role"`
}

type UserID struct {
	ID string `json:"id"`
}

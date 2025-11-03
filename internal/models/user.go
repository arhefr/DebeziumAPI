package models

type User struct {
	ID       int64    `json:"id"`
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
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	LastName string   `json:"last_name"`
	Email    string   `json:"email"`
	Role     []string `json:"role"`
}

type UserID struct {
	ID int64 `json:"id"`
}

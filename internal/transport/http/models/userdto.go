package models

type UsersDTO struct {
	Users []UserDTO `json:"users"`
	Total int       `json:"total"`
}

// Get /api/v1/users?limit=100&offset=1

type UserDTO struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

// DTO - data transfer object

// 1000000000 пользователей 100byte
// 1 сто из 10 млн
// offset, limit
// limit - 100
// offset - 1
// offset*limit   1*100 пользователей, 5*100 пользователей

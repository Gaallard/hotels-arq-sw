package dto

type UserDto struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

type UsersDto []UserDto

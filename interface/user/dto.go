package user

import "simple-ddd/domain/user"

type CreateUserRequest struct {
	Name  string `json:"name" example:"John Doe"`
	Email string `json:"email" example:"test@email.com"`
}

type CreateUserResponse struct {
	ID    int    `json:"id" example:"1"`
	Name  string `json:"name" example:"John Doe"`
	Email string `json:"email" example:"test@email.com"`
}

type GetUsersResponse struct {
	Data    []user.User
	Message string `json:"message" example:"ok"`
}

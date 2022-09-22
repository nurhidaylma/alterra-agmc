package public

import (
	"github.com/nurhidaylma/alterra-agmc/day-7/internal/repository"
)

type UserResponse struct {
	ID       uint                   `json:"id"`
	Email    string                 `json:"email"`
	Username string                 `json:"username"`
	Fullname string                 `json:"full_name"`
	Gender   repository.GenderTypes `json:"gender"`
	Age      int                    `json:"age"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	ID       uint                   `param:"id" validate:"required"`
	Email    string                 `json:"email"`
	Password string                 `json:"password"`
	Username string                 `json:"username"`
	Fullname string                 `json:"full_name"`
	Gender   repository.GenderTypes `json:"gender"`
	Age      int                    `json:"age"`
}

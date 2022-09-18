package public

import (
	"github.com/nurhidaylma/alterra-agmc/day-2/models"
)

type UserResponse struct {
	ID       uint               `json:"id"`
	Email    string             `json:"email"`
	Username string             `json:"username"`
	Fullname string             `json:"full_name"`
	Gender   models.GenderTypes `json:"gender"`
	Age      int                `json:"age"`
}

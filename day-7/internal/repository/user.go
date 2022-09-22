package repository

import (
	"time"
)

type User struct {
	ID        uint        `json:"id" gorm:"primaryKey,not null" validate:"required"`
	Email     string      `json:"email" gorm:"not null,unique" validate:"required"`
	Username  string      `json:"username" gorm:"not null" validate:"required"`
	Password  string      `json:"password" gorm:"not null" validate:"required"`
	FullName  string      `json:"full_name" gorm:"not null" validate:"required"`
	Gender    GenderTypes `json:"gender" gorm:"not null" validate:"required"`
	Age       int         `json:"age" gorm:"not null" validate:"required"`
	Token     string      `json:"token"`
	CreatedAt time.Time   `json:"created_at" gorm:"not null,autoCreateTime"`
	UpdatedAt time.Time   `json:"updated_at" gorm:"not null,autoUpdateTime"`
	DeletedAt *time.Time  `json:"deleted_at" sql:"index"`
}

// Gender types
type GenderTypes string

const (
	Female GenderTypes = "female"
	Male   GenderTypes = "male"
)

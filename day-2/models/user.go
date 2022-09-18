package models

import (
	"time"
)

type User struct {
	ID        uint        `json:"id" gorm:"primaryKey,not null"`
	Email     string      `json:"email" gorm:"not null,unique"`
	Username  string      `json:"username" gorm:"not null"`
	Password  string      `json:"password" gorm:"not null"`
	FullName  string      `json:"full_name" gorm:"not null"`
	Gender    GenderTypes `json:"gender" gorm:"not null"`
	Age       int         `json:"age" gorm:"not null"`
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

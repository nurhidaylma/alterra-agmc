package models

import "time"

type Book struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	Title     string     `json:"title" gorm:"not null"`
	Writer    string     `json:"writer" gorm:"not null"`
	NoOfPage  int        `json:"no_of_page" gorm:"not null"`
	Stock     int        `json:"stock" gorm:"not null"`
	Price     float32    `json:"price" gorm:"not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"not null,autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"not null,autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

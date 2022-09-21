package database

import (
	"math/rand"

	"github.com/nurhidaylma/alterra-agmc/day-4/config"
	"github.com/nurhidaylma/alterra-agmc/day-4/models"
	"gorm.io/gorm"
)

type seed struct {
	DB *gorm.DB
}

func NewSeeder() *seed {
	return &seed{config.GetConnection()}
}

func (s *seed) Seed() error {
	seeders := seedUp()
	for _, seed := range seeders {
		err := s.DB.Create(seed).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func seedUp() []interface{} {
	result := []interface{}{
		&models.User{
			ID:       uint(rand.Int()),
			Email:    "random@mail.com",
			Username: "randdom",
			Password: "ahdh8881ahQjalm0!",
			FullName: "Joice Tyler",
			Gender:   models.Female,
			Age:      20,
		},
	}

	return result
}

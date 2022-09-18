package database

import (
	"github.com/nurhidaylma/alterra-agmc/day-2/config"
	"github.com/nurhidaylma/alterra-agmc/day-2/models"
	"github.com/nurhidaylma/alterra-agmc/day-2/public"
)

func GetUsers() ([]public.UserResponse, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	result := []public.UserResponse{}
	for _, user := range users {
		result = append(result, public.UserResponse{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
			Fullname: user.FullName,
			Gender:   user.Gender,
			Age:      user.Age,
		})
	}

	return result, nil
}

func GetUserByID(userID uint) (*public.UserResponse, error) {
	user := models.User{}
	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	return &public.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Fullname: user.FullName,
		Gender:   user.Gender,
		Age:      user.Age,
	}, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(user *models.User) (*public.UserResponse, error) {
	if err := config.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return &public.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Fullname: user.FullName,
		Gender:   user.Gender,
		Age:      user.Age,
	}, nil
}

func UpdateUser(user *models.User) (*public.UserResponse, error) {
	if err := config.DB.Model(user).Updates(user).Error; err != nil {
		return nil, err
	}
	if err := config.DB.First(&user, "id = ?", user.ID).Error; err != nil {
		return nil, err
	}

	return &public.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Fullname: user.FullName,
		Gender:   user.Gender,
		Age:      user.Age,
	}, nil
}

func DeleteUser(id uint) error {
	user := models.User{}

	if err := config.DB.Delete(&user, id).Error; err != nil {
		return err
	}

	return nil
}

package mysql

import (
	"errors"

	"github.com/nurhidaylma/alterra-agmc/day-7/config"
	"github.com/nurhidaylma/alterra-agmc/day-7/internal/public"
	"github.com/nurhidaylma/alterra-agmc/day-7/internal/repository"
	"github.com/nurhidaylma/alterra-agmc/day-7/middlewares"
	"gorm.io/gorm"
)

func LoginUser(user *repository.User) (*string, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return &user.Token, nil
}

func GetUsers() ([]public.UserResponse, error) {
	var users []repository.User

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
	user := repository.User{}

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

func GetUserByEmail(email string) (*repository.User, error) {
	user := repository.User{}
	if err := config.DB.Where("email = ?", email).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, nil
}

func CreateUser(user *repository.User) (*public.UserResponse, error) {
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

func UpdateUser(user *repository.User) (*public.UserResponse, error) {
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
	user := repository.User{}

	if err := config.DB.Delete(&user, id).Error; err != nil {
		return err
	}

	return nil
}

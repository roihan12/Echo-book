package services

import (
	"echo-book/config"
	"echo-book/model"
	"errors"
)

func GetAllusers() ([]model.User, error) {

	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {
		return []model.User{}, err
	}

	return users, nil
}

func GetUserByID(id string) (model.User, error) {
	var user model.User = model.User{}

	if err := config.DB.Find(&user, "id = ?", id).Error; err != nil {
		return model.User{}, err
	}

	if user.ID == 0 {
		return model.User{}, errors.New("user not found")
	}

	return user, nil
}

func CreateUser(input model.User) (model.User, error) {
	var user model.User = model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil

}

func UpdateUser(input model.User, id string) (model.User, error) {
	user, err := GetUserByID(id)

	if err != nil {
		return model.User{}, nil
	}

	user.Email = input.Email
	user.Name = input.Name
	user.Password = input.Password

	if err := config.DB.Save(&user).Error; err != nil {
		return model.User{}, nil
	}

	return user, nil
}

func DeleteUser(id string) bool {
	user, err := GetUserByID(id)

	if err != nil {
		return false
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return false
	}

	return true
}

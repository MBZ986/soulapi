package services

import (
	"github.com/davecgh/go-spew/spew"
	"soulapi/global"
	"soulapi/models"
)

type UserService struct {
}

func (s UserService) QueryAllUsers() ([]models.User, error) {
	var users []models.User
	res := global.DB.Find(&users)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_user)
		return nil, res.Error
	}
	spew.Dump(users)
	return users, nil
}

func (s UserService) QueryUserById(id uint) (*models.User, error) {
	var user models.User
	res := global.DB.First(user, id)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_user)
		return nil, res.Error
	}

	return &user, nil
}

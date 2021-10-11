package services

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"soulapi/global"
	"soulapi/models"
)

type UserService struct {
	BaseService
}

func (s UserService) QueryAllUsers() ([]models.User, error) {
	var users []models.User
	res := global.DB.Find(&users)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_user)
		return nil, res.Error
	}
	return users, nil
}

func (s UserService) QueryUsersByPage(offset, limit int) (*models.Page, error) {
	var users []models.User
	res := global.DB.Offset(offset).Limit(limit).Find(&users)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败：%v", models_user, res.Error)
		return nil, res.Error
	}
	return s.pageHandler(users, offset, limit, s.CountUsers)
}

func (s UserService) FindUsers(user models.User, offset, limit int) (*models.Page, error) {
	var users []models.User
	res := global.DB.Where(&user).Offset(offset).Limit(limit).Find(&users)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_user)
		return nil, res.Error
	}

	return s.pageHandler(users, offset, limit, s.CountUsers)
}

func (s UserService) QueryById(id uint) (*models.User, error) {
	var user models.User
	res := global.DB.First(&user, id)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_user)
		return nil, res.Error
	}
	spew.Dump(user)
	return &user, nil
}
func (s UserService) DeleteById(id uint) error {
	var user models.User
	res := global.DB.Delete(&user, id)
	if res.Error != nil {
		global.Logger.Errorf("删除%s失败", models_user)
		return res.Error
	}
	return nil
}
func (s UserService) CreateUser(user models.User) (uint, error) {
	var tmp models.Title
	if err := global.DB.Where("username = ?", user.Username).First(&tmp).Error; err != nil {
		global.Logger.Errorf("%s名称重复,创建失败", models_user)
		return 0, err
	} else {
		if &tmp != nil {
			return 0, fmt.Errorf("%s名称重复,创建失败", models_music)
		}
	}
	res := global.DB.Where("username = ?", user.Username).Attrs(&user).FirstOrCreate(&user)
	if res.Error != nil {
		global.Logger.Errorf("创建%s失败", models_user)
		return 0, res.Error
	}
	return user.Id, nil
}
func (s UserService) UpdateUser(user models.User) (uint, error) {
	res := global.DB.Updates(user)
	if res.Error != nil {
		global.Logger.Errorf("更新%s失败", models_user)
		return 0, res.Error
	}
	return user.Id, nil
}

func (s UserService) CountUsers() (count int64, err error) {
	res := global.DB.Find(&models.User{}).Count(&count)
	if res.Error != nil {
		global.Logger.Errorf("创建%s失败", models_user)
		return count, res.Error
	}
	return count, nil
}

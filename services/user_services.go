package services

import (
	"fmt"
	"gorm.io/gorm"
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
	var tmp models.User
	if err := global.DB.Where("username = ?", user.Username).First(&tmp).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			global.Logger.Errorf("校验%s失败：%v", models_user, err)
			return 0, err
		}
	} else {
		return 0, fmt.Errorf("%s名称重复,创建失败", models_user)
	}

	if user.SoulTitleId == 0 {
		user.SoulTitleId = 1
	}

	res := global.DB.Create(&user)
	if res.Error != nil {
		global.Logger.Errorf("创建%s失败", models_user)
		return 0, res.Error
	}
	return user.Id, nil
}
func (s UserService) UpdateUser(user models.User) (uint, error) {
	res := global.DB.Updates(&user)
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
func (s UserService) HasUser(id uint) (has bool) {
	if user, err := s.QueryById(id); err != nil {
		global.Logger.Errorf("查询%s失败：%v", models_title, err)
		return false
	} else {
		if user != nil {
			return true
		}
	}
	return false
}

//添加头衔
func (s UserService) AddTitle(userId, titleId uint) (err error) {
	var (
		user  *models.User
		title *models.Title
	)
	if user, _ = s.QueryById(userId); user == nil {
		return fmt.Errorf("用户为空")
	}
	if title, _ = NewTitleService().QueryById(titleId); title == nil {
		return fmt.Errorf("头衔为空")
	}
	user.Titles = []*models.Title{title}
	if err = global.DB.Debug().Updates(user).Error; err != nil {
		global.Logger.Errorf("添加用户头衔失败:%v", err)
		return err
	}
	return nil
}

//关注
func (s UserService) Follow(user models.User, follower models.User) (err error) {
	if err = global.DB.Model(&user).Association("Friends").Append(&follower); err != nil {
		global.Logger.Errorf("添加用户头衔失败:%v", err)
		return err
	}
	return nil
}

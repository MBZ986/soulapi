package services

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"soulapi/global"
	"soulapi/models"
)

type MessageService struct {
	BaseService
}

func (s MessageService) QueryAllMessages() ([]models.Message, error) {
	var messages []models.Message
	res := global.DB.Find(&messages)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_message)
		return nil, res.Error
	}
	return messages, nil
}

func (s MessageService) QueryMessagesByPage(offset, limit int) (*models.Page, error) {
	var messages []models.Message
	res := global.DB.Offset(offset).Limit(limit).Find(&messages)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败：%v", models_message, res.Error)
		return nil, res.Error
	}
	return s.pageHandler(messages, offset, limit, s.CountMessages)
}

func (s MessageService) FindMessages(message models.Message, offset, limit int) (*models.Page, error) {
	var messages []models.Message
	res := global.DB.Where(&message).Offset(offset).Limit(limit).Find(&messages)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_message)
		return nil, res.Error
	}

	return s.pageHandler(messages, offset, limit, s.CountMessages)
}

func (s MessageService) QueryById(id uint) (*models.Message, error) {
	var message models.Message
	res := global.DB.First(&message, id)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_message)
		return nil, res.Error
	}
	spew.Dump(message)
	return &message, nil
}
func (s MessageService) DeleteById(id uint) error {
	var message models.Message
	res := global.DB.Delete(&message, id)
	if res.Error != nil {
		global.Logger.Errorf("删除%s失败", models_message)
		return res.Error
	}
	return nil
}
func (s MessageService) CreateMessage(message models.Message) (uint, error) {
	var tmp models.Title
	if err := global.DB.Where("title = ?", message.Title).First(&tmp).Error; err != nil {
		global.Logger.Errorf("%s标题重复,创建失败", models_message)
		return 0, err
	} else {
		if &tmp != nil {
			return 0, fmt.Errorf("%s名称重复,创建失败", models_music)
		}
	}
	res := global.DB.Where("title = ?", message.Title).Attrs(&message).FirstOrCreate(&message)
	if res.Error != nil {
		global.Logger.Errorf("创建%s失败", models_message)
		return 0, res.Error
	}
	return message.Id, nil
}
func (s MessageService) UpdateMessage(message models.Message) (uint, error) {
	res := global.DB.Updates(message)
	if res.Error != nil {
		global.Logger.Errorf("更新%s失败", models_message)
		return 0, res.Error
	}
	return message.Id, nil
}

func (s MessageService) CountMessages() (count int64, err error) {
	res := global.DB.Find(&models.Message{}).Count(&count)
	if res.Error != nil {
		global.Logger.Errorf("创建%s失败", models_message)
		return count, res.Error
	}
	return count, nil
}

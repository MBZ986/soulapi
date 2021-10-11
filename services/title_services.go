package services

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"soulapi/global"
	"soulapi/models"
)

type TitleService struct {
	BaseService
}

func (s TitleService) QueryAllTitles() ([]models.Title, error) {
	var titles []models.Title
	res := global.DB.Find(&titles)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_title)
		return nil, res.Error
	}
	return titles, nil
}

func (s TitleService) QueryTitlesByPage(offset, limit int) (*models.Page, error) {
	var titles []models.Title
	res := global.DB.Offset(offset).Limit(limit).Find(&titles)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败：%v", models_title, res.Error)
		return nil, res.Error
	}
	return s.pageHandler(titles, offset, limit, s.CountTitles)
}

//查询使用该头衔的用户
func (s TitleService) QueryUsersByTitleId(id uint) (*models.Title, error) {
	var title models.Title
	res := global.DB.Preload("Users").First(&title, id)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_title)
		return nil, res.Error
	}
	spew.Dump(title)
	return &title, nil
}

func (s TitleService) FindTitles(title models.Title, offset, limit int) (*models.Page, error) {
	var titles []models.Title
	res := global.DB.Where(&title).Offset(offset).Limit(limit).Find(&titles)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_title)
		return nil, res.Error
	}

	return s.pageHandler(titles, offset, limit, s.CountTitles)
}

func (s TitleService) QueryById(id uint) (*models.Title, error) {
	var title models.Title
	res := global.DB.First(&title, id)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_title)
		return nil, res.Error
	}
	spew.Dump(title)
	return &title, nil
}
func (s TitleService) DeleteById(id uint) error {
	var title models.Title
	res := global.DB.Delete(&title, id)
	if res.Error != nil {
		global.Logger.Errorf("删除%s失败", models_title)
		return res.Error
	}
	return nil
}
func (s TitleService) CreateTitle(title models.Title) (uint, error) {
	var tmp models.Title
	if err := global.DB.Where("title = ?", title.Title).First(&tmp).Error; err != nil {
		global.Logger.Errorf("%s名称重复,创建失败", models_title)
		return 0, err
	} else {
		if &tmp != nil {
			return 0, fmt.Errorf("%s名称重复,创建失败", models_music)
		}
	}
	res := global.DB.Where("title = ?", title.Title).Attrs(&title).FirstOrCreate(&title)
	if res.Error != nil {
		global.Logger.Errorf("创建%s失败", models_title)
		return 0, res.Error
	}
	return title.Id, nil
}
func (s TitleService) UpdateTitle(title models.Title) (uint, error) {
	res := global.DB.Updates(title)
	if res.Error != nil {
		global.Logger.Errorf("更新%s失败", models_title)
		return 0, res.Error
	}
	return title.Id, nil
}

func (s TitleService) CountTitles() (count int64, err error) {
	res := global.DB.Find(&models.Title{}).Count(&count)
	if res.Error != nil {
		global.Logger.Errorf("创建%s失败", models_title)
		return count, res.Error
	}
	return count, nil
}

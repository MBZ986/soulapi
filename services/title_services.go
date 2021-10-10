package services

import (
	"github.com/davecgh/go-spew/spew"
	"soulapi/global"
	"soulapi/models"
)

type TitleService struct {
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
	return s.pageHandler(titles, offset, limit)
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

	return s.pageHandler(titles, offset, limit)
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

func (s TitleService) pageHandler(data interface{}, offset, limit int) (*models.Page, error) {
	var (
		page models.Page
		err  error
	)
	page.Offset = offset
	page.Limit = limit
	page.Data = data
	if page.Total, err = s.CountTitles(); err != nil {
		global.Logger.Errorf("查询%s失败：%v", models_title, err)
		return nil, err
	}
	return &page, nil
}

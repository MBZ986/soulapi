package services

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"soulapi/global"
	"soulapi/models"
)

type LabelService struct {
	BaseService
}

func (s LabelService) QueryAllLabels() ([]models.Label, error) {
	var labels []models.Label
	res := global.DB.Find(&labels)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_label)
		return nil, res.Error
	}
	return labels, nil
}

func (s LabelService) QueryLabelsByPage(offset, limit int) (*models.Page, error) {
	var labels []models.Label
	res := global.DB.Offset(offset).Limit(limit).Find(&labels)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败：%v", models_label, res.Error)
		return nil, res.Error
	}
	return s.pageHandler(labels, offset, limit, s.CountLabels)
}

//查询使用该头衔的用户
func (s LabelService) QueryUsersByLabelId(id uint) (*models.Label, error) {
	var label models.Label
	res := global.DB.Preload("Users").First(&label, id)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_label)
		return nil, res.Error
	}
	spew.Dump(label)
	return &label, nil
}

func (s LabelService) FindLabels(label models.Label, offset, limit int) (*models.Page, error) {
	var labels []models.Label
	res := global.DB.Where(&label).Offset(offset).Limit(limit).Find(&labels)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_label)
		return nil, res.Error
	}

	return s.pageHandler(labels, offset, limit, s.CountLabels)
}

func (s LabelService) QueryById(id uint) (*models.Label, error) {
	var label models.Label
	res := global.DB.First(&label, id)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_label)
		return nil, res.Error
	}
	spew.Dump(label)
	return &label, nil
}
func (s LabelService) DeleteById(id uint) error {
	var label models.Label
	res := global.DB.Delete(&label, id)
	if res.Error != nil {
		global.Logger.Errorf("删除%s失败", models_label)
		return res.Error
	}
	return nil
}
func (s LabelService) CreateLabel(label models.Label) (uint, error) {
	var tmp models.Label
	if err := global.DB.Where("label_name = ?", label.LabelName).First(&tmp).Error; err != nil {
		global.Logger.Errorf("%s名称重复,创建失败", models_label)
		return 0, err
	} else {
		if &tmp != nil {
			return 0, fmt.Errorf("%s名称重复,创建失败", models_music)
		}
	}
	res := global.DB.Where("label_name = ?", label.LabelName).Attrs(&label).FirstOrCreate(&label)
	if res.Error != nil {
		global.Logger.Errorf("创建%s失败", models_label)
		return 0, res.Error
	}
	return label.Id, nil
}
func (s LabelService) UpdateLabel(label models.Label) (uint, error) {
	res := global.DB.Updates(label)
	if res.Error != nil {
		global.Logger.Errorf("更新%s失败", models_label)
		return 0, res.Error
	}
	return label.Id, nil
}

func (s LabelService) CountLabels() (count int64, err error) {
	res := global.DB.Find(&models.Label{}).Count(&count)
	if res.Error != nil {
		global.Logger.Errorf("创建%s失败", models_label)
		return count, res.Error
	}
	return count, nil
}

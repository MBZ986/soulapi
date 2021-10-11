package services

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"soulapi/global"
	"soulapi/models"
)

type PartitionTypeService struct {
	BaseService
}

func (s PartitionTypeService) QueryAllPartitionTypes() ([]models.PartitionType, error) {
	var partitionTypes []models.PartitionType
	res := global.DB.Find(&partitionTypes)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_partition_type)
		return nil, res.Error
	}
	return partitionTypes, nil
}

func (s PartitionTypeService) QueryPartitionTypesByPage(offset, limit int) (*models.Page, error) {
	var partitionTypes []models.PartitionType
	res := global.DB.Offset(offset).Limit(limit).Find(&partitionTypes)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败：%v", models_partition_type, res.Error)
		return nil, res.Error
	}
	return s.pageHandler(partitionTypes, offset, limit, s.CountPartitionTypes)
}

func (s PartitionTypeService) FindPartitionTypes(partitionType models.PartitionType, offset, limit int) (*models.Page, error) {
	var partitionTypes []models.PartitionType
	res := global.DB.Where(&partitionType).Offset(offset).Limit(limit).Find(&partitionTypes)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_partition_type)
		return nil, res.Error
	}

	return s.pageHandler(partitionTypes, offset, limit, s.CountPartitionTypes)
}

func (s PartitionTypeService) QueryById(id uint) (*models.PartitionType, error) {
	var partitionType models.PartitionType
	res := global.DB.First(&partitionType, id)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_partition_type)
		return nil, res.Error
	}
	spew.Dump(partitionType)
	return &partitionType, nil
}
func (s PartitionTypeService) DeleteById(id uint) error {
	var partitionType models.PartitionType
	res := global.DB.Delete(&partitionType, id)
	if res.Error != nil {
		global.Logger.Errorf("删除%s失败", models_partition_type)
		return res.Error
	}
	return nil
}
func (s PartitionTypeService) CreatePartitionType(partitionType models.PartitionType) (uint, error) {
	var tmp models.Title
	if err := global.DB.Where("part_type = ?", partitionType.PartType).First(&tmp).Error; err != nil {
		global.Logger.Errorf("%s名称重复,创建失败", models_partition_type)
		return 0, err
	} else {
		if &tmp != nil {
			return 0, fmt.Errorf("%s名称重复,创建失败", models_music)
		}
	}
	res := global.DB.Where("part_type = ?", partitionType.PartType).Attrs(&partitionType).FirstOrCreate(&partitionType)
	if res.Error != nil {
		global.Logger.Errorf("创建%s失败", models_partition_type)
		return 0, res.Error
	}
	return partitionType.Id, nil
}
func (s PartitionTypeService) UpdatePartitionType(partitionType models.PartitionType) (uint, error) {
	res := global.DB.Updates(partitionType)
	if res.Error != nil {
		global.Logger.Errorf("更新%s失败", models_partition_type)
		return 0, res.Error
	}
	return partitionType.Id, nil
}

func (s PartitionTypeService) CountPartitionTypes() (count int64, err error) {
	res := global.DB.Find(&models.PartitionType{}).Count(&count)
	if res.Error != nil {
		global.Logger.Errorf("统计%s失败", models_partition_type)
		return count, res.Error
	}
	return count, nil
}

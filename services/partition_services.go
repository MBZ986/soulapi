package services

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"soulapi/global"
	"soulapi/models"
)

type PartitionService struct {
	BaseService
}

func (s PartitionService) QueryAllPartitions() ([]models.Partition, error) {
	var partitions []models.Partition
	res := global.DB.Find(&partitions)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_partition)
		return nil, res.Error
	}
	return partitions, nil
}

func (s PartitionService) QueryPartitionsByPage(offset, limit int) (*models.Page, error) {
	var partitions []models.Partition
	res := global.DB.Offset(offset).Limit(limit).Find(&partitions)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败：%v", models_partition, res.Error)
		return nil, res.Error
	}
	return s.pageHandler(partitions, offset, limit, s.CountPartitions)
}

func (s PartitionService) FindPartitions(partition models.Partition, offset, limit int) (*models.Page, error) {
	var partitions []models.Partition
	res := global.DB.Where(&partition).Offset(offset).Limit(limit).Find(&partitions)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_partition)
		return nil, res.Error
	}

	return s.pageHandler(partitions, offset, limit, s.CountPartitions)
}

func (s PartitionService) QueryById(id uint) (*models.Partition, error) {
	var partition models.Partition
	res := global.DB.First(&partition, id)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_partition)
		return nil, res.Error
	}
	spew.Dump(partition)
	return &partition, nil
}
func (s PartitionService) DeleteById(id uint) error {
	var partition models.Partition
	res := global.DB.Delete(&partition, id)
	if res.Error != nil {
		global.Logger.Errorf("删除%s失败", models_partition)
		return res.Error
	}
	return nil
}
func (s PartitionService) CreatePartition(partition models.Partition) (uint, error) {
	var tmp models.Partition
	if err := global.DB.Where("part_title = ? && part_type_id = ?", partition.PartTitle, partition.PartTypeId).First(&tmp).Error; err != nil {
		global.Logger.Errorf("%s名称重复,创建失败", models_partition)
		return 0, err
	} else {
		if &tmp != nil {
			return 0, fmt.Errorf("%s重复,创建失败", models_partition)
		}
	}
	res := global.DB.Where("part_title = ? && part_type_id = ?", partition.PartTitle, partition.PartTypeId).Attrs(&partition).FirstOrCreate(&partition)
	if res.Error != nil {
		global.Logger.Errorf("创建%s失败", models_partition)
		return 0, res.Error
	}
	return partition.Id, nil
}
func (s PartitionService) UpdatePartition(partition models.Partition) (uint, error) {
	res := global.DB.Updates(partition)
	if res.Error != nil {
		global.Logger.Errorf("更新%s失败", models_partition)
		return 0, res.Error
	}
	return partition.Id, nil
}

func (s PartitionService) CountPartitions() (count int64, err error) {
	res := global.DB.Find(&models.Partition{}).Count(&count)
	if res.Error != nil {
		global.Logger.Errorf("统计%s失败", models_partition)
		return count, res.Error
	}
	return count, nil
}

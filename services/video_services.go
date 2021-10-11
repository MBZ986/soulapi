package services

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"soulapi/global"
	"soulapi/models"
)

type VideoService struct {
	BaseService
}

func (s VideoService) QueryAllVideos() ([]models.Video, error) {
	var videos []models.Video
	res := global.DB.Find(&videos)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_video)
		return nil, res.Error
	}
	return videos, nil
}

func (s VideoService) QueryVideosByPage(offset, limit int) (*models.Page, error) {
	var videos []models.Video
	res := global.DB.Offset(offset).Limit(limit).Find(&videos)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败：%v", models_video, res.Error)
		return nil, res.Error
	}
	return s.pageHandler(videos, offset, limit, s.CountVideos)
}

func (s VideoService) FindVideos(video models.Video, offset, limit int) (*models.Page, error) {
	var videos []models.Video
	res := global.DB.Where(&video).Offset(offset).Limit(limit).Find(&videos)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_video)
		return nil, res.Error
	}

	return s.pageHandler(videos, offset, limit, s.CountVideos)
}

func (s VideoService) QueryById(id uint) (*models.Video, error) {
	var video models.Video
	res := global.DB.First(&video, id)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_video)
		return nil, res.Error
	}
	spew.Dump(video)
	return &video, nil
}
func (s VideoService) DeleteById(id uint) error {
	var video models.Video
	res := global.DB.Delete(&video, id)
	if res.Error != nil {
		global.Logger.Errorf("删除%s失败", models_video)
		return res.Error
	}
	return nil
}
func (s VideoService) CreateVideo(video models.Video) (uint, error) {
	var tmp models.Video
	if err := global.DB.Where("title = ?", video.Title).First(&tmp).Error; err != nil {
		global.Logger.Errorf("%s名称重复,创建失败", models_video)
		return 0, err
	} else {
		if &tmp != nil {
			return 0, fmt.Errorf("%s标题重复,创建失败", models_video)
		}
	}
	res := global.DB.Where("title = ?", video.Title).Attrs(&video).FirstOrCreate(&video)
	if res.Error != nil {
		global.Logger.Errorf("创建%s失败", models_video)
		return 0, res.Error
	}
	return video.Id, nil
}
func (s VideoService) UpdateVideo(video models.Video) (uint, error) {
	res := global.DB.Updates(video)
	if res.Error != nil {
		global.Logger.Errorf("更新%s失败", models_video)
		return 0, res.Error
	}
	return video.Id, nil
}

func (s VideoService) CountVideos() (count int64, err error) {
	res := global.DB.Find(&models.Video{}).Count(&count)
	if res.Error != nil {
		global.Logger.Errorf("创建%s失败", models_video)
		return count, res.Error
	}
	return count, nil
}

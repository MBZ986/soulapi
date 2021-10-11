package services

import (
	"github.com/davecgh/go-spew/spew"
	"soulapi/global"
	"soulapi/models"
)

type PosterService struct {
	BaseService
}

func (s PosterService) QueryAllPosters() ([]models.Poster, error) {
	var posters []models.Poster
	res := global.DB.Find(&posters)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_poster)
		return nil, res.Error
	}
	return posters, nil
}

func (s PosterService) QueryPostersByPage(offset, limit int) (*models.Page, error) {
	var posters []models.Poster
	res := global.DB.Offset(offset).Limit(limit).Find(&posters)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败：%v", models_poster, res.Error)
		return nil, res.Error
	}
	return s.pageHandler(posters, offset, limit, s.CountPosters)
}

func (s PosterService) FindPosters(poster models.Poster, offset, limit int) (*models.Page, error) {
	var posters []models.Poster
	res := global.DB.Where(&poster).Offset(offset).Limit(limit).Find(&posters)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_poster)
		return nil, res.Error
	}

	return s.pageHandler(posters, offset, limit, s.CountPosters)
}

func (s PosterService) QueryById(id uint) (*models.Poster, error) {
	var poster models.Poster
	res := global.DB.First(&poster, id)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_poster)
		return nil, res.Error
	}
	spew.Dump(poster)
	return &poster, nil
}
func (s PosterService) DeleteById(id uint) error {
	var poster models.Poster
	res := global.DB.Delete(&poster, id)
	if res.Error != nil {
		global.Logger.Errorf("删除%s失败", models_poster)
		return res.Error
	}
	return nil
}
func (s PosterService) CreatePoster(poster models.Poster) (uint, error) {
	res := global.DB.FirstOrCreate(&poster)
	if res.Error != nil {
		global.Logger.Errorf("创建%s失败", models_poster)
		return 0, res.Error
	}
	return poster.Id, nil
}
func (s PosterService) UpdatePoster(poster models.Poster) (uint, error) {
	res := global.DB.Updates(poster)
	if res.Error != nil {
		global.Logger.Errorf("更新%s失败", models_poster)
		return 0, res.Error
	}
	return poster.Id, nil
}

func (s PosterService) CountPosters() (count int64, err error) {
	res := global.DB.Find(&models.Poster{}).Count(&count)
	if res.Error != nil {
		global.Logger.Errorf("创建%s失败", models_poster)
		return count, res.Error
	}
	return count, nil
}

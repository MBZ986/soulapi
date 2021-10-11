package services

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"soulapi/global"
	"soulapi/models"
)

type MusicService struct {
	BaseService
}

func (s MusicService) QueryAllMusics() ([]models.Music, error) {
	var musics []models.Music
	res := global.DB.Find(&musics)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_music)
		return nil, res.Error
	}
	return musics, nil
}

func (s MusicService) QueryMusicsByPage(offset, limit int) (*models.Page, error) {
	var musics []models.Music
	res := global.DB.Offset(offset).Limit(limit).Find(&musics)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败：%v", models_music, res.Error)
		return nil, res.Error
	}
	return s.pageHandler(musics, offset, limit, s.CountMusics)
}

func (s MusicService) FindMusics(music models.Music, offset, limit int) (*models.Page, error) {
	var musics []models.Music
	res := global.DB.Where(&music).Offset(offset).Limit(limit).Find(&musics)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_music)
		return nil, res.Error
	}

	return s.pageHandler(musics, offset, limit, s.CountMusics)
}

func (s MusicService) QueryById(id uint) (*models.Music, error) {
	var music models.Music
	res := global.DB.First(&music, id)
	if res.Error != nil {
		global.Logger.Errorf("查询%s失败", models_music)
		return nil, res.Error
	}
	spew.Dump(music)
	return &music, nil
}
func (s MusicService) DeleteById(id uint) error {
	var music models.Music
	res := global.DB.Delete(&music, id)
	if res.Error != nil {
		global.Logger.Errorf("删除%s失败", models_music)
		return res.Error
	}
	return nil
}
func (s MusicService) CreateMusic(music models.Music) (uint, error) {
	var tmp models.Music
	if err := global.DB.Where("name = ?", music.Name).First(&tmp).Error; err != nil {
		global.Logger.Errorf("%s名称重复,创建失败", models_music)
		return 0, err
	} else {
		if &tmp != nil {
			return 0, fmt.Errorf("%s名称重复,创建失败", models_music)
		}
	}
	res := global.DB.Where("name = ?", music.Name).Attrs(&music).FirstOrCreate(&music)
	if res.Error != nil {
		global.Logger.Errorf("创建%s失败", models_music)
		return 0, res.Error
	}
	return music.Id, nil
}
func (s MusicService) UpdateMusic(music models.Music) (uint, error) {
	res := global.DB.Updates(music)
	if res.Error != nil {
		global.Logger.Errorf("更新%s失败", models_music)
		return 0, res.Error
	}
	return music.Id, nil
}

func (s MusicService) CountMusics() (count int64, err error) {
	res := global.DB.Find(&models.Music{}).Count(&count)
	if res.Error != nil {
		global.Logger.Errorf("创建%s失败", models_music)
		return count, res.Error
	}
	return count, nil
}

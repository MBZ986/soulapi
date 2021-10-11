package services

import (
	"soulapi/global"
	"soulapi/models"
)

type BaseService struct {
}

var (
	models_user           = "用户"
	models_title          = "头衔"
	models_message        = "文章"
	models_music          = "音乐"
	models_video          = "视频"
	models_partition      = "分区"
	models_partition_type = "分区类型"
	models_poster         = "帖子"
	models_label          = "标签"
)

func (s BaseService) pageHandler(data interface{}, offset, limit int, CountFunc func() (count int64, err error)) (*models.Page, error) {
	var (
		page models.Page
		err  error
	)
	//page.Offset = offset
	page.Offset = offset
	page.Limit = limit
	page.Data = data
	if page.Total, err = CountFunc(); err != nil {
		global.Logger.Errorf("查询%s失败：%v", models_title, err)
		return nil, err
	}
	return &page, nil
}

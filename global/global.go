package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"soulapi/conf"
)

var (
	GLOBAL_CONFIG *conf.Server
	DB            *gorm.DB
	Logger        *zap.SugaredLogger
)

const (
	MEDIA_VIDEO   = 100
	MEDIA_MUSIC   = 200
	MEDIA_POSTER  = 300
	MEDIA_MESSAGE = 400
)

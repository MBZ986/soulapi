package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"soulapi/conf"
)

var (
	GLOBAL_CONFIG *conf.Server
	DB            *gorm.DB
	Logger        *zap.SugaredLogger
	Validate      *validator.Validate
	Trans         ut.Translator
)

const (
	MEDIA_VIDEO   = 100
	MEDIA_MUSIC   = 200
	MEDIA_POSTER  = 300
	MEDIA_MESSAGE = 400

	USER_STATE_DEFAULT = 20
	USER_STATE_ADMIN   = 24
	USER_STATE_DISABLE = 28
)

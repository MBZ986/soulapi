package validate

import (
	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"soulapi/global"
)

func init() {
	global.Validate = validator.New()
	uni := ut.New(zh_Hans_CN.New())
	global.Trans, _ = uni.GetTranslator("zh")
	if err := zh.RegisterDefaultTranslations(global.Validate, global.Trans); err != nil {
		panic(err)
	}
}

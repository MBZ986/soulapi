package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/adapter"
	"github.com/go-playground/validator/v10"
	"net/http"
	"soulapi/global"
)

type BaseController struct {
	beego.Controller
}

type ReturnMsg struct {
	Code int
	Msg  string
	Data interface{}
}

func (this *BaseController) Prepare() {
	//fmt.Println("预处理")
	//	github.com/beego/beego/v2/server/web.(*FilterRouter).filter
	//D:/GoPath/pkg/mod/github.com/beego/beego/v2@v2.0.1/server/web/filter.go
	//		81

	if offset, err := this.GetInt("offset", 0); err == nil {
		this.Data["offset"] = offset
		//global.Logger.Debug("解析请求JSON参数失败：%v", err)
	}
	if limit, err := this.GetInt("limit", 10); err == nil {
		this.Data["limit"] = limit
		//global.Logger.Debug("解析请求JSON参数失败：%v", err)
	}

	//if strings.Contains(this.Ctx.Input.Header("content-type"), "application/json") {
	//	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &Page); err != nil {
	//		global.Logger.Debug("解析请求JSON参数失败：%v", err)
	//	}
	//} else {

	//}
}
func (this *BaseController) getPage() (offset int, limit int) {
	return this.Data["offset"].(int), this.Data["limit"].(int)
}
func (this *BaseController) ResData(data interface{}, err error) {
	if err != nil {
		this.Data["json"] = ReturnMsg{
			500, err.Error(), data,
		}
	} else {
		this.Data["json"] = ReturnMsg{
			200, "success", data,
		}
	}
	this.ServeJSON() //对json进行序列化输出
}
func (this *BaseController) Res(err error) {
	if err != nil {
		this.Data["json"] = ReturnMsg{
			500, err.Error(), nil,
		}
	} else {
		this.Data["json"] = ReturnMsg{
			200, "success", nil,
		}
	}
	this.ServeJSON() //对json进行序列化输出
}

func (this *BaseController) Succ(data interface{}) {

	this.Data["json"] = ReturnMsg{
		200, "success", data,
	}
	this.ServeJSON() //对json进行序列化输出
}

func (this *BaseController) Err(code int, msg string, data interface{}) {

	this.Data["json"] = ReturnMsg{
		code, msg, data,
	}
	this.ServeJSON() //对json进行序列化输出
}
func (this *BaseController) ErrMsg(msg string, args ...interface{}) {

	this.Data["json"] = ReturnMsg{
		http.StatusBadRequest, fmt.Sprintf(msg, args...), nil,
	}
	this.ServeJSON() //对json进行序列化输出
}

func (this *BaseController) ErrE(err error) bool {
	if err == nil {
		return true
	}
	if errors, ok := err.(validator.ValidationErrors); ok {
		this.Err(http.StatusBadRequest, "参数校验失败", errors.Translate(global.Trans))
		return false
	}

	this.Data["json"] = ReturnMsg{
		http.StatusBadRequest, err.Error(), nil,
	}
	this.ServeJSON() //对json进行序列化输出
	return false
}

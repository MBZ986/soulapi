package controllers

import (
	"encoding/json"
	"fmt"
	"soulapi/global"
	"soulapi/models"
	"soulapi/services"
)

type UserController struct {
	BaseController
	services.UserService
}

func (c UserController) All() {
	c.ResData(c.QueryAllUsers())
}

func (c UserController) Page() {
	c.ResData(c.QueryUsersByPage(c.getPage()))
}
func (c UserController) Find() {
	var user models.User
	offset, limit := c.getPage()
	if len(c.Ctx.Input.RequestBody) > 0 {
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err != nil {
			c.ErrMsg("参数解析错误:%v", err)
			return
		}
	}

	c.ResData(c.FindUsers(user, offset, limit))
}

func (c UserController) Add() {
	var user models.User
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err != nil {
		c.ErrMsg("参数解析错误:%v", err)
		return
	}
	ok := c.ErrE(global.Validate.Struct(user))
	fmt.Println(ok)
	c.ResData(c.CreateUser(user))
}
func (c UserController) Upd() {
	var user models.User
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err != nil {
		c.ErrMsg("参数解析错误:%v", err)
		return
	}
	fmt.Println("用户", user)
	c.ResData(c.UpdateUser(user))
}
func (c UserController) Del() {
	if id, err := c.GetInt("id"); err != nil {
		c.ErrMsg("参数解析错误:%v", err)
	} else {
		c.Res(c.DeleteById(uint(id)))
	}
}

func (c UserController) Get() {
	if id, err := c.GetInt("id"); err != nil {
		c.ErrMsg("参数解析错误:%v", err)
	} else {
		c.ResData(c.QueryById(uint(id)))
	}
}

func (c UserController) Count() {
	c.ResData(c.CountUsers())
}

func (c UserController) InsertTitle() {
	var (
		userId  int
		titleId int
	)
	userId, _ = c.GetInt("user_id")
	titleId, _ = c.GetInt("title_id")
	if userId == 0 || titleId == 0 {
		c.ErrMsg("用户或者头衔ID不能为空")
		return
	}
	c.Res(c.AddTitle(uint(userId), uint(titleId)))
}

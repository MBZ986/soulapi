package controllers

import (
	"encoding/json"
	"soulapi/models"
	"soulapi/services"
)

//头衔
type TitleController struct {
	BaseController
	services.TitleService
}

func (c TitleController) All() {
	c.ResData(c.QueryAllTitles())
}

func (c TitleController) Page() {
	c.ResData(c.QueryTitlesByPage(c.getPage()))
}
func (c TitleController) Find() {
	var title models.Title
	offset, limit := c.getPage()
	if len(c.Ctx.Input.RequestBody) > 0 {
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &title); err != nil {
			c.ErrMsg("参数解析错误:%v", err)
			return
		}
	}

	c.ResData(c.FindTitles(title, offset, limit))
}

func (c TitleController) Add() {
	var title models.Title
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &title); err != nil {
		c.ErrMsg("参数解析错误:%v", err)
		return
	}
	c.ResData(c.CreateTitle(title))
}
func (c TitleController) Upd() {
	var title models.Title
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &title); err != nil {
		c.ErrMsg("参数解析错误:%v", err)
		return
	}
	c.ResData(c.UpdateTitle(title))
}
func (c TitleController) Del() {
	if id, err := c.GetInt("id"); err != nil {
		c.ErrMsg("参数解析错误:%v", err)
	} else {
		c.Res(c.DeleteById(uint(id)))
	}
}

func (c TitleController) Get() {
	if id, err := c.GetInt("id"); err != nil {
		c.ErrMsg("参数解析错误:%v", err)
	} else {
		c.ResData(c.QueryById(uint(id)))
	}
}

func (c TitleController) Count() {
	c.ResData(c.CountTitles())
}

func (c TitleController) GetUsersById() {
	if id, err := c.GetInt("id"); err != nil {
		c.ErrMsg("参数解析错误:%v", err)
	} else {
		c.ResData(c.QueryUsersByTitleId(uint(id)))
	}
}

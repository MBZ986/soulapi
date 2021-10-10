package controllers

import (
	"soulapi/services"
)

type UserController struct {
	BaseController
	services.UserService
}

func (c UserController) All() {
	c.ResData(c.QueryAllUsers())
}

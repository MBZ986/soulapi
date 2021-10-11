package controllers

import (
	"soulapi/services"
)

type UserController struct {
	BaseController
	services.UserService
}

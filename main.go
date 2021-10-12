package main

import (
	beego "github.com/beego/beego/v2/adapter"
	_ "soulapi/initialize/conf"
	_ "soulapi/initialize/logger"
	_ "soulapi/initialize/mysql"
	_ "soulapi/initialize/validate"
	_ "soulapi/routers"
)

func main() {
	beego.Run()
}

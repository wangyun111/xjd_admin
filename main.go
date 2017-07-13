package main

import (
	"github.com/astaxie/beego"
	_ "xjd_admin/models"
	_ "xjd_admin/routers"
	"xjd_admin/services"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	go services.GetChaheLoginRecord()
	beego.Run()
}

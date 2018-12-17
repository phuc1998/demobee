package main

import (
	_ "demobee/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root@tcp(localhost:3306)/ormdemo?charset=utf8")
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

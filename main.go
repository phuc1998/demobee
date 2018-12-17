package main

import (
	_ "demobee/routers"
	"fmt"
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// log := logs.NewLogger(10000)
	logs.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root@tcp(localhost:3306)/ormdemo?charset=utf8")
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true

	l := logs.GetLogger()
	l.Println("this is a message of http")
	//an official log.Logger with prefix ORM
	logs.GetLogger("ORM").Println("this is a message of orm")

	logs.Debug("my book is bought in the year of ", 2016)
	logs.Info("this %s cat is %v years old", "yellow", 3)
	logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
	logs.Error(1024, "is a very", "good game")
	logs.Critical("oh,crash")

	mapParam := make(map[string]interface{})
	mapParam["appuser"] = "phuc"
	mapParam["appid"] = 33
	mapParam["key1"] = "xBtAOYobmHRIl8UXPN22czME4jgpECUz"
	mapParam["amount"] = 100000
	mapParam["items"] = "[{\"id\":3, \"name\":\"Cafe đen đá\", \"price\":1000, \"quantity\":3}]"
	mapParam["machineid"] = "1542163372950-353"
	req := httplib.Post("http://mep.zpapps.vn/demo-golang/api/createordertk")
	formBody := url.Values{}
	for key, value := range mapParam {
		formBody.Set(key, fmt.Sprintf("%v", value))
	}
	requestBody := formBody.Encode()
	req = req.Body(requestBody)
	response, _ := req.String()
	logs.Info(response)
	beego.Run()

}

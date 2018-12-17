package controllers

import (
	"demobee/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type PhoneController struct {
	beego.Controller
}

// @router /list [get]
func (this *PhoneController) MyGet() {
	o := orm.NewOrm()
	o.Using("default")
	var phones []*models.Phone
	qs := o.QueryTable("phone")
	qs.All(&phones)
	this.Ctx.Output.JSON(phones, false, false)
}

// @router /insert [post]
func (this *PhoneController) MyPost() {
	o := orm.NewOrm()
	o.Using("default")
	phone := new(models.Phone)
	this.Ctx.Request.ParseForm()
	params := this.Ctx.Request.Form
	phone.Id, _ = strconv.Atoi(params["id"][0])
	phone.Name = params["name"][0]
	phone.Screen_size, _ = strconv.Atoi(params["screen_size"][0])
	category := new(models.Category)
	category.Id, _ = strconv.Atoi(params["id_category"][0])
	//o.Read(category)
	phone.Category = category
	_, err := o.Insert(phone)
	if err != nil {
		this.Ctx.WriteString("Insert failed! " + err.Error())
	} else {
		this.Ctx.WriteString("Insert successful!")
	}
}

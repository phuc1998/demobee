package controllers

import (
	"demobee/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type StatisticController struct {
	beego.Controller
}

// @router /list [get]
func (this *StatisticController) MyGet() {
	o := orm.NewOrm()
	o.Using("default")
	var statistic []*models.Statistic
	qs := o.QueryTable("statistic")
	qs.All(&statistic)
	this.Ctx.Output.JSON(statistic, false, false)
}

// @router /insert [post]
func (this *StatisticController) MyPost() {
	phone := new(models.Phone)
	discount := new(models.Discount)
	this.Ctx.Request.ParseForm()
	params := this.Ctx.Request.Form
	phone.Id, _ = strconv.Atoi(params["id_phone"][0])
	discount.Id, _ = strconv.Atoi(params["id_discount"][0])
	statistic := new(models.Statistic)
	statistic.Phone = phone
	statistic.Discount = discount
	statistic.Count = 2
	statistic.Stt, _ = strconv.Atoi(params["stt"][0])
	o := orm.NewOrm()
	//m2m := o.QueryM2M(phone, "discount")
	_, err := o.Insert(statistic)
	if err != nil {
		this.Ctx.WriteString("Insert failed! " + err.Error())
	} else {
		this.Ctx.WriteString("Insert successful!")
	}
}

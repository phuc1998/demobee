package controllers

import (
	"demobee/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type DiscountController struct {
	beego.Controller
}

// @router /list [get]
func (this *DiscountController) MyGet() {

}

// @router /insert [post]
func (this *DiscountController) MyPost() {
	o := orm.NewOrm()
	o.Using("default")
	discount := new(models.Discount)
	this.Ctx.Request.ParseForm()
	params := this.Ctx.Request.Form
	discount.Id, _ = strconv.Atoi(params["id"][0])
	discount.Percent, _ = strconv.Atoi(params["percent"][0])
	_, err := o.Insert(discount)
	if err != nil {
		this.Ctx.WriteString("Insert failed!")
	} else {
		this.Ctx.WriteString("Insert successful!")
	}
}

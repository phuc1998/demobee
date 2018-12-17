package controllers

import (
	"demobee/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CategoryController struct {
	beego.Controller
}

// @router /list [get]
func (this *CategoryController) MyGet() {
	o := orm.NewOrm()
	o.Using("default")
	var category []*models.Category
	qs := o.QueryTable("category")
	qs.All(&category)
	for _, v := range category {
		o.LoadRelated(v, "Phones")
	}
	this.Ctx.Output.JSON(category, false, false)
}

// @router /insert [post]
func (this *CategoryController) MyPost() {
	o := orm.NewOrm()
	o.Using("default")
	this.Ctx.Request.ParseForm()
	params := this.Ctx.Request.Form
	category := new(models.Category)
	category.Id, _ = strconv.Atoi(params["id"][0])
	category.Manufacturer = params["manufacturer"][0]
	_, err := o.Insert(category)
	if err != nil {
		this.Ctx.WriteString("Insert failed!")
	} else {
		this.Ctx.WriteString("Insert successful!")
	}
}

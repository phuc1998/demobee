package routers

import (
	"demobee/controllers"
	"demobee/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {

	orm.RegisterModel(new(models.Phone), new(models.Category), new(models.Discount), new(models.Statistic))

	// ns :=
	// 	beego.NewNamespace("/v1",
	// 		// beego.NSCond(func(ctx *context.Context) bool {
	// 		// 	if ctx.Input.Domain() == "api.beego.me" {
	// 		// 		return true
	// 		// 	}
	// 		// 	return false
	// 		// }),
	// 		// beego.NSGet("/notallowed", func(ctx *context.Context) {
	// 		// 	ctx.Output.Body([]byte("notAllowed"))
	// 		// }),
	// 		beego.NSNamespace("/cms",
	// 			beego.NSInclude(
	// 				&controllers.MainController{},
	// 			),
	// 		),
	// 	)

	// //register namespace
	// beego.AddNamespace(ns)

	//beego.Router("/category", &controllers.CategoryController{})
	//beego.Router("/phone", &controllers.PhoneController{})
	//beego.Router("/discount", &controllers.DiscountController{})
	//beego.Router("/statistic", &controllers.StatisticController{})
	ns := beego.NewNamespace("/demobee",
		beego.NSNamespace("/statistic", beego.NSInclude(&controllers.StatisticController{})),
		beego.NSNamespace("/phone", beego.NSInclude(&controllers.PhoneController{})),
		beego.NSNamespace("/discount", beego.NSInclude(&controllers.DiscountController{})),
		beego.NSNamespace("/category", beego.NSInclude(&controllers.CategoryController{})),
	)
	beego.AddNamespace(ns)
	//beego.Router("", &controllers.MainController{}, "Get:MyMethod")
}

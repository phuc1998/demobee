package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["demobee/controllers:CategoryController"] = append(beego.GlobalControllerRouter["demobee/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "MyPost",
            Router: `/insert`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demobee/controllers:CategoryController"] = append(beego.GlobalControllerRouter["demobee/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "MyGet",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demobee/controllers:DiscountController"] = append(beego.GlobalControllerRouter["demobee/controllers:DiscountController"],
        beego.ControllerComments{
            Method: "MyPost",
            Router: `/insert`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demobee/controllers:DiscountController"] = append(beego.GlobalControllerRouter["demobee/controllers:DiscountController"],
        beego.ControllerComments{
            Method: "MyGet",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demobee/controllers:PhoneController"] = append(beego.GlobalControllerRouter["demobee/controllers:PhoneController"],
        beego.ControllerComments{
            Method: "MyPost",
            Router: `/insert`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demobee/controllers:PhoneController"] = append(beego.GlobalControllerRouter["demobee/controllers:PhoneController"],
        beego.ControllerComments{
            Method: "MyGet",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demobee/controllers:StatisticController"] = append(beego.GlobalControllerRouter["demobee/controllers:StatisticController"],
        beego.ControllerComments{
            Method: "MyPost",
            Router: `/insert`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demobee/controllers:StatisticController"] = append(beego.GlobalControllerRouter["demobee/controllers:StatisticController"],
        beego.ControllerComments{
            Method: "MyGet",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"cmdb/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"fmt"
	"github.com/astaxie/beego/context"
)

var FilterPrintBodyLog = func(ctx *context.Context) {
	var info string
	if ctx.Request.Method == "GET" {
		info = fmt.Sprintf("方法:%s  路径:%s",
			ctx.Request.Method, ctx.Request.RequestURI)
	} else {
		info = fmt.Sprintf("方法:%s  路径:%s\n 请求参数:\n%s",
			ctx.Request.Method, ctx.Request.RequestURI, ctx.Input.RequestBody)
	}
	logs.Info(info)
}

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, FilterPrintBodyLog)

	ns := beego.NewNamespace("/cmdb",
		beego.NSRouter("/ping",
			&controllers.PingController{}, "get:CheckHealth"),

	beego.NSNamespace("/ckeys",
		beego.NSInclude(
			&controllers.CkeysController{},
		),
	),

	beego.NSNamespace("/clink",
		beego.NSInclude(
			&controllers.ClinkController{},
		),
	),

	beego.NSNamespace("/clog",
		beego.NSInclude(
			&controllers.ClogController{},
		),
	),

	beego.NSNamespace("/cobject",
		beego.NSInclude(
			&controllers.CobjectController{},
		),
	),

	beego.NSNamespace("/ctype",
		beego.NSInclude(
			&controllers.CtypeController{},
		),
	),

	beego.NSNamespace("/cview",
		beego.NSInclude(
			&controllers.CviewController{},
		),
	),

	beego.NSNamespace("/cdkey",
		beego.NSInclude(
			&controllers.CdkeyController{},
		),
	),
	)
	beego.AddNamespace(ns)
}

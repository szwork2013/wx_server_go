// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"wx_server_go/controllers/web/v1/cus"
	"wx_server_go/controllers/web/v1/wx"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/web/v1",
		beego.NSNamespace("/subscribe",
			beego.NSInclude(
				&wx.WxSubscribeController{},
			),
		),
		beego.NSNamespace("/wxtask",
			beego.NSInclude(
				&wx.WxTaskController{},
			),
		),
		beego.NSNamespace("/cusmbr",
			beego.NSInclude(
				&cus.CusMbrController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

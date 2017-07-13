// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool utils to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"xjd_admin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/sysuser", beego.NSInclude(&controllers.SysUserController{})),
		beego.NSNamespace("/user", beego.NSInclude(&controllers.UsersController{})),
	)
	beego.AddNamespace(ns)
}

package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["xjd_admin/controllers:SysUserController"] = append(beego.GlobalControllerRouter["xjd_admin/controllers:SysUserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["xjd_admin/controllers:SysUserController"] = append(beego.GlobalControllerRouter["xjd_admin/controllers:SysUserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["xjd_admin/controllers:SysUserController"] = append(beego.GlobalControllerRouter["xjd_admin/controllers:SysUserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["xjd_admin/controllers:SysUserController"] = append(beego.GlobalControllerRouter["xjd_admin/controllers:SysUserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["xjd_admin/controllers:SysUserController"] = append(beego.GlobalControllerRouter["xjd_admin/controllers:SysUserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/all`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["xjd_admin/controllers:SysUserController"] = append(beego.GlobalControllerRouter["xjd_admin/controllers:SysUserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["xjd_admin/controllers:UsersController"] = append(beego.GlobalControllerRouter["xjd_admin/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["xjd_admin/controllers:UsersController"] = append(beego.GlobalControllerRouter["xjd_admin/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["xjd_admin/controllers:UsersController"] = append(beego.GlobalControllerRouter["xjd_admin/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["xjd_admin/controllers:UsersController"] = append(beego.GlobalControllerRouter["xjd_admin/controllers:UsersController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/all`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["xjd_admin/controllers:UsersController"] = append(beego.GlobalControllerRouter["xjd_admin/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}

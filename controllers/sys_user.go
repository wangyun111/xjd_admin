package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"xjd_admin/models"
	// "xjd_admin/services"
	"xjd_admin/utils"
)

//管理员模块
type SysUserController struct {
	beego.Controller
}

//@Summary 管理员登录
//@Description
//@Param  account      query  string  true   "账号"
//@Param  password     query  string  true   "密码"
//@Success    200  {object} models.SysUser
//@Failure    304  code 304  登录失败
//@Failure    404  code 404  参数为空
//@router /login [get]
func (this *SysUserController) Login() {
	defer this.ServeJSON()
	account := this.GetString("account")
	password := this.GetString("password")
	if account == "" || password == "" {
		this.Data["json"] = utils.ParamNull()
	} else {
		_SysUser, _err := models.Login(account, password)
		// services.AddCacheLoginRecord(account, password)
		if _err == nil {
			this.Data["json"] = _SysUser
		} else {
			this.Data["json"] = utils.OperationFalse("登录失败")
		}
	}
}

//@Summary 添加管理员信息
//@Description
//@Param    body      body     models.SysUser    true   "账号信息"
//@Success   200    成功
//@Failure   304    添加失败
//@router / [post]
func (this *SysUserController) Post() {
	defer this.ServeJSON()
	var _SysUser models.SysUser
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &_SysUser)
	if err != nil {
		this.Data["json"] = utils.ParamError("参数有误:" + err.Error())
	} else {
		account := _SysUser.Name
		password := _SysUser.Password
		if account == "" || password == "" {
			this.Data["json"] = utils.ParamNull()
		} else {
			_err := models.AddSysUser(_SysUser)
			if _err == nil {
				this.Data["json"] = utils.OperationSuccess()
			} else {
				this.Data["json"] = utils.OperationFalse()
			}
		}
	}
}

//@Summary 删除管理员信息
//@Description
//@Param  _id      query  string  true   "账号"
//@Success   200    操作成功
//@Failure   304    登录失败
//@router / [delete]
func (this *SysUserController) Delete() {
	defer this.ServeJSON()
	_id := this.GetString("_id")
	if _id == "" {
		this.Data["json"] = utils.ParamNull()
	} else {
		if !models.CheckIsBsonId(_id) {
			this.Data["json"] = utils.ParamError("参数有误。")
		} else {
			_err := models.DeleteSysUser(_id)
			beego.Info(_err, _id)
			if _err == nil {
				this.Data["json"] = utils.OperationSuccess()
			} else {
				this.Data["json"] = utils.OperationFalse()
			}
		}
	}
}

//@Summary 修改管理员信息
//@Description
//@Param    body      body     models.SysUser    true   "账号信息"
//@Success   200    操作成功
//@Failure   304    登录失败
//@router / [put]
func (this *SysUserController) Put() {
	defer this.ServeJSON()
	var _SysUser models.SysUser
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &_SysUser)
	if err != nil {
		this.Data["json"] = utils.ParamError("参数有误:" + err.Error())
	} else {
		_id := _SysUser.Id
		if _id == "" {
			this.Data["json"] = utils.ParamNull()
		} else {
			_err := models.UpdateSysUser(_SysUser)
			if _err == nil {
				this.Data["json"] = utils.OperationSuccess()
			} else {
				this.Data["json"] = utils.OperationFalse()
			}
		}
	}
}

//@Summary 分页查询管理员信息
//@Description
//@Param  account      query  string  false   "账号"
//@Param  beginTime    query  string  false   "开始时间"
//@Param  endTime      query  string  false   "结束时间"
//@Success   200     {object} models.SysUser
//@Failure   304    登录失败
//@router /all [get]
func (this *SysUserController) GetAll() {
	defer this.ServeJSON()
	account := this.GetString("account")
	beginTime := this.GetString("beginTime")
	endTime := this.GetString("endTime")
	list, _err := models.QuerySysUser(account, beginTime, endTime)
	if _err == nil && len(list) > 0 {
		this.Data["json"] = list
	} else {
		this.Data["json"] = utils.ReturnDataNull()
	}
}

//@Summary 查询单个管理员信息
//@Description
//@Param     _id      query  string  true   "id值"
//@Success   200     {object} models.SysUser
//@Failure   404     参数为空
//@Failure   401     参数有误
//@Failure   400     返回结果为空
//@router / [get]
func (this *SysUserController) Get() {
	defer this.ServeJSON()
	_id := this.GetString("_id")
	if _id == "" {
		this.Data["json"] = utils.ParamNull()
	} else {
		if !models.CheckIsBsonId(_id) {
			this.Data["json"] = utils.ParamError("参数有误。")
		} else {
			_SysUser, _err := models.FindByIdSysUser(_id)
			if _err == nil {
				this.Data["json"] = _SysUser
			} else {
				this.Data["json"] = utils.ReturnDataNull()
			}
		}
	}
}

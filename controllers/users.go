package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"sexy_tools/time"
	"strings"
	"xjd_admin/models"
	"xjd_admin/services"
	"xjd_admin/utils"
)

//用户模块
type UsersController struct {
	beego.Controller
}

//@Summary 用户登录
//@Description
//@Param  account           query  string  true   "账号"
//@Param  loginPassword     query  string  true   "密码"
//@Success    200  {object} models.Users
//@Failure    304  code 304  登录失败
//@Failure    404  code 404  参数为空
//@router /login [get]
func (this *UsersController) Login() {
	defer this.ServeJSON()
	account := this.GetString("account")
	loginPassword := this.GetString("loginPassword")
	if account == "" || loginPassword == "" {
		this.Data["json"] = utils.ParamNull()
	} else {
		_User, _err := services.GetChaheUserInfo(utils.UserInfo+account, account)
		if _err != nil {
			this.Data["json"] = utils.OperationFalse("账号不存在,请仔细核对！")
			return
		} else {
			if !strings.EqualFold(_User.LoginPassword, loginPassword) {
				count := services.CheckPwdErrSeveral(utils.PsdErrSeveral + account)
				if count == 0 {
					this.Data["json"] = utils.OperationFalse("密码错误达到5次,请明天再试！")
					return
				}
			} else {
				models.Rc.Delete(utils.PsdErrSeveral + loginPassword)
			}
		}
		_User, _err = models.UserLogin(account, loginPassword)
		if _err == nil {
			ip := this.Ctx.Input.IP()
			var _Logs models.Logs
			_Logs.OperateTime = time.GetNowTime()
			_Logs.Account = account
			_Logs.Uid = _User.Id.Hex()
			_Logs.AppVersion = _User.AppVersion
			_Logs.Content = "用户登录"
			_Logs.Model = "login"
			_Logs.OperateIp = ip
			models.Rc.LPush(utils.UserOperationLogs, &_Logs) //记录用户操作日志
			//添加登录记录
			//修改登录时间,生成新的token等
			if _User.State == 1 {
				this.Data["json"] = utils.OperationFalse("该账户被冻结。")
			} else {
				this.Data["json"] = map[string]interface{}{"code": 200}
			}
		} else {
			this.Data["json"] = utils.OperationFalse("账号密码错误")
		}
	}
}

//@Summary 添加用户信息(用户注册)
//@Description
//@Param    body      body     models.UsersParams    true   "用户信息"
//@Success   200    成功
//@Failure   304    添加失败
//@router / [post]
func (this *UsersController) Post() {
	defer this.ServeJSON()
	var _Users models.UsersParams
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &_Users)
	if err != nil {
		this.Data["json"] = utils.ParamError("参数有误:" + err.Error())
	} else {
		account := _Users.Account
		loginPassword := _Users.LoginPassword
		if account == "" || loginPassword == "" {
			this.Data["json"] = utils.ParamNull()
		} else {
			_err := models.AddUsers(_Users.Users)
			if _err == nil {
				this.Data["json"] = utils.OperationSuccess()
			} else {
				this.Data["json"] = utils.OperationFalse()
			}
		}
	}
}

//@Summary 修改会员信息
//@Description
//@Param    body      body     models.Users    true   "会员信息"
//@Success   200    操作成功
//@Failure   304    登录失败
//@router / [put]
func (this *UsersController) Put() {
	defer this.ServeJSON()
	var _Users models.Users
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &_Users)
	if err != nil {
		this.Data["json"] = utils.ParamError("参数有误:" + err.Error())
	} else {
		_id := _Users.Id
		if _id == "" {
			this.Data["json"] = utils.ParamNull()
		} else {
			_err := models.UpdateUsers(_Users)
			if _err == nil {
				this.Data["json"] = utils.OperationSuccess()
			} else {
				this.Data["json"] = utils.OperationFalse()
			}
		}
	}
}

//@Summary 分页查询会员信息
//@Description
//@Param  account      query  string  false   "账号"
//@Param  beginTime    query  string  false   "开始时间"
//@Param  endTime      query  string  false   "结束时间"
//@Success   200     {object} models.Page
//@Failure   304    登录失败
//@router /all [get]
func (this *UsersController) GetAll() {
	defer this.ServeJSON()
	pageNo, _ := this.GetInt("pageNo")
	pageSize, _ := this.GetInt("pageSize")
	account := this.GetString("account")
	beginTime := this.GetString("beginTime")
	endTime := this.GetString("endTime")
	smap := map[string]interface{}{}
	timeMap := map[string]string{}
	if account != "" {
		smap["name"] = account
	}
	if beginTime != "" {
		timeMap["$gte"] = beginTime
	}
	if endTime != "" {
		timeMap["$lte"] = endTime
	}
	if len(timeMap) > 0 {
		smap["createTime"] = timeMap
	}
	count, _ := models.QueryUsersCount(smap)
	if count < 1 {
		this.Data["json"] = utils.ReturnDataNull()
	} else {
		if pageNo < 1 {
			pageNo = 1
		}
		if pageSize < 1 {
			pageSize = 20
		}
		list, _ := models.QueryUsers(pageNo, pageSize, smap)
		page := models.PageUtil(count, pageNo, pageSize, list)
		this.Data["json"] = page
	}
}

//@Summary 查询单个管理员信息
//@Description
//@Param     account      query  string  true   "id值"
//@Success   200     {object} models.Users
//@Failure   404     参数为空
//@Failure   401     参数有误
//@Failure   400     返回结果为空
//@router / [get]
func (this *UsersController) Get() {
	defer this.ServeJSON()
	account := this.GetString("account")
	if account == "" {
		this.Data["json"] = utils.ParamNull()
	} else {
		_Users, _err := models.FindByIdUsers(account)
		if _err == nil {
			this.Data["json"] = _Users
		} else {
			this.Data["json"] = utils.ReturnDataNull()
		}
	}
}

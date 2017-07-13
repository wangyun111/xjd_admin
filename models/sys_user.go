package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"sexy_tools/time"
	// "xjd_admin/models/config"
)

type SysUser struct {
	Id            bson.ObjectId `bson:"_id"           json:"_id"           description:"主键编号"`
	Name          string        `bson:"name"          json:"name"          description:"账号"`
	Password      string        `bson:"password"      json:"password"      description:"密码"`
	Displayname   string        `bson:"displayName"   json:"displayName"   description:"使用人姓名"`
	Sex           string        `bson:"sex"           json:"sex"           description:"性别"`
	Email         string        `bson:"email"         json:"email"         description:"邮箱"`
	Phone         string        `bson:"phone"         json:"phone"         description:"联系电话"`
	Accountstatus string        `bson:"accountStatus" json:"accountStatus" description:"状态(启用,禁用)"`
	Remark        string        `bson:"remark"        json:"remark"        description:"备注"`
	CreateTime    string        `bson:"createTime"    json:"createTime"    description:"创建时间"`
	Code          string        `bson:"code"          json:"code"          description:"code标识"`
	RoleId        int           `bson:"roleId"        json:"roleId"        description:"角色id"`
	StationId     int           `bson:"stationId"     json:"stationId"     description:"岗位id"`
	QnAccount     string        `bson:"qnAccount"     json:"qnAccount"     description:"青牛账号"`
	QnPassword    string        `bson:"qnPassword"    json:"qnPassword"    description:"青牛密码"`
	PlaceRole     string        `bson:"placeRole"     json:"placeRole"     description:"坐席角色"`
	TapeId        int           `bson:"tapeId"        json:"tapeId"        description:"最新坐席通话id"`
	AccountType   int           `bson:"accountType"   json:"accountType"   description:"账号类型(0内部,1外部)"`
}

//管理员登录
func Login(account, password string) (_SysUser SysUser, _err error) {
	session := GetSession()
	c := session.DB("wangyun").C("sys_user")
	smap := map[string]interface{}{}
	smap["name"] = account
	smap["password"] = password
	_err = c.Find(&smap).One(&_SysUser)
	return
}

//添加管理员
func AddSysUser(_SysUser SysUser) (_err error) {
	session := GetSession()
	c := session.DB("wangyun").C("sys_user")
	_SysUser.Id = bson.NewObjectId()
	_SysUser.CreateTime = time.GetNowTime()
	_err = c.Insert(&_SysUser)
	return
}

func DeleteSysUser(_id string) (_err error) {
	session := GetSession()
	c := session.DB("wangyun").C("sys_user")
	_err = c.RemoveId(bson.ObjectIdHex(_id))
	beego.Info(_err)
	return
}

func UpdateSysUser(_SysUser SysUser) (_err error) {
	_id := _SysUser.Id
	password := _SysUser.Password
	displayName := _SysUser.Displayname
	email := _SysUser.Email
	phone := _SysUser.Phone
	accountStatus := _SysUser.Accountstatus
	roleId := _SysUser.RoleId
	stationId := _SysUser.StationId
	placeRole := _SysUser.PlaceRole
	accountType := _SysUser.AccountType
	smap := map[string]interface{}{}
	if password != "" {
		smap["password"] = password
	}
	if displayName != "" {
		smap["displayName"] = displayName
	}
	if email != "" {
		smap["email"] = email
	}
	if phone != "" {
		smap["phone"] = phone
	}
	if accountStatus != "" {
		smap["accountStatus"] = accountStatus
	}
	if roleId > 0 {
		smap["roleId"] = roleId
	}
	if stationId > 0 {
		smap["stationId"] = stationId
	}
	if placeRole != "" {
		smap["placeRole"] = placeRole
	}
	if accountType > 0 {
		smap["accountType"] = accountType
	}
	session := GetSession()
	c := session.DB("wangyun").C("sys_user")
	_err = c.UpdateId(_id, bson.M{"$set": smap})
	return
}

func QuerySysUser(account, beginTime, endTime string) (list []SysUser, _err error) {
	session := GetSession()
	c := session.DB("wangyun").C("sys_user")
	smap := map[string]interface{}{}
	timeMap := map[string]string{}
	tMap := map[string]interface{}{}
	// []string
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

	tMap["$in"] = []string{"汪赟", "lishan"}
	smap["displayName"] = tMap
	beego.Info(smap)
	_err = c.Find(&smap).Sort("-createTime").All(&list)
	beego.Info(_err)
	return
}

func FindByIdSysUser(_id string) (_SysUser SysUser, _err error) {
	session := GetSession()
	c := session.DB("wangyun").C("sys_user")
	_err = c.FindId(bson.ObjectIdHex(_id)).One(&_SysUser)
	return
}

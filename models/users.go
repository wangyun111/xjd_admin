package models

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"sexy_tools/time"
	"xjd_admin/utils"
)

type Users struct {
	Id                  bson.ObjectId `bson:"_id"                   json:"_id"                   description:"主键编号"`
	Code                string        `bson:"code"                  json:"code"                  description:"账号表示"`
	WrCode              string        `bson:"wrCode"                json:"wrCode"                description:"微融code"`
	Account             string        `bson:"account"               json:"account"               description:"账号"`
	LoginPassword       string        `bson:"loginPassword"         json:"loginPassword"         description:"登录密码"`
	State               int           `bson:"state"                 json:"State"                 description:"状态"`
	CreateTime          string        `bson:"createTime"            json:"createTime"            description:"创建时间"`
	MobileType          string        `bson:"mobileType"            json:"mobileType"            description:"手机类型"`
	MobileVersion       string        `bson:"mobileVersion"         json:"mobileVersion"         description:"手机状态"`
	LoginTime           string        `bson:"loginTime"             json:"loginTime"             description:"登录时间"`
	HisDevicecode       string        `bson:"hisDevicecode"         json:"hisDevicecode"         description:"手机设备标识"`
	MobileTypeRecent    string        `bson:"mobileTypeRecent"      json:"mobileTypeRecent"      description:"最近操作的结果，手机型号"`
	OperationTime       string        `bson:"operationTime"         json:"operationTime"         description:"最近操作的时间（1.2.1 版本加入的获取用户接口）"`
	App                 int           `bson:"app"                   json:"app"                   description:"1 ios,2 android,3 wx,4pc5 wd 6 wap"`
	Token               string        `bson:"token"                 json:"token"                 description:"服务器保存一致信息"`
	PkgType             string        `bson:"pkgType"               json:"pkgType"               description:"包的标识"`
	Source              string        `bson:"source"                json:"source"                description:"来源"`
	AppVersion          string        `bson:"appVersion"            json:"appVersion"            description:"app版本"`
	AppVersionRecent    string        `bson:"appVersionRecent"      json:"appVersionRecent"      description:"最近操作app版本"`
	UserImg             string        `bson:"userImg"               json:"userImg"               description:"用户头像"`
	Isemulator          string        `bson:"isemulator"            json:"isemulator"            description:"是否是模拟器"`
	Jpushid             string        `bson:"jpushid"               json:"jpushid"               description:"手机推送标识"`
	LoanControl         int           `bson:"loanControl"           json:"loanControl"           description:"借款控制: 0进行中有一笔不能申请第二笔,1:不做限制"`
	Ip                  int           `bson:"ip"                    json:"ip"                    description:"ip地址"`
	HisDevicecodeRecent int           `bson:"hisDevicecodeRecent"   json:"hisDevicecodeRecent"   description:"最近手机设备标识"`
	Location            int           `bson:"location"              json:"location"              description:"地址到省市"`
	Address             int           `bson:"address"               json:"address"               description:"详细地址到具体街道"`
	IpLocation          int           `bson:"ipLocation"            json:"ipLocation"            description:"ip地址到省市"`
	IpAddress           int           `bson:"ipAddress"             json:"ipAddress"             description:"ip详细地址"`
	TagType             int           `bson:"tagType"               json:"tagType"               description:"标记欺诈类型"`
	TagSysId            int           `bson:"tagSysId"              json:"tagSysId"              description:"标记人账号id"`
}

type UsersParams struct {
	Users
	VerifyCode string
}

//用户登录
func UserLogin(account, loginPassword string) (_Users Users, _err error) {
	session := GetSession()
	c := session.DB("wangyun").C("users")
	smap := map[string]interface{}{}
	smap["account"] = account
	smap["loginPassword"] = loginPassword
	_err = c.Find(&smap).One(&_Users)
	return
}

//添加管理员
func AddUsers(_Users Users) (_err error) {
	session := GetSession()
	c := session.DB("wangyun").C("users")
	_Users.Id = bson.NewObjectId()
	_Users.CreateTime = time.GetNowTime()
	_err = c.Insert(&_Users)
	return
}

func DeleteUsers(_id string) (_err error) {
	session := GetSession()
	c := session.DB("wangyun").C("users")
	_err = c.RemoveId(bson.ObjectIdHex(_id))
	beego.Info(_err)
	return
}

func UpdateUsers(_Users Users) (_err error) {
	// _id := _Users.Id
	account := _Users.Account
	loginTime := _Users.LoginTime
	loginPassword := _Users.LoginPassword
	smap := map[string]interface{}{}
	if loginPassword != "" {
		smap["loginPassword"] = loginPassword
	}
	if loginTime != "" {
		smap["loginTime"] = loginTime
	}
	session := GetSession()
	c := session.DB("wangyun").C("users")
	c.Update(bson.M{"account": account}, bson.M{"$set": smap})
	// _err = c.UpdateId(_id, bson.M{"$set": smap})
	return
}

func QueryUsersCount(smap map[string]interface{}) (count int, err error) {
	session := GetSession()
	c := session.DB("wangyun").C("users")
	count, err = c.Find(&smap).Count()
	return
}

func QueryUsers(pageNo, pageSize int, smap map[string]interface{}) (list []Users, _err error) {
	session := GetSession()
	c := session.DB("wangyun").C("users")
	_err = c.Find(&smap).Skip((pageNo - 1) * pageSize).Limit(pageSize).
		Sort("-createTime").All(&list)
	beego.Info(_err)
	return
}

func FindByIdUsers(account string) (_Users Users, _err error) {
	smap := map[string]interface{}{}
	if account != "" {
		redisKey := utils.UserAccount
		if Rc.IsExist(redisKey + account) {
			if byteS, _err1 := Rc.RedisBytes(redisKey); _err1 == nil {
				if _err = json.Unmarshal(byteS, &_Users); _err == nil {
					return
				}
			}
		}
		smap["account"] = account
	}
	session := GetSession()
	c := session.DB("wangyun").C("users")
	_err = c.Find(&smap).One(&_Users)
	beego.Info(_err)
	if _err == nil {
		byteS, _err2 := json.Marshal(&_Users)
		if _err2 == nil {
			redisKeyUserAccount := utils.UserAccount + _Users.Account
			Rc.Put(redisKeyUserAccount, byteS, utils.OneHour)
		}
	}
	return
}

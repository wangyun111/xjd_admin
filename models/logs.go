package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"sexy_tools/time"
	// "xjd_admin/utils"
)

//操作日志
type Logs struct {
	Id            bson.ObjectId `bson:"_id"                  json:"_id"             description:"主键编号"`
	Uid           string        `bson:"uid"                  json:"uid"             description:"用户id"`
	Account       string        `bson:"account"              json:"account"         description:"操作账号"`
	OperateTime   string        `bson:"operateTime"          json:"operateTime"     description:"操作时间"`
	Content       string        `bson:"content"              json:"content"         description:"日志内容"`
	OperateIp     string        `bson:"operateIp"            json:"operateIp"       description:"登录ip"`
	Model         string        `bson:"model"                json:"model"           description:"操作模式"`
	IsEmailSend   int           `bson:"isEmailSend"          json:"isEmailSend"     description:"是否模拟器"`
	AppVersion    string        `bson:"appVersion"           json:"appVersion"      description:"app版本号"`
	MobileType    string        `bson:"mobileType"           json:"mobileType"      description:"手机类型"`
	MobileVersion string        `bson:"mobileVersion"        json:"mobileVersion"   description:"手机型号"`
}

//添加管理员
func AddLogs(_Logs Logs) (_err error) {
	session := GetSession()
	c := session.DB("wangyun").C("logs")
	_Logs.Id = bson.NewObjectId()
	_Logs.OperateTime = time.GetNowTime()
	_err = c.Insert(&_Logs)
	return
}

func QueryLogs(account, beginTime, endTime string) (list []Logs, _err error) {
	session := GetSession()
	c := session.DB("wangyun").C("logs")
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
	beego.Info(len(timeMap))
	_err = c.Find(&smap).Sort("-createTime").All(&list)
	beego.Info(_err)
	return
}

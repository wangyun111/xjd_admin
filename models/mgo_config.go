package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	mgo_url = "127.0.0.1:27017"
	// "mongodb://ludan:123456@101.37.86.154:27017,otherhost:27017/weixin"
)

var (
	MgoSession *mgo.Session
	DataBase   = "wangyun"
)

/**
 * 公共方法，获取session，如果存在则拷贝一份
 */
func GetSession() *mgo.Session {
	if MgoSession == nil {
		var err error
		MgoSession, err = mgo.Dial(mgo_url)
		if err != nil {
			panic(err) //直接终止程序运行
		}
	}
	//最大连接池默认为4096
	return MgoSession.Clone()
}

//验证id是否符合mongodbId类型
func CheckIsBsonId(_id string) bool {
	return bson.IsObjectIdHex(_id)
}

//公共方法，获取collection对象
func WitchCollection(collection string, s func(*mgo.Collection) error) error {
	session := GetSession()
	defer session.Close()
	c := session.DB(DataBase).C(collection)
	return s(c)
}

func init() {
	MgoSession = GetSession()
}

package services

import (
	"encoding/json"
	"fmt"
	"sexy_tools/time"
	"xjd_admin/models"
)

//添加用户登录记录
func AddCacheLoginRecord(_Logs models.Logs) {
	if models.Rc != nil {
		models.Rc.LPush("loginLogs", &_Logs)
	}
}

//获取用户登录记录
func GetChaheLoginRecord() {
	// var _SysUser models.SysUser
	for {
		models.Rc.Brpop("loginRecord", func(b []byte) {
			fmt.Println(string(b))
		})
	}

}

func GetChaheUserInfo(key, account string) (_Users models.Users, _err error) {
	rc := models.Rc
	if rc.IsExist(key) {
		if byteS, _err1 := rc.RedisBytes(key); _err1 == nil {
			_err = json.Unmarshal(byteS, &_Users)
			if _err == nil {
				return
			}
		}
	}
	_Users, _err = models.FindByIdUsers(account)
	return
}

func CheckPwdErrSeveral(key string) (count int) {
	count = 5
	rc := models.Rc
	if rc.IsExist(key) {
		count, _ = rc.RedisInt(key)
	}
	count--
	rc.Put(key, count, time.GetTodayLastSecond())
	return
}

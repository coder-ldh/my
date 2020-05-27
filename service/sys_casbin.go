package service

import (
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/util"
	gormadapter "github.com/casbin/gorm-adapter"
	"my/global"
	"my/model/request"
	"strings"
)

// @title    GetPolicyPathByAuthorityId
// @description   get policy path by authorityId, 获取权限列表
// @auth                     （2020/04/05  20:22）
// @param     authorityId     string
// @return                    []string

func GetPolicyPathByAuthorityId(authorityId string) (pathMaps []request.CasbinInfo) {
	e := Casbin()
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

// @title    ClearCasbin
// @description   清除匹配的权限
// @auth                     （2020/04/05  20:22）
// @param     v               int
// @param     p               string
// @return                    bool

func ClearCasbin(v int, p ...string) bool {
	e := Casbin()
	return e.RemoveFilteredPolicy(v, p...)

}

// @title    Casbin
// @description   store to DB, 持久化到数据库  引入自定义规则
// @auth                     （2020/04/05  20:22）

func Casbin() *casbin.Enforcer {
	a := gormadapter.NewAdapterByDB(global.GVA_DB)
	e := casbin.NewEnforcer(global.GVA_CONFIG.Casbin.ModelPath, a)
	e.AddFunction("ParamsMatch", ParamsMatchFunc)
	_ = e.LoadPolicy()
	return e
}

// @title    ParamsMatch
// @description   customized rule, 自定义规则函数
// @auth                     （2020/04/05  20:22）
// @param     fullNameKey1    string
// @param     key2            string
// @return                    bool

func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	//剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

// @title    ParamsMatchFunc
// @description   customized function, 自定义规则函数
// @auth                     （2020/04/05  20:22）
// @param     args            ...interface{}
// @return                    interface{}
// @return                    error

func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return (bool)(ParamsMatch(name1, name2)), nil
}

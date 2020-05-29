package model

type AdminResource struct {
	Base
	//用户Id
	AdminId int "json:adminId"
	//资源ID
	ResourceId int "json:resourceId"
}

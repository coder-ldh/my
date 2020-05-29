package model

type RoleResource struct {
	Base
	//角色ID
	RoleId int "json:roleId"
	//资源ID
	ResourceId int "json:resourceId"
}

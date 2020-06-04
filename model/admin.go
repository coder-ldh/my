package model

type Admin struct {
	Base
	//登陆名
	LoginName *string "json:loginName"
	//密码
	Password *string "json:-"
	//手机号
	Phone *string "json:phone"
	//角色ID
	RoleId *int "json:roleId"
	//头像
	Photo *string "json:photo"
}

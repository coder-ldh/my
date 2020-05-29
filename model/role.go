package model

type Role struct {
	Base
	//角色名称
	Name string "json:name"
	//角色说明
	Content string "json:content"
}

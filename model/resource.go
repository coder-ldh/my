package model

type Resource struct {
	Base
	//权限名称 权限名称
	Name string "json:name"
	//访问路径 访问路径
	Path string "json:path"
}

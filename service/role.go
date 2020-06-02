package service

import (
	"github.com/jinzhu/gorm"
	"my/global"
	"my/model"
)

func GetRoleById(roleId int) (*model.Role, error) {
	var role *model.Role
	error := global.GVA_DB.Find(&role, roleId).Error
	if error != nil && error != gorm.ErrRecordNotFound {
		return nil, error
	}
	return role, nil
}

func GetRoleList(pageNum int, pageSize int) ([]*model.Role, error) {
	var roles []*model.Role
	error := global.GVA_DB.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&roles).Error
	if error != nil && error != gorm.ErrRecordNotFound {
		return nil, error
	}
	return roles, nil
}

func GetRoleByName(name string) (*model.Role, error) {
	var role *model.Role
	error := global.GVA_DB.Find(&role).Error
	if error != nil && error != gorm.ErrRecordNotFound {
		return nil, error
	}
	return role, nil
}

func UpdateRole(model.Role) (*model.Role, error) {
	var role *model.Role
	error := global.GVA_DB.Update(&role).Error
	if error != nil {
		return nil, error
	}
	return role, nil
}

func AddRole(role model.Role) (*model.Role, error) {
	error := global.GVA_DB.Save(&role).Error
	if error != nil {
		return nil, error
	}
	return &role, nil
}

func DeleteRole(roleId int) error {
	var role = new(model.Role)
	role.Delete = true
	role.Id = roleId
	error := global.GVA_DB.Update(&role).Error
	if error != nil {
		return error
	}
	return nil
}

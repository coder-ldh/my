package service

import (
	"github.com/jinzhu/gorm"
	"my/global"
	"my/model"
)

func GetRoleResourceList(roleResource model.RoleResource) ([]*model.RoleResource, error) {
	var roleResources []*model.RoleResource
	error := global.GVA_DB.Find(&roleResources).Error
	if error != nil && error != gorm.ErrRecordNotFound {
		return nil, error
	}
	return roleResources, nil
}

func UpdateRoleResource(roleResource *model.RoleResource) (*model.RoleResource, error) {
	error := global.GVA_DB.Update(roleResource).Error
	if error != nil {
		return nil, error
	}
	return roleResource, nil
}

func AddRoleResource(roleResource model.RoleResource) (*model.RoleResource, error) {
	error := global.GVA_DB.Save(&roleResource).Error
	if error != nil {
		return nil, error
	}
	return &roleResource, nil
}

func DeleteRoleResource(roleResourceId int) error {
	var roleResource = new(model.RoleResource)
	roleResource.Delete = true
	roleResource.Id = roleResourceId
	error := global.GVA_DB.Update(&roleResource).Error
	if error != nil {
		return error
	}
	return nil
}

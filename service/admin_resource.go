package service

import (
	"github.com/jinzhu/gorm"
	"my/global"
	"my/model"
)

func GetAdminResourceList(adminResource model.AdminResource) ([]*model.AdminResource, error) {
	var adminResources []*model.AdminResource
	error := global.GVA_DB.Find(&adminResources).Error
	if error != nil && error != gorm.ErrRecordNotFound {
		return nil, error
	}
	return adminResources, nil
}

func UpdateAdminResource(model.AdminResource) (*model.AdminResource, error) {
	var adminResource *model.AdminResource
	error := global.GVA_DB.Update(&adminResource).Error
	if error != nil {
		return nil, error
	}
	return adminResource, nil
}

func AddAdminResource(adminResource model.AdminResource) (*model.AdminResource, error) {
	error := global.GVA_DB.Save(&adminResource).Error
	if error != nil {
		return nil, error
	}
	return &adminResource, nil
}

func DeleteAdminResource(adminResourceId int) error {
	var adminResource = new(model.AdminResource)
	adminResource.Delete = true
	adminResource.Id = adminResourceId
	error := global.GVA_DB.Update(&adminResource).Error
	if error != nil {
		return error
	}
	return nil
}

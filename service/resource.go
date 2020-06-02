package service

import (
	"github.com/jinzhu/gorm"
	"my/global"
	"my/model"
)

func GetResourceById(resourceId int) (*model.Resource, error) {
	var resource *model.Resource
	error := global.GVA_DB.Find(&resource, resourceId).Error
	if error != nil && error != gorm.ErrRecordNotFound {
		return nil, error
	}
	return resource, nil
}

func GetResourceList(pageNum int, pageSize int) ([]*model.Resource, error) {
	var resources []*model.Resource
	error := global.GVA_DB.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&resources).Error
	if error != nil && error != gorm.ErrRecordNotFound {
		return nil, error
	}
	return resources, nil
}

func GetResourceByName(name string) (*model.Resource, error) {
	var resource *model.Resource
	error := global.GVA_DB.Find(&resource).Error
	if error != nil && error != gorm.ErrRecordNotFound {
		return nil, error
	}
	return resource, nil
}

func UpdateResource(model.Resource) (*model.Resource, error) {
	var resource *model.Resource
	error := global.GVA_DB.Update(&resource).Error
	if error != nil {
		return nil, error
	}
	return resource, nil
}

func AddResource(resource model.Resource) (*model.Resource, error) {
	error := global.GVA_DB.Save(&resource).Error
	if error != nil {
		return nil, error
	}
	return &resource, nil
}

func DeleteResource(resourceId int) error {
	var resource = new(model.Resource)
	resource.Delete = true
	resource.Id = resourceId
	error := global.GVA_DB.Update(&resource).Error
	if error != nil {
		return error
	}
	return nil
}

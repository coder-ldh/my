package service

import (
	"github.com/jinzhu/gorm"
	"my/global"
	"my/model"
)

func GetAdminById(adminId int) (*model.Admin, error) {
	var admin *model.Admin
	error := global.GVA_DB.Find(&admin, adminId).Error
	if error != nil && error != gorm.ErrRecordNotFound {
		return nil, error
	}
	return admin, nil
}

func GetAdminList(pageNum int, pageSize int) ([]*model.Admin, error) {
	var admins []*model.Admin
	error := global.GVA_DB.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&admins).Error
	if error != nil && error != gorm.ErrRecordNotFound {
		return nil, error
	}
	return admins, nil
}

func GetAdminByLoginName(LoginName string) (*model.Admin, error) {
	var admin = new(model.Admin)
	error := global.GVA_DB.Where("login_name = ?", LoginName).First(&admin).Error
	if error != nil && error == gorm.ErrRecordNotFound {
		return nil, error
	}
	return admin, nil
}

func UpdateAdmin(model.Admin) (*model.Admin, error) {
	var admin *model.Admin
	error := global.GVA_DB.Update(&admin).Error
	if error != nil {
		return nil, error
	}
	return admin, nil
}

func AddAdmin(admin model.Admin) (*model.Admin, error) {
	error := global.GVA_DB.Save(&admin).Error
	if error != nil {
		return nil, error
	}
	return &admin, nil
}

func DeleteAdmin(adminId int) error {
	var admin = new(model.Admin)
	admin.Delete = true
	admin.Id = adminId
	error := global.GVA_DB.Update(&admin).Error
	if error != nil {
		return error
	}
	return nil
}

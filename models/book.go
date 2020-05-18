package models

import (
	"github.com/jinzhu/gorm"

)

type Book struct {
	Id int `gorm:"size:11;primary_key;AUTO_INCREMENT;not null" json:"id"`
	Name string `gorm:"type:varchar(255);"column:name" json:"name"`
}

func Books()([]*Book,error) {
	println(DB)
	println(name)
	var books []*Book
	error := DB.Find(&books).Error
	if error != nil && error != gorm.ErrRecordNotFound {
		return nil, error
	}
	return books, nil
}

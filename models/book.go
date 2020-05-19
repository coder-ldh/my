package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	orm "my/database"
	"strconv"
)

type Book struct {
	Id         int    `gorm:"size:11;primary_key;AUTO_INCREMENT;not null" json:"id"`
	BookName   string `gorm:"type:varchar(255);"column:book_name" json:"bookName"`
	BookIntro  string `gorm:"type:varchar(255);"column:book_intro" json:"bookIntro"`
	BookAuthor string `gorm:"type:varchar(255);"column:book_author" json:"bookAuthor"`
}

func Books() ([]*Book, error) {
	var books []*Book
	error := orm.DB.Find(&books).Error
	if error != nil && error != gorm.ErrRecordNotFound {
		return nil, error
	}
	return books, nil
}

func BookMysqlToEs() {
	var books []*Book
	error := orm.DB.Find(&books).Error
	if error != nil && error != gorm.ErrRecordNotFound {
		fmt.Errorf("BookMysqlToEs（） 查询返回失败")
		return
	}
	for i := range books {
		orm.Es.Index().Index("book").Id(strconv.Itoa(books[i].Id)).BodyJson(books[i])
	}
}

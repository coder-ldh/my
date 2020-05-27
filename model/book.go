package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic/v7"
	"my/global"
	"strconv"
)

type Book struct {
	Id           int    `gorm:"size:11;primary_key;AUTO_INCREMENT;not null" json:"id"`
	BookName     string `gorm:"type:varchar(255);"column:book_name" json:"bookName"`
	BookIntro    string `gorm:"type:varchar(255);"column:book_intro" json:"bookIntro"`
	BookAuthor   string `gorm:"type:varchar(255);"column:book_author" json:"bookAuthor"`
	BookImageUrl string `gorm:"type:varchar(255);"column:book_image_url" json:"bookImageUrl"`
}

func Books(pageNum int, pageSize int) ([]*Book, error) {
	var books []*Book
	error := global.GVA_DB.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&books).Error
	if error != nil && error != gorm.ErrRecordNotFound {
		return nil, error
	}
	return books, nil
}

func GetBookByIdFromEs(bookId string) (*Book, error) {
	var book *Book
	getResult, err := global.GVA_ES.Get().Index("book").Id(bookId).Do(context.Background())
	if err != nil {
		return nil, err
	}
	json.Unmarshal(getResult.Source, &book)
	return book, nil
}

func BookMysqlToEs() (err error) {
	var books []*Book
	error := global.GVA_DB.Find(&books).Error
	if error != nil && error != gorm.ErrRecordNotFound {
		fmt.Errorf("BookMysqlToEs（） 查询返回失败")
		return error
	}
	var bulk = global.GVA_ES.Bulk()
	for i := range books {
		indexRequest := elastic.NewBulkIndexRequest().Index("book").Id(strconv.Itoa(books[i].Id)).Doc(books[i])
		bulk.Add(indexRequest)
	}
	bulk.Do(context.Background())
	return nil
}

package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic/v7"
	"my/global"
	"my/model"
	"strconv"
)

func Books(pageNum int, pageSize int) ([]*model.Book, error) {
	var books []*model.Book
	error := global.GVA_DB.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&books).Error
	if error != nil && error != gorm.ErrRecordNotFound {
		return nil, error
	}
	return books, nil
}

func GetBookByIdFromEs(bookId string) (*model.Book, error) {
	var book *model.Book
	getResult, err := global.GVA_ES.Get().Index("book").Id(bookId).Do(context.Background())
	if err != nil {
		return nil, err
	}
	json.Unmarshal(getResult.Source, &book)
	return book, nil
}

func BookMysqlToEs() (err error) {
	var books []*model.Book
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

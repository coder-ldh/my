package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic/v7"
	"my/global"
	"strconv"
)

type Section struct {
	Id             int    `gorm:"size:20;primary_key;AUTO_INCREMENT;not null" json:"id"`
	BookId         int    `gorm:"type:bigint(20);"column:book_id" json:"bookId"`
	SectionTitle   string `gorm:"type:varchar(255);"column:section_title" json:"sectionTitle"`
	SectionSeq     int    `gorm:"type:bigint(20);"column:section_seq" json:"sectionSeq"`
	SectionContent string `gorm:"type:text;"column:section_content" json:"sectionContent"`
}

func Sections(pageNum int, pageSize int) ([]*Section, error) {
	var s []*Section
	err := global.GVA_DB.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&s).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return s, nil
}

func SectionMysqlToEs() (err error) {
	var section []*Section
	error := global.GVA_DB.Find(&section).Error
	if error != nil && error != gorm.ErrRecordNotFound {
		fmt.Errorf("SectionMysqlToEs（） 查询返回失败")
		return error
	}
	var bulk = global.GVA_BulkProcessor
	for i := range section {
		indexRequest := elastic.NewBulkIndexRequest().Index("section").Id(strconv.Itoa(section[i].Id)).Doc(section[i])
		bulk.Add(indexRequest)
	}
	return nil
}

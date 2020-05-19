package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic/v7"
	orm "my/database"
	"strconv"
)

type Section struct {
	Id             int    `gorm:"size:20;primary_key;AUTO_INCREMENT;not null" json:"id"`
	BookId         int    `gorm:"type:bigint(20);"column:book_id" json:"bookId"`
	SectionTitle   string `gorm:"type:varchar(255);"column:section_title" json:"sectionTitle"`
	SectionSeq     int    `gorm:"type:bigint(20);"column:section_seq" json:"sectionSeq"`
	SectionContent string `gorm:"type:text;"column:section_content" json:"sectionContent"`
}

func Sections() ([]*Section, error) {
	var s []*Section
	err := orm.DB.Find(s).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return s, nil
}

func SectionMysqlToEs() (err error) {
	var section []*Section
	error := orm.DB.Find(&section).Error
	if error != nil && error != gorm.ErrRecordNotFound {
		fmt.Errorf("SectionMysqlToEs（） 查询返回失败")
		return error
	}
	var bulk = orm.BulkService
	for i := range section {
		indexRequest := elastic.NewBulkIndexRequest().Index("section").Id(strconv.Itoa(section[i].Id)).Doc(section[i])
		bulk.Add(indexRequest)
	}
	return nil
}

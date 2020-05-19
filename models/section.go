package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	orm "my/database"
	"strconv"
)

type Section struct {
	Id             int    `gorm:"size:20;primary_key;AUTO_INCREMENT;not null" json:"id"`
	BookId         int    `gorm:"type:bigint(20);"column:book_id" json:"bookId"`
	SectionTitle   string `gorm:"type:varchar(255);"column:section_title" json:"sectionTitle"`
	SectionSeq     string `gorm:"type:varchar(255);"column:section_seq" json:"sectionSeq"`
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

func SectionMysqlToEs() {
	var s []*Section
	error := orm.DB.Find(&s).Error
	if error != nil && error != gorm.ErrRecordNotFound {
		fmt.Errorf("SectionMysqlToEs（） 查询返回失败")
		return
	}
	for i := range s {
		orm.Es.Index().Index("section").Id(strconv.Itoa(s[i].Id)).BodyJson(s[i])
	}
}

package models

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
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
	var s []*Section
	error := orm.DB.Find(&s).Error
	if error != nil && error != gorm.ErrRecordNotFound {
		fmt.Errorf("SectionMysqlToEs（） 查询返回失败")
		return error
	}
	for i := range s {
		_, err := orm.Es.Index().Index("section").Id(strconv.Itoa(s[i].Id)).BodyJson(s[i]).Do(context.Background())
		if err != nil {
			fmt.Sprintf(" SectionMysqlToEs Index err:%s", err.Error())
		}
	}
	return nil
}

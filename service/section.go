package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic/v7"
	"log"
	"my/global"
	"my/model"
	"strconv"
)

func Sections(pageNum int, pageSize int) ([]*model.Section, error) {
	var s []*model.Section
	err := global.GVA_DB.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&s).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return s, nil
}

func SectionMysqlToEs() (err error) {
	var section []*model.Section
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

func SectionListByBookId(bookId string, c *gin.Context) ([]*model.Section, error) {
	var sectionList []*model.Section
	searchSource := elastic.NewSearchSource()
	var include, exclude []string
	include = []string{"id", "BookId", "SectionTitle", "SectionSeq"}
	exclude = []string{"SectionContent"}
	searchSource.FetchSourceIncludeExclude(include, exclude)
	esQuery := elastic.NewMatchQuery("BookId", bookId)
	searchResult, err := global.GVA_ES.Search().Index("section").SearchSource(searchSource).Query(esQuery).Sort("SectionSeq", false).Do(c.Request.Context())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for _, hit := range searchResult.Hits.Hits {
		var section *model.Section
		json.Unmarshal(hit.Source, &section)
		sectionList = append(sectionList, section)
	}
	return sectionList, nil
}

func SectionById(sectionId int) (*model.Section, error) {
	s := new(model.Section)
	err := global.GVA_DB.First(&s, sectionId).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return s, nil
}

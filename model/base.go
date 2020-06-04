package model

import (
	"time"
)

type Base struct {
	Id int `gorm:"size:11;primary_key;AUTO_INCREMENT;not null" json:"id"`
	//是否删除
	Delete bool `json:"delete"`
	//乐观锁
	Revision *int `json:"revision"`
	//创建人
	CreatedBy int `json:"createdBy"`
	//创建时间
	CreatedTime time.Time `json:"createdTime"`
	//修改人
	UpdatedBy int `json:"updatedBy"`
	//修改时间
	UpdatedTime time.Time `json:"updatedTime"`
}

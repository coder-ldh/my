package request

import "my/model"

// Paging common input parameter structure
type PageInfo struct {
	PageNum  int `json:"pageNum" form:"pageNum"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

// Find by id structure
type GetById struct {
	Id float64 `json:"id" form:"id"`
}

type SectionResponse struct {
	Time string          `json:"time"`
	Hits string          `json:"hits"`
	Data []model.Section `json:"sections"`
}

type BookResponse struct {
	Time int64        `json:"time"`
	Hits int64        `json:"hits"`
	Data []model.Book `json:"books"`
}

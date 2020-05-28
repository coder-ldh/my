package model

type Book struct {
	Id           int    `gorm:"size:11;primary_key;AUTO_INCREMENT;not null" json:"id"`
	BookName     string `gorm:"type:varchar(255);"column:book_name" json:"bookName"`
	BookIntro    string `gorm:"type:varchar(255);"column:book_intro" json:"bookIntro"`
	BookAuthor   string `gorm:"type:varchar(255);"column:book_author" json:"bookAuthor"`
	BookImageUrl string `gorm:"type:varchar(255);"column:book_image_url" json:"bookImageUrl"`
}

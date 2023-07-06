package code_memory

import "time"

type Article struct {
	ID          int       `gorm:"primaryKey;column:id" json:"id"`
	Title       string    `gorm:"column:title" json:"title"`
	ArticleType string    `gorm:"column:article_type;" json:"article_type"`
	ColumnID    int       `gorm:"column:column_id" json:"column_id"`
	SubColumnID int       `gorm:"column:sub_column_id" json:"sub_column_id"`
	URL         string    `gorm:"column:URL;type:longtext" json:"URL"`
	Code        string    `gorm:"column:code;type:longtext" json:"code"`
	Create_time time.Time `gorm:"column:create_time" json:"create_time"`
}

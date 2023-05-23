package db

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
type Category struct {
	ID          int       `gorm:"primaryKey;column:id" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Create_time time.Time `gorm:"column:create_time" json:"create_time"`
}
type SubCategory struct {
	ID          int       `gorm:"primaryKey;column:id" json:"id"`
	Name        string    `gorm:"column:unique;not null;name" json:"name"`
	Category_id int       `gorm:"column:category_id" json:"category_id"`
	Create_time time.Time `gorm:"column:create_time" json:"create_time"`
}

package code_memory

import "time"

type CategoryInfo struct {
	Name    string `json:"name"`
	SubName string `json:"subName"`
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

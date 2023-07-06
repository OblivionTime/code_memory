package code_memory

type Info struct {
	ArticleTotal     int64             `json:"article_total"`
	CategoryTotal    int64             `json:"category_total"`
	SubCategoryTotal int64             `json:"subCategory_total"`
	CategoryList     []*CategoryDetail `json:"category_list"`
}
type CategoryDetail struct {
	ID       int           `json:"id"`
	Name     string        `json:"name"`
	Children []SubCategory `json:"children"`
}

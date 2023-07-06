package code_memory

import (
	"code_memory/global"
	"code_memory/model/code_memory"
	"code_memory/model/common"

	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

func (b *BaseApi) GetInfo(ctx *gin.Context) {
	info := &code_memory.Info{}
	//计算文章总数
	global.DB.Model(code_memory.Article{}).Count(&info.ArticleTotal)
	//计算栏目总数
	global.DB.Model(code_memory.Category{}).Count(&info.CategoryTotal)
	//计算标签总数
	global.DB.Model(code_memory.SubCategory{}).Count(&info.SubCategoryTotal)
	//获取各个栏目的标签信息
	var category_list []code_memory.Category
	global.DB.Model(code_memory.Category{}).Scan(&category_list)
	for _, cat := range category_list {
		detail := &code_memory.CategoryDetail{
			ID:       cat.ID,
			Name:     cat.Name,
			Children: make([]code_memory.SubCategory, 0),
		}
		global.DB.Model(code_memory.SubCategory{}).Where("category_id =?", detail.ID).Scan(&detail.Children)
		info.CategoryList = append(info.CategoryList, detail)
	}
	common.OkWithData(info, ctx)
}

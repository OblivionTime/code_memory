package code_memory

import (
	"code_memory/global"
	"code_memory/model/code_memory"
	"code_memory/model/common"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ArticleApi struct{}

//获取文章列表
func (a *ArticleApi) GetArticleList(ctx *gin.Context) {
	// page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	categoryId, _ := strconv.Atoi(ctx.DefaultQuery("category_id", "0"))
	sub_categoryId, _ := strconv.Atoi(ctx.DefaultQuery("sub_category_id", "0"))
	title := ctx.DefaultQuery("title", "")
	article_type := ctx.DefaultQuery("article_type", "")
	var article_list []code_memory.Article
	// 计算偏移量
	// offset := (page - 1) * 10
	var count int64
	query := global.DB.Model(&code_memory.Article{})
	if categoryId != 0 {
		query = query.Where("column_id=? and sub_column_id=?", categoryId, sub_categoryId)
	}
	if title != "" {
		query = query.Where("title LIKE ?", fmt.Sprintf(`%%%s%%`, title))
	}
	if article_type != "" {
		query = query.Where("article_type=?", article_type)
	}
	query.Count(&count)
	query.Order("create_time DESC").Find(&article_list)
	// 构造响应数据
	response := map[string]interface{}{
		// "page":  page,
		"total": count,
		"data":  article_list,
	}
	common.OkWithData(response, ctx)
}

//添加/修改文章
func (a *ArticleApi) CreateOrUpdateArticle(ctx *gin.Context) {
	var resd code_memory.Article
	if err := ctx.ShouldBindJSON(&resd); err != nil {
		common.FailWithMessage("参数错误", ctx)
		return
	}
	resd.Create_time = time.Now()
	var res *gorm.DB
	if resd.ID != 0 {
		res = global.DB.Model(&code_memory.Article{}).Where("id=?", resd.ID).Updates(&resd)
	} else {
		res = global.DB.Model(&code_memory.Article{}).Create(&resd)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
		common.FailWithMessage(res.Error.Error(), ctx)
		return
	}
	common.Ok(ctx)

}

//删除文章
func (a *ArticleApi) DeleteArticle(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		common.FailWithMessage("参数错误", ctx)
		return
	}
	result := global.DB.Model(&code_memory.Article{}).Delete(&code_memory.Article{}, id)
	if result.Error != nil {
		fmt.Println(result.Error)
		common.FailWithMessage(result.Error.Error(), ctx)
		return
	} else if result.RowsAffected == 0 {
		common.FailWithMessage("不存在当前数据", ctx)
		return
	}
	common.Ok(ctx)
}

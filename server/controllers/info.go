package controllers

import (
	"code_memory/db"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}
type Info struct {
	ArticleTotal     int               `json:"article_total"`
	CategoryTotal    int               `json:"category_total"`
	SubCategoryTotal int               `json:"subCategory_total"`
	CategoryList     []*CategoryDetail `json:"Category_list"`
}
type CategoryDetail struct {
	ID       int              `json:"id"`
	Name     string           `json:"name"`
	Children []db.SubCategory `json:"children"`
}

//获取信息
func GetInfo(ctx *gin.Context) {
	info := &Info{}
	//计算文章总数
	db.DB.Model(db.Article{}).Count(&info.ArticleTotal)
	//计算栏目总数
	db.DB.Model(db.Category{}).Count(&info.CategoryTotal)
	//计算标签总数
	db.DB.Model(db.SubCategory{}).Count(&info.SubCategoryTotal)
	//获取各个栏目的标签信息
	var category_list []db.Category
	db.DB.Model(db.Category{}).Scan(&category_list)
	for _, cat := range category_list {
		detail := &CategoryDetail{
			ID:       cat.ID,
			Name:     cat.Name,
			Children: make([]db.SubCategory, 0),
		}
		db.DB.Model(db.SubCategory{}).Where("category_id =?", detail.ID).Scan(&detail.Children)
		info.CategoryList = append(info.CategoryList, detail)
	}
	resp := &Response{
		Code: 0,
		Msg:  "success",
		Data: info,
	}
	ctx.JSON(200, resp)
}

//获取文章列表
func GetArticleList(ctx *gin.Context) {
	// page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	categoryId, _ := strconv.Atoi(ctx.DefaultQuery("category_id", "0"))
	sub_categoryId, _ := strconv.Atoi(ctx.DefaultQuery("sub_category_id", "0"))
	title := ctx.DefaultQuery("title", "")
	article_type := ctx.DefaultQuery("article_type", "")
	var article_list []db.Article
	// 计算偏移量
	// offset := (page - 1) * 10
	var count int64
	query := db.DB.Model(&db.Article{})
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

	ctx.JSON(200, response)
}

//添加/修改文章
func CreateOrUpdateArticle(ctx *gin.Context) {
	resp := &Response{
		Code: 0,
		Msg:  "success",
	}
	var resd db.Article
	if err := ctx.ShouldBindJSON(&resd); err != nil {
		ctx.JSON(200, &Response{
			Code: 4001,
			Msg:  "参数有误",
		})
		return
	}
	resd.Create_time = time.Now()
	var res *gorm.DB
	if resd.ID != 0 {
		res = db.DB.Model(db.Article{}).Where("id=?", resd.ID).Updates(&resd)
	} else {
		res = db.DB.Model(db.Article{}).Create(&resd)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
		resp.Code = 4001
		resp.Msg = "error"
		ctx.JSON(200, resp)
		return
	}
	ctx.JSON(200, resp)
}

//删除文章
func DeleteArticle(ctx *gin.Context) {
	id := ctx.Query("id")
	resp := &Response{
		Code: 0,
		Msg:  "success",
	}
	if id == "" {
		ctx.JSON(200, &Response{
			Code: 4001,
			Msg:  "参数有误",
		})
		return
	}
	result := db.DB.Model(&db.Article{}).Delete(&db.Article{}, id)
	if result.Error != nil {
		fmt.Println(result.Error)
		resp.Code = 4001
		resp.Msg = "删除失败"
		ctx.JSON(200, resp)
		return
	} else if result.RowsAffected == 0 {
		// 记录不存在
		resp.Code = 4001
		resp.Msg = "不存在当前数据"
		ctx.JSON(200, resp)
		return
	}
	ctx.JSON(200, resp)
}

type CategoryInfo struct {
	Name    string `json:"name"`
	SubName string `json:"subName"`
}

//添加栏目
func CreateCategory(ctx *gin.Context) {
	resp := &Response{
		Code: 0,
		Msg:  "success",
	}
	var resd CategoryInfo
	if err := ctx.ShouldBindJSON(&resd); err != nil {
		fmt.Println(err)
		ctx.JSON(200, &Response{
			Code: 4001,
			Msg:  "参数有误",
		})
		return
	}
	category := &db.Category{
		Name:        resd.Name,
		Create_time: time.Now(),
	}
	//判断当前栏目是否已经存在,如果存在不插入,不存在则插入
	res := db.DB.Model(db.Category{}).FirstOrCreate(&category, db.Category{Name: category.Name})

	if res.Error != nil {
		resp.Code = 4001
		resp.Msg = "插入失败"
		ctx.JSON(200, resp)
		return
	}
	//插入子栏目,先判断子栏目是否存在,存在则返回错误,不存在则插入
	subCategory := &db.SubCategory{
		Name:        resd.SubName,
		Category_id: category.ID,
		Create_time: time.Now(),
	}
	res = db.DB.Model(db.SubCategory{}).Create(&subCategory)
	if res.Error != nil {

		resp.Code = 4001
		resp.Msg = "子栏目已存在!!!"
		ctx.JSON(200, resp)
		return
	}
	ctx.JSON(200, resp)
}

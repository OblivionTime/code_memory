package code_memory

import (
	"code_memory/global"
	"code_memory/model/code_memory"
	"code_memory/model/common"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type CategoryApi struct{}

//添加栏目
func (c *CategoryApi) CreateCategory(ctx *gin.Context) {

	var resd code_memory.CategoryInfo
	if err := ctx.ShouldBindJSON(&resd); err != nil {
		fmt.Println(err)
		common.FailWithMessage("参数错误", ctx)
		return
	}
	category := &code_memory.Category{
		Name:        resd.Name,
		Create_time: time.Now(),
	}
	//判断当前栏目是否已经存在,如果存在不插入,不存在则插入
	res := global.DB.Model(code_memory.Category{}).FirstOrCreate(&category, code_memory.Category{Name: category.Name})

	if res.Error != nil {
		common.FailWithMessage("插入失败", ctx)
		return
	}
	//插入子栏目,先判断子栏目是否存在,存在则返回错误,不存在则插入
	subCategory := &code_memory.SubCategory{
		Name:        resd.SubName,
		Category_id: category.ID,
		Create_time: time.Now(),
	}
	res = global.DB.Model(code_memory.SubCategory{}).Create(&subCategory)
	if res.Error != nil {
		common.FailWithMessage("子栏目已存在", ctx)
		return
	}
	common.Ok(ctx)
}

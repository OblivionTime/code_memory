

<template>
  <div class="container">
    <sidebar :menu="menu" :category_total="category_total" :article_total="article_total"
      :subCategory_total="subCategory_total" @switch_article="switch_article"></sidebar>
    <div class="content">
      <div class="header">
        <div class="search">
          <img src="./assets/research.svg" alt="" class="form-icon">
          <input v-model="search" type="text" class="search-input" placeholder="搜索文章" @input="loadArticleList()">
        </div>
        <div style="display: flex; align-items: center">
          <div style="display: flex; align-items: center;flex: 1;">
            <div style="
              text-align: center;
              background-color: #f5f7fa;
              line-height: 29px;
              width: 80px;
              color: #909399;
              border-radius: 4px 0 0px 4px;
              padding: 0 20px;
              border: 1px solid #dcdfe6;
              font-size: 14px;
            ">
              文章类型
            </div>
            <el-select v-model="article_type" placeholder="请选择" style="width: 120px;">
              <el-option label="所有" value="所有"> </el-option>
              <el-option label="文章" value="文章"> </el-option>
              <el-option label="代码" value="代码"> </el-option>
            </el-select>
          </div>
        </div>
        <div class="btn" style="display: flex;    justify-content: center;    align-items: center;">
          <div>
            <button @click="showAdd">发布文章</button>
          </div>
          <div>
            <button @click="showCategory">创建栏目</button>
          </div>
        </div>
      </div>
      <div class="article-lists">
        <div class="article" v-for="item in article_list" :key="item.id">
          <div>
            <h2 class="article-title" v-if="item.article_type == '代码'">{{ item.title }}</h2>
            <h2 class="article-title article-URL" v-else @click="openURL(item.URL)">{{ item.title }}</h2>
            <div class="edit" @click="showEdit(item)">
              <img src="./assets/edit.svg" style="width: 20px;">
              编辑文章
            </div>
            <div class="delete" @click="deleteArticle(item.id)">
              <img src="./assets/delete.svg" style="width: 20px;">
              删除文章
            </div>
          </div>

          <div class="article-content" v-if="item.article_type == '代码'">
            <div class="hljs-container">
              <highlightjs :autodetect="true" :code="item.code"></highlightjs>
            </div>

            <div class="copy" @click="copyCode(item.code)">
              <img src="./assets/copy.svg" style="width: 20px;">
              复制代码
            </div>
            <div class="more" @click="showMore(item.id)" style="display: none;">
              <span>显示更多</span>
            </div>
          </div>

        </div>
      </div>


    </div>
    <el-dialog v-model="showAddArticle" :title="articleTitle" width="35%" draggable :close-on-click-modal="false"
      align-center>
      <el-form :model="formData" label-width="120px" label-position="left">
        <el-form-item label="标题:">
          <el-input v-model="formData.title" />
        </el-form-item>
        <el-form-item label="栏目:">
          <el-cascader v-model="formData.tag" :options="options" @change="handleChange" style="width: 100%;" />
        </el-form-item>
        <el-form-item label="文章类型:">
          <el-select v-model="formData.article_type" style="width: 100%;">
            <el-option label="代码" value="代码" />
            <el-option label="文章" value="文章" />
          </el-select>
        </el-form-item>
        <el-form-item label="URL:" v-if="formData.article_type == '文章'">
          <el-input v-model="formData.URL" />
        </el-form-item>
        <el-form-item label="代码:" v-if="formData.article_type == '代码'">
          <el-input v-model="formData.code" :autosize="{ minRows: 1, maxRows: 4 }" type="textarea" placeholder="请输入代码"
            resize="none" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="showAddArticle = false">取消</el-button>
          <el-button type="primary" @click="HandlerArticle">
            {{ articleTitle == '发布文章' ? '添加' : '修改' }}
          </el-button>
        </span>
      </template>
    </el-dialog>
    <el-dialog v-model="showAddColumn" title="添加栏目" width="35%" draggable :close-on-click-modal="false" align-center>
      <el-form :model="colFormData" label-width="120px" label-position="left">
        <el-form-item label="栏目:">
          <el-autocomplete v-model="colFormData.name" :fetch-suggestions="querySearch" clearable
            placeholder="Please Input" style="width:100%" />
        </el-form-item>
        <el-form-item label="子栏目:">
          <el-input v-model="colFormData.subName" />
        </el-form-item>

      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="showAddColumn = false">取消</el-button>
          <el-button type="primary" @click="SubmitCategory">
            添加
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup>
import sidebar from './components/sidebar.vue';
import { addLine, copyCode } from '@/utils/dom.js';
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, onMounted, reactive, nextTick } from 'vue';
import { getInfo, getArticleList, postCreateArticle, postUpdateArticle, getDeleteArticle, postCreateCategory } from '@/api/info.js'
/**
 * 发布/修改文章
 */
let showAddArticle = ref(false)
let articleTitle = ref("发布文章")
let formData = ref({ title: "", code: "", URL: "", tag: reactive([]), column_id: "", sub_column_id: "", article_type: "代码" })
let options = ref([])
function handleChange(tag) {
  formData.value.column_id = tag[0]
  formData.value.sub_column_id = tag[1]
}
//发布文章/修改文章
function HandlerArticle() {
  let formData2 = formData.value
  if (!(formData2.title && formData2.column_id && formData2.sub_column_id)) {
    return ElMessage({
      message: '请填写完整信息',
      type: 'warning',
    })
  }
  if (formData2.article_type == '代码' && !formData2.code) {
    return ElMessage({
      message: '请填写代码',
      type: 'warning',
    })
  } else if (formData2.article_type == '文章' && !formData2.URL) {
    return ElMessage({
      message: '请填写URL',
      type: 'warning',
    })
  }
  if (articleTitle.value == '发布文章') {
    postCreateArticle(formData2)
      .then((res) => {
        if (res.code == 0) {
          ElMessage({
            message: '操作成功',
            type: 'success',
          })
          showAddArticle.value = false
          loadInfo()
          loadArticleList()
        } else {
          ElMessage({
            message: res.message,
            type: 'error',
          })
        }
      })
  } else {
    postUpdateArticle(formData2)
      .then((res) => {
        if (res.code == 0) {
          ElMessage({
            message: '操作成功',
            type: 'success',
          })
          showAddArticle.value = false
          loadInfo()
          loadArticleList()
        } else {
          ElMessage({
            message: res.message,
            type: 'error',
          })
        }
      })
  }
}
//显示添加文章窗口
function showAdd() {
  formData.value = { title: "", code: "", URL: "", tag: reactive([]), column_id: "", sub_column_id: "", article_type: "代码" }
  showAddArticle.value = true
  articleTitle.value = '发布文章'
}
//显示编辑文章窗口
function showEdit(row) {
  formData.value = JSON.parse(JSON.stringify(row))
  formData.value.tag = [formData.value.column_id, formData.value.sub_column_id]
  showAddArticle.value = true
  articleTitle.value = '编辑文章'
}
//删除文章
function deleteArticle(id) {
  ElMessageBox.confirm('此操作将永久删除该文章, 是否继续?')
    .then(() => {
      console.log(id);
      getDeleteArticle({ id: id })
        .then((res) => {
          if (res.code == 0) {
            ElMessage({
              message: '删除成功',
              type: 'success',
            })
            loadInfo()
            loadArticleList()
          } else {
            ElMessage({
              message: res.message,
              type: 'error',
            })
          }
        })
    })
    .catch(() => {
      // catch error
    })
}
//搜索文章
let search = ref("")
let article_type = ref("")
onMounted(() => {
  loadInfo()
  loadArticleList()
})

function showMore(id) {
  let index = code_dict.value[id]
  let block = document.querySelectorAll('pre')[index]
  block.style.overflow = "auto"
  block.parentNode.parentNode.querySelectorAll('.more')[0].style.display = "none"
}
/**
 * 加载文章
 */
let article_list = ref([])
let category_id = ref(0)
let sub_category_id = ref(0)
let code_dict = ref({})
function switch_article(category_id2, sub_category_id2) {
  category_id.value = category_id2
  sub_category_id.value = sub_category_id2
  loadArticleList()
}

function loadArticleList() {
  getArticleList({ category_id: category_id.value, sub_category_id: sub_category_id.value, "title": search.value, "article_type": article_type.value })
    .then((res) => {
      article_list.value = res.data
      nextTick(() => {
        let code_list = res.data.filter(item => item.article_type == '代码')
        let blocks = document.querySelectorAll('pre');
        blocks.forEach((block, index) => {
          code_dict.value[code_list[index].id] = index
          addLine(block)
          if (block.scrollHeight > 600) {
            block.style.overflow = "hidden"
            block.parentNode.parentNode.querySelectorAll('.more')[0].style.display = "flex"
          }
        })
      })
    })
}
//打开新网页
function openURL(url) {
  window.open(url);
}
/**
 * 基本数据
 */
const getImgUrl = (name) => {
  return new URL(name, import.meta.url).href
}

let article_total = ref(0)
let category_total = ref(0)
let subCategory_total = ref(0)
let menu = ref([])

function loadInfo() {
  getInfo()
    .then((res) => {
      if (res.code == 0) {
        article_total.value = res.data.article_total;
        category_total.value = res.data.category_total;
        subCategory_total.value = res.data.subCategory_total;
        let Category_list = []
        let options2 = []
        let restaurants2 = []
        for (const item of res.data.Category_list) {
          let value = { name: item.name, id: item.id, icon: getImgUrl("./assets/frontend.svg"), children: item.children }
          let child = item.children.map((t) => ({ label: t.name, value: t.id }))
          let option = { label: item.name, value: item.id, children: child }
          restaurants2.push({ value: item.name })
          Category_list.push(value)
          options2.push(option)
        }
        menu.value = Category_list
        options.value = options2
        restaurants.value = restaurants2
      }
    })
}
/**
 * 栏目相关
 */
let showAddColumn = ref(false)
let colFormData = ref({ name: "", subName: "", })
let restaurants = ref([])
function showCategory() {
  showAddColumn.value = true
  colFormData.value = { name: "", subName: "", }
}
function SubmitCategory() {
  let formData2 = colFormData.value
  if (!(formData2.name && formData2.subName)) {
    return ElMessage({
      message: '请填写完整信息',
      type: 'warning',
    })
  }
  postCreateCategory(formData2)
    .then((res) => {
      if (res.code == 0) {
        ElMessage({
          message: '操作成功',
          type: 'success',
        })
        showAddColumn.value = false
        loadInfo()
      } else {
        ElMessage({
          message: res.message,
          type: 'error',
        })
      }
    })

}
const querySearch = (queryString, cb) => {
  const results = queryString
    ? restaurants.value.filter(createFilter(queryString))
    : restaurants.value
  // call callback function to return suggestions
  cb(results)
}
const createFilter = (queryString) => {
  return (restaurant) => {
    return (
      restaurant.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0
    )
  }
}
</script>
<style lang="scss" scoped>
@import "./code.scss";

.container {
  display: flex;
  align-items: center;

  .content {
    overflow: auto;
    width: 60vw;
    height: 98vh;
    border-radius: 10px;
    background: #fff;
    box-sizing: border-box;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, .5);

    .header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      width: 94%;
      box-sizing: border-box;
      padding: 30px 40px;
      border-radius: 20px;
      margin: 20px auto;
      box-shadow: 0 2px 12px 0 rgba(0, 0, 0, .5);

      .search {
        position: relative;
        flex: 1;

        .form-icon {
          position: absolute;
          top: 4px;
          left: 10px;
          width: 20px;
        }

        .search-input {
          // width: 300px;
          border: 1px solid #5ccdf8;
          border-radius: 10px;
          padding: 5px 0 5px 40px;
          box-sizing: border-box;
          font-size: 16px;
          outline: none;
          margin-bottom: 5px;

          &:focus {
            border-bottom-color: #5ccdf8;
          }
        }
      }

      .btn {
        flex: 1;
        display: flex;
        justify-content: center;
        align-items: center;
        gap: 10px;
        div {
          flex: 1;
        }

        button {
          width: 100%;
          cursor: pointer;
          height: 35px;
          margin: 0 10px;
          border-radius: 10px;
          color: #fff;
          border: 0;
          background: #409eff;
          box-shadow: 0 0 1px 0 rgba(0, 0, 0, .5);
        }
      }

    }

    .article-lists {
      position: relative;



      .article {
        width: 94%;
        box-sizing: border-box;
        padding: 30px 40px;
        border-radius: 20px;
        margin: 20px auto;
        box-shadow: 0 2px 12px 0 rgba(0, 0, 0, .5);
        position: relative;
        text-align: center;
        background: rgba($color: #fff, $alpha: 0.5);

        .edit {
          width: 100px;
          letter-spacing: 2px;
          height: 40px;
          position: absolute;
          top: 0px;
          right: 90px;
          color: white;
          font-size: 11px;
          display: flex;
          align-items: center;
          cursor: pointer;
          color: #409eff;
          z-index: 10;
        }

        .delete {
          width: 100px;
          letter-spacing: 2px;
          height: 40px;
          position: absolute;
          top: 0px;
          right: 0px;
          color: white;
          font-size: 11px;
          display: flex;
          align-items: center;
          cursor: pointer;
          color: red;
          z-index: 10;
        }

        .article-title {
          width: 100%;
          white-space: nowrap;
          overflow: hidden;
          text-overflow: ellipsis;
          margin-bottom: 20px;
        }

        .article-URL {
          &:hover {
            text-decoration: underline;
            color: #409eff;
            cursor: pointer;
          }
        }

        .article-content {
          // max-height: 600px;
          // overflow: auto;
          border-radius: 20px;
          position: relative;
        }



        .copy {
          width: 100px;
          letter-spacing: 2px;
          height: 40px;
          position: absolute;
          top: -5px;
          color: white;
          right: 0px;
          font-size: 11px;
          display: flex;
          align-items: center;
          cursor: pointer;
        }

        .more {
          align-items: center;
          justify-content: center;
          position: absolute;
          bottom: 0px;
          width: 100%;
          // border-bottom-left-radius: 20px;
          // border-bottom-right-radius: 20px;
          height: 50px;
          font-size: 18px;
          z-index: 999;
          color: white;
          background: rgba(0, 0, 0, 0.5);

          span:hover {
            cursor: pointer;
            color: burlywood;
          }
        }

      }
    }
  }

}
</style>

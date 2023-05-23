import request from '@/utils/request'
//获取基本信息
export function getInfo() {
    return request({
        url: '/info',
        method: 'get',
    })
}
//获取文章列表
export function getArticleList(params) {
    return request({
        url: '/list_article',
        method: 'get',
        params
    })
}
//添加文章
export function postCreateArticle(data){
    return request({
        url: '/add_article',
        method: 'post',
        data
    })
}
//修改文章
export function postUpdateArticle(data){
    return request({
        url: '/update_article',
        method: 'post',
        data
    })
}
//删除文章
export function getDeleteArticle(params) {
    return request({
        url: '/delete_article',
        method: 'get',
        params
    })
}
//添加栏目
export function postCreateCategory(data){
    return request({
        url: '/add_category',
        method: 'post',
        data
    })
}
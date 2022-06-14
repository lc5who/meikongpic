import service from '@/utils/request'

// @Tags Author
// @Summary 创建Author
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Author true "创建Author"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /author/createAuthor [post]
export const createAuthor = (data) => {
  return service({
    url: '/author/createAuthor',
    method: 'post',
    data
  })
}

// @Tags Author
// @Summary 删除Author
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Author true "删除Author"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /author/deleteAuthor [delete]
export const deleteAuthor = (data) => {
  return service({
    url: '/author/deleteAuthor',
    method: 'delete',
    data
  })
}

// @Tags Author
// @Summary 删除Author
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Author"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /author/deleteAuthor [delete]
export const deleteAuthorByIds = (data) => {
  return service({
    url: '/author/deleteAuthorByIds',
    method: 'delete',
    data
  })
}

// @Tags Author
// @Summary 更新Author
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Author true "更新Author"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /author/updateAuthor [put]
export const updateAuthor = (data) => {
  return service({
    url: '/author/updateAuthor',
    method: 'put',
    data
  })
}

// @Tags Author
// @Summary 用id查询Author
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Author true "用id查询Author"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /author/findAuthor [get]
export const findAuthor = (params) => {
  return service({
    url: '/author/findAuthor',
    method: 'get',
    params
  })
}

// @Tags Author
// @Summary 分页获取Author列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Author列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /author/getAuthorList [get]
export const getAuthorList = (params) => {
  return service({
    url: '/author/getAuthorList',
    method: 'get',
    params
  })
}

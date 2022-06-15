import service from '@/utils/request'

// @Tags Authors
// @Summary 创建Authors
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Authors true "创建Authors"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authors/createAuthors [post]
export const createAuthors = (data) => {
  return service({
    url: '/authors/createAuthors',
    method: 'post',
    data
  })
}

// @Tags Authors
// @Summary 删除Authors
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Authors true "删除Authors"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /authors/deleteAuthors [delete]
export const deleteAuthors = (data) => {
  return service({
    url: '/authors/deleteAuthors',
    method: 'delete',
    data
  })
}

// @Tags Authors
// @Summary 删除Authors
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Authors"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /authors/deleteAuthors [delete]
export const deleteAuthorsByIds = (data) => {
  return service({
    url: '/authors/deleteAuthorsByIds',
    method: 'delete',
    data
  })
}

// @Tags Authors
// @Summary 更新Authors
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Authors true "更新Authors"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /authors/updateAuthors [put]
export const updateAuthors = (data) => {
  return service({
    url: '/authors/updateAuthors',
    method: 'put',
    data
  })
}

// @Tags Authors
// @Summary 用id查询Authors
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Authors true "用id查询Authors"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /authors/findAuthors [get]
export const findAuthors = (params) => {
  return service({
    url: '/authors/findAuthors',
    method: 'get',
    params
  })
}

// @Tags Authors
// @Summary 分页获取Authors列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Authors列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authors/getAuthorsList [get]
export const getAuthorsList = (params) => {
  return service({
    url: '/authors/getAuthorsList',
    method: 'get',
    params
  })
}

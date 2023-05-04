import service from '../utils/request'


  // @Tags SysApi
  // @Summary 获取权限客户列表
  // @Security ApiKeyAuth
  // @accept application/json
  // @Produce application/json
  // @Param data body modelInterface.PageInfo true "获取权限客户列表"
  // @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
  // @Router /customer/customerList [get]
  export const getCfDomainList = () => {
    return service({
      url: '/cloudflare/all',
      method: 'get',
    })
  }
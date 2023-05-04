import service from '../utils/request'

// 创建cf api
export const createCfApi = (data) => {
    return service({
        url: '/apiKeyManager/cloudflare',
        method: 'post',
        data
    })
}

//获取cf api list
export const getCfApiList = () => {
    return service({
        url: '/apiKeyManager/cloudflareList',
        method: 'get',
    })
}

export const deleteCfApi = (data) => {
    return service({
        url: '/apiKeyManager/cloudflare',
        method: 'delete',
        data
    })
}

export const getCfApiById = (params) => {
    return service({
        url: '/apiKeyManager/cloudflare',
        method: 'get',
        params
    })
}

export const updateCfApiForm = (data) => {
    return service({
        url: '/apiKeyManager/cloudflare',
        method: 'put',
        data
    })
}
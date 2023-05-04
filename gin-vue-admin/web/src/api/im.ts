import service from '../utils/request'




export const getImList = () => {
    return service({
        url: '/im',
        method: 'get',
    })
}

export const deleteIm = (data) => {
    return service({
        url: '/im',
        method: 'delete',
        data
    })
}

export const getImById = (params) => {
    return service({
        url: '/im/single',
        method: 'get',
        params
    })
}

export const updateIm = (data) => {
    return service({
        url: '/im',
        method: 'put',
        data
    })
}

export const createIm = (data) => {
    return service({
        url: '/im',
        method: 'post',
        data
    })
}
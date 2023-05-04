import service from '../utils/request'


export const getImLine = (params) => {
    return service({
        url: '/im/lines',
        method: 'get',
        params
    })
}

export const changeLine = (data,id) => {
    return service({
        url: '/im/lines/'+id,
        method: 'put',
        data
    })
}
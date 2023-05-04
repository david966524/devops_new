import service from '../utils/request'


export const getAirborneList = () => {
    return service({
        url: '/airborne',
        method: 'get',
    })
}

export const deleteAirborne = (data) => {
    return service({
        url: '/airborne',
        method: 'delete',
        data
    })
}

export const getAirborneById = (params) => {
    return service({
        url: '/airborne/single',
        method: 'get',
        params
    })
}

export const updateAirborne = (data) => {
    return service({
        url: '/airborne',
        method: 'put',
        data
    })
}


export const createAirborne = (data) => {
    return service({
        url: '/airborne',
        method: 'post',
        data
    })
}
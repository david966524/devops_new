import service from '../utils/request'




export const getAsiaCloudVhost = () => {
    return service({
        url: '/asiacloud/vhost',
        method: 'get',
    })
}

export const getAsiaCloudDomainList = (vhost) => {
    return service({
        url: '/asiacloud/' + vhost,
        method: 'get',
    })
}

export const deleteAsiaCloudDomain = (data) => {
    return service({
        url: '/asiacloud',
        method: 'delete',
        data
    })
}

export const addAsiaCloudDomain = (data) => {
    return service({
        url: '/asiacloud',
        method: 'post',
        data
    })
}


export const updateAsiaCloudDomain = (data) => {
    return service({
        url: '/asiacloud',
        method: 'put',
        data
    })
}


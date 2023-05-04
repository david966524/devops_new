import service from '../utils/request'


export const getjobs = () => {
    return service({
        url: '/jenkins/jobs',
        method: 'get'
    })
}


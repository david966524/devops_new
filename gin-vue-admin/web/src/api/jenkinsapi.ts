import service from '../utils/request'


export const getjobs = () => {
    return service({
        url: '/jenkins/jobs',
        method: 'get'
    })
}

export const getBuildHistory = (params) => {
    return service({
        url: '/jenkins/jobs/build',
        method: 'get',
        params
    })
}

export const getBuildInfo = (params) => {
    return service({
        url: '/jenkins/jobs/build/info',
        method: 'get',
        params
    })
}

export const getConsoleOut = (params) => {
    return service({
        url: '/jenkins/jobs/build/info/consoleout',
        method: 'get',
        params
    })
}



export const getBuildParameters = (params) => {
    return service({
        url: '/jenkins/jobs/build/parameters',
        method: 'get',
        params
    })
}


export const onlyBuildJob = (data) => {
    return service({
        url: '/jenkins/jobs/build',
        method: 'post',
        data
    })
}
import service from '../utils/request'

export const getdnsList = (params) => {
    return service({
      url: '/dns/all',
      method: 'get',
      params
    })
  }

export const getdnsrecord= (params)=>{
    return service({
        url: '/dns/record',
        method: 'get',
        params
      })
}

export const adddnsrecord=(data)=>{
    return service({
        url: '/dns/record',
        method: 'post',
        data
      })
}

export const addautodnsrecord=(data)=>{
    return service({
        url: '/dns/record/auto',
        method: 'post',
        data
      })
}
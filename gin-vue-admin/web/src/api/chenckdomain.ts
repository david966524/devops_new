import service from '../utils/request'



  export const CheckDomain = (params) => {
    return service({
      url: '/checkdomain',
      method: 'get',
      params
    })
  }
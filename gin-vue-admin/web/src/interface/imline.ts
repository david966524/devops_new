// im line对象
export interface Line{
    base_url?:string
    res_url?:string
    socket_ip?:string
    socket_port?:number
    timeout:number
    ssl?:number
    remark?:string
    type?:number
}

// im 对象
export interface Im {
    ID?:string
    CreatedAt?: string
    UpdatedAt?: string
    projectName:string
    serverip:string
    groupid?:string
    jsonconfig?:string
    remark?:string
}
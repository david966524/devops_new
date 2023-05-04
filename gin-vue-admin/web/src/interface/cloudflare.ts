// 域名接口
export  interface CfDomain {
    id: string
    name: string
    created_on: string
    modified_on: string
    name_servers: string[]
    status: string
}

// apikeyManager
export interface ApiKeyManager{
    id?: number
    name: string 
    email: string
    key: string
    type: number
}
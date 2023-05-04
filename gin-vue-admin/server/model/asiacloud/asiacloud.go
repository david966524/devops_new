package asiacloud

import "gorm.io/gorm"

// 亚洲云海
type ResponseBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type VhostResult struct {
	ResponseBody
	Data struct {
		Total int `json:"total"`
		List  []struct {
			Pid    int    `json:"pid";gorm:"primary_key";"type:int(50);not null;comment:VhostPid"`
			Name   string `json:"name";gorm:"type:varchar(50);not null;comment:vhostName"`
			Status int    `json:"status";gorm:"type:int(50);not null;comment:vhost状态"`
		} `json:"list"`
	} `json:"data"`
}

type LoginResult struct {
	ResponseBody
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

type LoginPayload struct {
	Sign string `json:"sign"`
	Time int64  `json:"time"`
	Uid  string `json:"uid"`
}

type AsiaCloudAccount struct {
	gorm.Model
	Uid    string `form:"uid";gorm:"type:varchar(50);uniqueIndex;not null;comment:uid"`
	Apikey string `form:"apikey";gorm:"type:varchar(50);not null;comment:密钥"`
}

type DomainData struct {
	Total int             `json:"total"`
	List  []DomainDetails `json:"list"`
}

type DomainDetails struct {
	ID            int                 `json:"id"`
	Domain        string              `json:"domain"`
	Host          string              `json:"host"`
	Vhost         string              `json:"vhost"`
	Status        int                 `json:"status"`
	Cname         string              `json:"cname"`
	PublicSetting DomainPublicSetting `json:"publicSetting"`
}
type DomainPublicSetting struct {
	Sslcsr        string `json:"sslcsr"`
	Sslkey        string `json:"sslkey"`
	Hash          string `json:"hash"`
	Hsts          int    `json:"hsts"`
	ForceSSL      int    `json:"force_ssl"`
	MaxErrorCount string `json:"max_error_count"`
	ErrorTryTime  string `json:"error_try_time"`
	Portmap       int    `json:"portmap"`
}

type Testvhost struct {
	Name string `json:"name"`
}

type Vhost struct {
	Name   string `json:"name";gorm:"type:varchar(50);not null;comment:vhostName"`
	Pid    uint   `json:"pid";gorm:"primary_key";"type:int(50);not null;comment:VhostPid"`
	Status uint   `json:"status";gorm:"type:int(50);not null;comment:vhost状态"`
}

type DomainResult struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    DomainData `json:"data"`
}

package dns

type Dnsresq struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		RecordCount string `json:"recordCount"`
		PageSize    int    `json:"pageSize"`
		Page        int    `json:"page"`
		PageCount   int    `json:"pageCount"`
		Data        []struct {
			DomainsID   string `json:"domainsID"`
			NsGroupID   string `json:"nsGroupID"`
			GroupID     string `json:"groupID"`
			Domains     string `json:"domains"`
			State       int    `json:"state"`
			UserLock    int    `json:"userLock"`
			AdminLock   int    `json:"adminLock"`
			HealthState int    `json:"healthState"`
			ViewType    string `json:"view_type"`
		} `json:"data"`
		NextPage int `json:"nextPage"`
	} `json:"data"`
}

type DnsInfo struct {
	DomainID string `json:"domainID"`
}

type DnsRecord struct {
	DomainID string `json:"domainID"`
	Type     string `json:"type"`
	Host     string `json:"host"`
	Value    string `json:"value"`
	TTL      string `json:"TTL"`
	Remark   string `json:"remark"`
}

type AutoRecord struct {
	DomainID   string `json:"domainID"`   //dns 域名ID
	DomainName string `json:"domainName"` //dns 域名名称
	Type       string `json:"type"`       // dns 解析类型
	Host       string `json:"host"`       // dns 域名前缀
	Value      string `json:"value"`      //亚洲云海 vhost cname值
	TTL        string `json:"TTL"`
	Remark     string `json:"remark"` //dns 解析备注
	ImIp       string `json:"imIp"`   //im包网ip
}

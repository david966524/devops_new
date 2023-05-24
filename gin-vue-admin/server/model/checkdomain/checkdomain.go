package checkdomain

type CheckDomainInfo struct {
	Code     int    `json:"code"`
	Url      string `json:"url"`
	Qqmsg    string `json:"qq_msg"`
	Vxmsg    string `json:"vx_msg"`
	Cause    string `json:"cause"`
	Icp_name string `json:"icp_name"`
	Icp      string `json:"icp"`
	Tips     string `json:"tips"`
}

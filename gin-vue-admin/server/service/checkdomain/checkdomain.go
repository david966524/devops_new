package checkdomain

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/model/checkdomain"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/myhttp"
)

type CheckDomainService struct{}

func (cds *CheckDomainService) CheckDomain(url string) (*checkdomain.CheckDomainInfo, error) {
	var domaininfo checkdomain.CheckDomainInfo
	requrl := fmt.Sprintf("https://v.api.aa1.cn/api/api-lj-gf/index.php?url=%v", url)
	_, err := myhttp.HttpClinet().R().SetResult(&domaininfo).EnableTrace().Get(requrl)
	if err != nil {
		return nil, err
	}
	return &domaininfo, err
}

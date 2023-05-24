package checkdomain

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	checkdomainApi
}

var (
	checkDomainService = service.ServiceGroupApp.CheckDomainGroup.CheckDomainService
)

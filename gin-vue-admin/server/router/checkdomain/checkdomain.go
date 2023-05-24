package checkdomain

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CheckDomain struct{}

func (cd *CheckDomain) InitCheckDomainRouter(Router *gin.RouterGroup) {
	checkDomainRouter := Router.Group("/checkdomain")
	checkdomain := v1.ApiGroupApp.CheckDomainGroup.CheckDomain
	{
		checkDomainRouter.GET("", checkdomain)
	}
}

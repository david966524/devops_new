package asiacloud

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type AsiaCloudRouter struct{}

func (acr *AsiaCloudRouter) InitAsiaCloudRouter(Router *gin.RouterGroup) {
	asiacloudRouter := Router.Group("/asiacloud")
	asiacloudGetvhost := v1.ApiGroupApp.AsiacloudGroup.GetVhost
	addDomain := v1.ApiGroupApp.AsiacloudGroup.AddDomain
	deleteDomain := v1.ApiGroupApp.AsiacloudGroup.DeleteDomain
	updateDomain := v1.ApiGroupApp.AsiacloudGroup.UpdateDomain
	getDomainList := v1.ApiGroupApp.AsiacloudGroup.GetDomainListByVhost
	{
		asiacloudRouter.GET("/vhost", asiacloudGetvhost)
		asiacloudRouter.GET("/:vhost", getDomainList)
		asiacloudRouter.PUT("", updateDomain)
		asiacloudRouter.POST("", addDomain)
		asiacloudRouter.DELETE("", deleteDomain)
	}

}

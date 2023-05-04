package cloudflare

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CloudflareRouter struct{}

func (cf *CloudflareRouter) InitCloudflareRouter(Router *gin.RouterGroup) {
	cfRouter := Router.Group("/cloudflare")
	getCfList := v1.ApiGroupApp.CfApiGroup.GetAll
	{
		cfRouter.GET("/all", getCfList)
	}
}

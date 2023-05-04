package apimanager

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ApiManagerRouter struct{}

func (cfr *ApiManagerRouter) InitApiManagerRouter(Router *gin.RouterGroup) {
	apimanagerRouter := Router.Group("/apiKeyManager")
	CreateCfApi := v1.ApiGroupApp.ApiManagerGroup.CreateCfApi
	GetList := v1.ApiGroupApp.ApiManagerGroup.GetList
	DeleteCfApi := v1.ApiGroupApp.ApiManagerGroup.DeleteCfApi
	GetCfApiById := v1.ApiGroupApp.ApiManagerGroup.GetCfApiById
	UpdateCfApi := v1.ApiGroupApp.ApiManagerGroup.UpdateCfApi
	{
		apimanagerRouter.POST("/cloudflare", CreateCfApi)
		apimanagerRouter.GET("/cloudflareList", GetList)
		apimanagerRouter.DELETE("/cloudflare", DeleteCfApi)
		apimanagerRouter.GET("/cloudflare", GetCfApiById)
		apimanagerRouter.PUT("/cloudflare", UpdateCfApi)
	}

}

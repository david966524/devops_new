package packagenet

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ImRouter struct{}

func (ir *ImRouter) InitImRouter(Router *gin.RouterGroup) {
	imrouter := Router.Group("/im")
	imgetlist := v1.ApiGroupApp.PackageNetGroup.ImApi.GetImList
	imgetSingle := v1.ApiGroupApp.PackageNetGroup.ImApi.GetImById
	imUpdate := v1.ApiGroupApp.PackageNetGroup.ImApi.UpdateIm
	imDelete := v1.ApiGroupApp.PackageNetGroup.ImApi.DeleteIm
	creatIM := v1.ApiGroupApp.PackageNetGroup.ImApi.CreateIm
	getLines := v1.ApiGroupApp.PackageNetGroup.GetLines
	saveLines := v1.ApiGroupApp.PackageNetGroup.SaveLines
	{
		imrouter.GET("", imgetlist)
		imrouter.GET("/single", imgetSingle)
		imrouter.PUT("", imUpdate)
		imrouter.DELETE("", imDelete)
		imrouter.POST("", creatIM)
		imrouter.GET("/lines", getLines)
		imrouter.PUT("/lines/:id", saveLines)
	}
}

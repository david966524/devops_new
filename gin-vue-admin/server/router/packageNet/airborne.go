package packagenet

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type AirborneRouter struct{}

func (ir *AirborneRouter) InitAirborneRouter(Router *gin.RouterGroup) {
	airborneRouter := Router.Group("/airborne")
	getAirborneList := v1.ApiGroupApp.PackageNetGroup.GetAirborneList
	getAirborneById := v1.ApiGroupApp.PackageNetGroup.GetAirborneById
	deleteAirborne := v1.ApiGroupApp.PackageNetGroup.DeleteAirborne
	updateAirborne := v1.ApiGroupApp.PackageNetGroup.UpdateAirborne
	createAirborne := v1.ApiGroupApp.PackageNetGroup.CreateAirborne
	{
		airborneRouter.GET("", getAirborneList)
		airborneRouter.GET("/single", getAirborneById)
		airborneRouter.POST("", createAirborne)
		airborneRouter.PUT("", updateAirborne)
		airborneRouter.DELETE("", deleteAirborne)
	}
}

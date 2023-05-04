package jenkins

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type JenkinsRouter struct{}

func (jr *JenkinsRouter) InitJenkinsRouter(Router *gin.RouterGroup) {
	jenkinsRouter := Router.Group("/jenkins")
	getJobs := v1.ApiGroupApp.JenkinsGroup.JenkinsApi.GetJobs

	{
		jenkinsRouter.GET("/jobs", getJobs)
	}

}

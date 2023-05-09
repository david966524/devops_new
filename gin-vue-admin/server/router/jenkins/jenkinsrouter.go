package jenkins

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type JenkinsRouter struct{}

func (jr *JenkinsRouter) InitJenkinsRouter(Router *gin.RouterGroup) {
	jenkinsRouter := Router.Group("/jenkins")
	getJobs := v1.ApiGroupApp.JenkinsGroup.JenkinsApi.GetJobs
	getJobBuildHistory := v1.ApiGroupApp.JenkinsGroup.JenkinsApi.GetJobBuildHistory
	getBuildInfo := v1.ApiGroupApp.JenkinsGroup.JenkinsApi.GetBuildInfoById
	getConsoleOut := v1.ApiGroupApp.JenkinsGroup.JenkinsApi.GetBuildConsoleOut
	getBuildParameters := v1.ApiGroupApp.JenkinsGroup.JenkinsApi.GetBuildParameters
	buildJob := v1.ApiGroupApp.JenkinsGroup.JenkinsApi.BuildJob
	{
		jenkinsRouter.GET("/jobs", getJobs)
		jenkinsRouter.GET("/jobs/build", getJobBuildHistory)
		jenkinsRouter.GET("/jobs/build/info", getBuildInfo)
		jenkinsRouter.GET("/jobs/build/info/consoleout", getConsoleOut)
		jenkinsRouter.GET("/jobs/build/parameters", getBuildParameters)
		jenkinsRouter.POST("/jobs/build", buildJob)
	}

}

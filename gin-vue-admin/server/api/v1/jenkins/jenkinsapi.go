package jenkins

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/jenkins"
	"github.com/gin-gonic/gin"
)

type JenkinsApi struct{}

func (ja *JenkinsApi) GetJobs(c *gin.Context) {
	jobsinfo, err := jenkinsService.GetJobs()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(jobsinfo, c)
}

func (ja *JenkinsApi) GetJobBuildHistory(c *gin.Context) {
	var jobinfo jenkins.JobInfo
	if err := c.ShouldBindQuery(&jobinfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(jobinfo.Name)
	history, err := jenkinsService.GetJobBuildHistory(jobinfo.Name)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(history, c)
}

func (ja *JenkinsApi) GetBuildInfoById(c *gin.Context) {
	var jobinfo jenkins.JobInfo
	if err := c.ShouldBindQuery(&jobinfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	info, err := jenkinsService.GetBuildInfo(jobinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(info, c)
}
func (ja *JenkinsApi) GetBuildConsoleOut(c *gin.Context) {
	var jobinfo jenkins.JobInfo
	if err := c.ShouldBindQuery(&jobinfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	consoleout, err := jenkinsService.GetBuildConsoleOut(jobinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(consoleout, c)
}

func (ja *JenkinsApi) GetBuildParameters(c *gin.Context) {
	var jobinfo jenkins.JobInfo
	if err := c.ShouldBindQuery(&jobinfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(jobinfo.Name)
	parameters, err := jenkinsService.GetBuildParameters(jobinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(parameters, c)
}

func (ja *JenkinsApi) BuildJob(c *gin.Context) {
	var jobinfo jenkins.JobInfo
	err := c.ShouldBindJSON(&jobinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	buildid, err := jenkinsService.Buildjob(&jobinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	msg := fmt.Sprintf("构建中 %v", buildid)
	response.OkWithMessage(msg, c)
}

func (ja *JenkinsApi) BuildJobParameter(c *gin.Context) {
	var jobinfo jenkins.JobInfo
	err := c.ShouldBindJSON(&jobinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(jobinfo)
	buildid, err := jenkinsService.BuildJobParameter(&jobinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	msg := fmt.Sprintf("构建中 %v", buildid)
	response.OkWithMessage(msg, c)
}

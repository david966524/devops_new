package jenkins

import "github.com/gin-gonic/gin"

type JenkinsApi struct{}

func (ja *JenkinsApi) GetJobs(c *gin.Context) {
	jenkinsService.GetJobs()
}

package jenkins

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	JenkinsApi
}

var (
	jenkinsService = service.ServiceGroupApp.JenkinsGroup
)

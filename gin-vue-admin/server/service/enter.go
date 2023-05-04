package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/apimanager"
	"github.com/flipped-aurora/gin-vue-admin/server/service/asiacloud"
	"github.com/flipped-aurora/gin-vue-admin/server/service/cloudflare"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/jenkins"
	packagenet "github.com/flipped-aurora/gin-vue-admin/server/service/packageNet"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup    system.ServiceGroup
	ExampleServiceGroup   example.ServiceGroup
	CfManagerServiceGroup apimanager.ServiceGroup
	CfServiceGroup        cloudflare.ServiceGroup
	PackageNetGroup       packagenet.ServiceGroup
	AsiacloudGroup        asiacloud.ServiceGroup
	JenkinsGroup          jenkins.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

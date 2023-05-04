package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/apimanager"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/asiacloud"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/jenkins"

	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/cloudflare"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	packagenet "github.com/flipped-aurora/gin-vue-admin/server/api/v1/packageNet"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	CfApiGroup      cloudflare.ApiGroup
	ApiManagerGroup apimanager.ApiGroup
	PackageNetGroup packagenet.ApiGroup
	AsiacloudGroup  asiacloud.ApiGroup
	JenkinsGroup    jenkins.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

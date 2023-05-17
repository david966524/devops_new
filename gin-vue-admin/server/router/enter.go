package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/apimanager"
	"github.com/flipped-aurora/gin-vue-admin/server/router/asiacloud"
	"github.com/flipped-aurora/gin-vue-admin/server/router/cloudflare"
	"github.com/flipped-aurora/gin-vue-admin/server/router/dns"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/jenkins"
	packagenet "github.com/flipped-aurora/gin-vue-admin/server/router/packageNet"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System     system.RouterGroup
	Example    example.RouterGroup
	Cloudflare cloudflare.RouterGroup
	ApiManager apimanager.ApiManagerRouter
	PackageNet packagenet.RouterGroup
	Asiacloud  asiacloud.RouterGroup
	Jenkins    jenkins.RouterGroup
	Dns        dns.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

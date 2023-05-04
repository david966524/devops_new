package asiacloud

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	Asiacloud
}

var (
	asiacloudService = service.ServiceGroupApp.AsiacloudGroup
)

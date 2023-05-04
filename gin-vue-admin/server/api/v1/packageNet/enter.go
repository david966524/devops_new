package packagenet

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

type ApiGroup struct {
	AirborneApi
	ImApi
}

var (
	imService       = service.ServiceGroupApp.PackageNetGroup
	airborneService = service.ServiceGroupApp.PackageNetGroup
)

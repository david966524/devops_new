package dns

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DnsApi
}

var (
	dnsService = service.ServiceGroupApp.DnsGroup.DnsService
)

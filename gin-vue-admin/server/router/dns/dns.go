package dns

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type DnsRouter struct{}

func (dns *DnsRouter) InitDnsRouter(Router *gin.RouterGroup) {
	dnsRouter := Router.Group("/dns")
	getDomainList := v1.ApiGroupApp.DnsGroup.DnsApi.GetDomainList
	getDomainRecord := v1.ApiGroupApp.DnsGroup.DnsApi.GetDomainRecordById
	addDomainRecord := v1.ApiGroupApp.DnsGroup.DnsApi.AddDomainRecordById
	autoaddIm := v1.ApiGroupApp.DnsGroup.DnsApi.AutoAddIm
	{
		dnsRouter.GET("/all", getDomainList)
		dnsRouter.GET("/record", getDomainRecord)
		dnsRouter.POST("/record", addDomainRecord)
		dnsRouter.POST("/record/auto", autoaddIm)
	}
}

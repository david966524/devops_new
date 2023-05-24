package dns

import (
	"fmt"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dns"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

type DnsApi struct{}

func (da *DnsApi) GetDomainList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	dnsresq := dnsService.GetDomainList(pageInfo)
	totel, _ := strconv.ParseInt(dnsresq.Data.RecordCount, 10, 64)
	response.OkWithDetailed(response.PageResult{
		List:     dnsresq.Data.Data,
		Total:    totel,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (da *DnsApi) GetDomainRecordById(c *gin.Context) {
	var dnsinfo dns.DnsInfo
	err := c.ShouldBindQuery(&dnsinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	recordvalue := dnsService.GetDomainRecordById(dnsinfo)
	response.OkWithData(recordvalue.Value(), c)
}

func (da *DnsApi) AddDomainRecordById(c *gin.Context) {
	var dnsRecord dns.DnsRecord
	err := c.ShouldBindJSON(&dnsRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	resp := dnsService.AddDomainRecordById(dnsRecord)
	respvalue := gjson.Get(resp, "data").Value()
	response.OkWithData(respvalue, c)
}

func (da *DnsApi) AutoAddIm(c *gin.Context) {
	var autoRecordim dns.AutoRecord
	err := c.ShouldBindJSON(&autoRecordim)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(autoRecordim)
	err1 := dnsService.AutoAddIm(autoRecordim)
	if err1 != nil {
		response.FailWithMessage(err1.Error(), c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

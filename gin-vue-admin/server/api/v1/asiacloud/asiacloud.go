package asiacloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/asiacloud"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type Asiacloud struct{}

func (ac *Asiacloud) GetVhost(c *gin.Context) {
	vhostResult, err := asiacloudService.GetVhost()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(vhostResult.Data.List, c)
}

func (ac *Asiacloud) GetDomainListByVhost(c *gin.Context) {
	vhost := c.Param("vhost")
	data, err := asiacloudService.GetDomainListByVhost(vhost)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(data.Data.List, c)
}

func (ac *Asiacloud) AddDomain(c *gin.Context) {
	var ad asiacloud.DomainDetails
	if err := c.ShouldBindJSON(&ad); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := asiacloudService.AddDomain(&ad); err != nil {
		response.FailWithMessage("添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

func (ac *Asiacloud) DeleteDomain(c *gin.Context) {
	var ad asiacloud.DomainDetails
	err := c.ShouldBindJSON(&ad)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := asiacloudService.DeleteDomain(&ad); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (ac *Asiacloud) UpdateDomain(c *gin.Context) {
	var ad asiacloud.DomainDetails
	err := c.ShouldBindJSON(&ad)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := asiacloudService.UpdateDomain(&ad); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

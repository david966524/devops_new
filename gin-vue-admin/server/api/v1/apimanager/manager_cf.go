package apimanager

import (
	"fmt"

	cloudflaremodel "github.com/flipped-aurora/gin-vue-admin/server/model/cloudflare"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
)

type ApiManagerCf struct{}

func (amf *ApiManagerCf) CreateCfApi(c *gin.Context) {
	fmt.Println("api run ..........")
	var account cloudflaremodel.CloudFlareAccount
	if err := c.ShouldBind(&account); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := service.ServiceGroupApp.CfManagerServiceGroup.CreateApiAccount(&account)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

func (amf *ApiManagerCf) GetList(c *gin.Context) {
	fmt.Println("api get all run")
	list, err := service.ServiceGroupApp.CfManagerServiceGroup.GetCfApiList()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(list, c)
}

func (amf *ApiManagerCf) DeleteCfApi(c *gin.Context) {
	var account cloudflaremodel.CloudFlareAccount
	if err := c.BindJSON(&account); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Printf("删除cfapi的id为%v\n", account.ID)
	err := service.ServiceGroupApp.CfManagerServiceGroup.DeleteCfApi(&account)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (amf *ApiManagerCf) GetCfApiById(c *gin.Context) {
	var account cloudflaremodel.CloudFlareAccount
	if err := c.ShouldBindQuery(&account); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(account.ID)
	result, err := service.ServiceGroupApp.CfManagerServiceGroup.GetCfApiById(account.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(result, c)
}

func (amf *ApiManagerCf) UpdateCfApi(c *gin.Context) {
	var account cloudflaremodel.CloudFlareAccount
	if err := c.BindJSON(&account); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := service.ServiceGroupApp.CfManagerServiceGroup.UpdateCfApi(&account)
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)

}

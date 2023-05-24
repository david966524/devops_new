package cloudflare

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
)

type cloudflareApi struct{}

func (cf *cloudflareApi) GetAll(c *gin.Context) {
	r, err := service.ServiceGroupApp.CfServiceGroup.GetList()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(r.Result, c)
}

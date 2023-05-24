package checkdomain

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

type checkdomainApi struct{}
type reqData struct {
	Url string `json:"url"`
}

func (cda *checkdomainApi) CheckDomain(c *gin.Context) {
	var reqdata reqData
	// err := c.ShouldBindJSON(&reqdata)
	err := c.ShouldBindQuery(&reqdata)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(reqdata, utils.UrlVerify)
	if err != nil {
		response.FailWithMessage("不能为空", c)
		return
	}

	result, err := checkDomainService.CheckDomain(reqdata.Url)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(result, c)
}

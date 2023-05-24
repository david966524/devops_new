package packagenet

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/packageNet/im"
	"github.com/gin-gonic/gin"
)

type ImApi struct{}

func (imapi *ImApi) GetImList(c *gin.Context) {
	ims, err := imService.GetImList()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(ims, c)
}
func (imapi *ImApi) GetImById(c *gin.Context) {
	var im im.Im
	err := c.ShouldBindQuery(&im)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	i, err := imService.GetImById(im)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(i, c)
}

func (imapi *ImApi) UpdateIm(c *gin.Context) {
	var im im.Im
	err := c.ShouldBindJSON(&im)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := imService.UpdataIm(&im); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}
func (imapi *ImApi) DeleteIm(c *gin.Context) {
	var im im.Im
	if err := c.ShouldBindJSON(&im); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := imService.DeleteIm(&im); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (imapi *ImApi) CreateIm(c *gin.Context) {
	var im im.Im
	if err := c.ShouldBindJSON(&im); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := imService.CreateIm(&im); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

func (imapi *ImApi) GetLines(c *gin.Context) {
	var im im.Im

	if err := c.ShouldBindQuery(&im); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	lines, err := imService.GetLines(int(im.ID))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(lines, c)
}

func (imapi *ImApi) SaveLines(c *gin.Context) {
	var lines []im.Line
	id := c.Param("id")
	if err := c.ShouldBindJSON(&lines); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := imService.SaveLines(id, lines); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

package packagenet

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/packageNet/airborne"
	"github.com/gin-gonic/gin"
)

type AirborneApi struct{}

func (aba *AirborneApi) GetAirborneList(c *gin.Context) {
	airbornes, err := airborneService.GetAirborneList()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(airbornes, c)
}

func (aba *AirborneApi) GetAirborneById(c *gin.Context) {
	var ab airborne.AirBorne
	if err := c.ShouldBindQuery(&ab); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	result, err := airborneService.GetAirborneById(&ab)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(result, c)
}

func (aba *AirborneApi) UpdateAirborne(c *gin.Context) {
	var ab airborne.AirBorne
	if err := c.ShouldBindJSON(&ab); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := airborneService.UpdateAirborne(&ab); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (aba *AirborneApi) DeleteAirborne(c *gin.Context) {
	var ab airborne.AirBorne
	if err := c.ShouldBindJSON(&ab); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := airborneService.DeleteAirborne(&ab); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (aba *AirborneApi) CreateAirborne(c *gin.Context) {
	var ab airborne.AirBorne
	if err := c.ShouldBindJSON(&ab); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := airborneService.CreateAirborne(&ab); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

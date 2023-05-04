package packagenet

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/packageNet/airborne"
)

type AirborneService struct{}

func (as *AirborneService) GetAirborneList() ([]airborne.AirBorne, error) {
	var airbornes []airborne.AirBorne
	result := global.GVA_DB.Find(&airbornes)
	if result.Error != nil {
		return nil, result.Error
	}
	return airbornes, nil
}

func (as *AirborneService) CreateAirborne(ab *airborne.AirBorne) error {

	result := global.GVA_DB.Create(&ab)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (as *AirborneService) UpdateAirborne(ab *airborne.AirBorne) error {

	err := global.GVA_DB.Save(ab)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
func (as *AirborneService) DeleteAirborne(ab *airborne.AirBorne) error {
	err := global.GVA_DB.Delete(ab).Error
	if err != nil {
		return err
	}
	return nil
}

func (as *AirborneService) GetAirborneById(ab *airborne.AirBorne) (airborne.AirBorne, error) {
	var airborne airborne.AirBorne
	result := global.GVA_DB.Where("id=?", ab.ID).Find(&airborne)
	if result.Error != nil {
		return airborne, result.Error
	}
	return airborne, nil
}

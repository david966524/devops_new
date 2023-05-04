package apimanager

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	cloudflaremodel "github.com/flipped-aurora/gin-vue-admin/server/model/cloudflare"
)

type CloudFlareManager struct{}

func (cfm *CloudFlareManager) CreateApiAccount(account *cloudflaremodel.CloudFlareAccount) (err error) {
	fmt.Println(account)
	err = global.GVA_DB.Create(&account).Error
	return err
}

func (cfm *CloudFlareManager) GetCfApiList() (List []cloudflaremodel.CloudFlareAccount, err error) {
	var apiAccountList []cloudflaremodel.CloudFlareAccount
	result := global.GVA_DB.Find(&apiAccountList)
	if result.Error != nil {
		return nil, result.Error
	}
	return apiAccountList, nil
}

func (cfm *CloudFlareManager) DeleteCfApi(account *cloudflaremodel.CloudFlareAccount) error {
	err := global.GVA_DB.Delete(account).Error
	return err
}

func (cfm *CloudFlareManager) GetCfApiById(id uint) (account cloudflaremodel.CloudFlareAccount, err error) {
	result := global.GVA_DB.Where("id = ? ", id).Find(&account)
	if result.Error != nil {
		return account, result.Error
	}
	return account, result.Error
}

func (cfm *CloudFlareManager) UpdateCfApi(account *cloudflaremodel.CloudFlareAccount) (err error) {
	err = global.GVA_DB.Save(account).Error
	return err
}

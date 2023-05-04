package cloudflare

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	cloudflaremodel "github.com/flipped-aurora/gin-vue-admin/server/model/cloudflare"
)

func CfClient() (*cloudflare.API, error) {
	var account cloudflaremodel.CloudFlareAccount
	result := global.GVA_DB.Where("type=?", 1).Find(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	client, err := cloudflare.New(account.Key, account.Email)
	if err != nil {
		return nil, err
	}
	return client, nil
}

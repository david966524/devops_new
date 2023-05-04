package cosutils

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tencen"
	"github.com/tencentyun/cos-go-sdk-v5"
)

// 腾讯云 cos client
func CosClient() *cos.Client {
	var tencentAccess tencen.TencentAccess
	result := global.GVA_DB.Find(&tencentAccess)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}

	u, _ := url.Parse("https://newim-txy4-123qwer-1312026995.cos.ap-hongkong.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  tencentAccess.AccessKeyId,
			SecretKey: tencentAccess.AccessKeySecret,
		},
	})
	return client
}

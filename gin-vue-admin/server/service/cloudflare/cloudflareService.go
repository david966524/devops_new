package cloudflare

import (
	"context"
	"fmt"

	cloudflareapi "github.com/cloudflare/cloudflare-go"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/cloudflare"
)

type CloudflareService struct{}

func (cfs *CloudflareService) GetList() (*cloudflareapi.ZonesResponse, error) {
	client, err := cloudflare.CfClient()
	if err != nil {
		return nil, err
	}
	// pageOpts := cloudflareapi.PaginationOptions{
	// 	Page:    1,
	// 	PerPage: 50,
	// }
	// r, err := client.ListZonesContext(context.Background(), cloudflareapi.WithPagination(pageOpts))
	r, err := client.ListZonesContext(context.Background())
	if err != nil {
		fmt.Println("获取区域列表失败：", err)
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(r.ResultInfo)
	fmt.Println(len(r.Result))

	return &r, nil
}

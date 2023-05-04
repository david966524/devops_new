package packagenet

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/packageNet/im"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/cosutils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/myhttp"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type ImService struct{}

func (is *ImService) GetImList() ([]im.Im, error) {
	var ims []im.Im
	result := global.GVA_DB.Find(&ims)
	if result.Error != nil {
		return nil, result.Error
	}
	return ims, nil
}

func (is *ImService) GetImById(i im.Im) (im.Im, error) {
	var im im.Im
	result := global.GVA_DB.Where("id = ?", i.ID).Find(&im)
	if result.Error != nil {
		return im, result.Error
	}
	return im, nil
}
func (is *ImService) UpdataIm(im *im.Im) error {
	err := global.GVA_DB.Save(im).Error
	return err
}
func (is *ImService) DeleteIm(im *im.Im) error {

	result := global.GVA_DB.Delete(im)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (is *ImService) CreateIm(im *im.Im) error {

	result := global.GVA_DB.Create(im)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (is *ImService) GetLines(id int) ([]im.Line, error) {
	var im im.Im
	result := global.GVA_DB.Where("id=?", id).Find(&im)
	if result.Error != nil {
		return nil, result.Error
	}
	lineLinst, err := reqLines(im.Jsonconfig)
	return lineLinst, err
}

func (is *ImService) SaveLines(imid string, lines []im.Line) error {
	filename := getfilename(imid)
	err := saveTocos(lines, filename)
	return err
}

// 通过 im 的url 字段获取im线路表
func reqLines(url string) ([]im.Line, error) {
	if url == "" {
		return nil, errors.New("json 为空")
	}
	client := myhttp.HttpClinet()
	var datas im.Datas
	resp, err := client.R().SetResult(&datas).EnableTrace().Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp)
	//b1, _ := json.MarshalIndent(datas.Data.Lines, "", "    ")
	//fmt.Println(string(b1))
	return datas.Data.Lines, nil
}

// 向 腾讯cos 保存 修改好的 im 线路
func saveTocos(lines []im.Line, filename string) error {
	data := &im.Datas{
		Data: im.Data{
			Lines: lines,
		},
		Ok: true,
	}
	jsonbytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	// log.Println(string(jsonbytes))
	// log.Println(filename)
	client := cosutils.CosClient()
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: "application/json",
		},
		ACLHeaderOptions: &cos.ACLHeaderOptions{
			// 如果不是必要操作，建议上传文件时不要给单个文件设置权限，避免达到限制。若不设置默认继承桶的权限。
		},
	}
	objKey := fmt.Sprintf("configs/%v", filename)
	resp, err := client.Object.Put(context.Background(), objKey, bytes.NewReader(jsonbytes), opt)
	if err != nil {
		return err
	}
	fmt.Println("Object updated successfully.")
	if resp.StatusCode != 200 {
		return errors.New("cos 响应非200")
	}
	return nil
}

// 根据im id查询 json name
func getfilename(imid string) string {
	var im im.Im
	db := global.GVA_DB.Where("id=?", imid).Find(&im)
	if db.Error != nil {
		log.Println(db.Error.Error())
	}
	s := strings.Split(im.Jsonconfig, "/")
	filename := s[len(s)-1]
	return filename
}

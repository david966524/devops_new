package asiacloud

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/asiacloud"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/myhttp"
)

type AsiacloudService struct{}

func (as *AsiacloudService) GetVhost() (asiacloud.VhostResult, error) {

	asiaCloudToken := getAsiaCloudToken()
	var vhostResult asiacloud.VhostResult
	vhostUrl := "https://cdnportal.myasiacloud.com/api/proxy/vhost"
	resp, err := myhttp.HttpClinet().R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+asiaCloudToken).
		Get(vhostUrl)
	if err != nil {
		return vhostResult, err
	}
	json.Unmarshal(resp.Body(), &vhostResult)
	log.Println(vhostResult.Data.List)
	return vhostResult, nil
}

func (as *AsiacloudService) GetDomainListByVhost(vhost string) (*asiacloud.DomainResult, error) {
	asiaCloudToken := getAsiaCloudToken()
	var domainResult asiacloud.DomainResult
	getdomaintUrl := "https://cdnportal.myasiacloud.com/api/site/" + vhost + "/domain/page"
	resp, err := myhttp.HttpClinet().R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+asiaCloudToken).
		Get(getdomaintUrl)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(resp.Body(), &domainResult)
	return &domainResult, nil
}

func (as *AsiacloudService) AddDomain(ad *asiacloud.DomainDetails) error {
	asiaCloudToken := getAsiaCloudToken()
	addDomainUrl := "https://cdnportal.myasiacloud.com/api/site/" + ad.Vhost + "/domain"
	_, err := myhttp.HttpClinet().R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+asiaCloudToken).
		SetBody(ad).
		Post(addDomainUrl)
	if err != nil {
		return err
	}
	// log.Println(resp.Body())
	return nil
}
func (as *AsiacloudService) DeleteDomain(ad *asiacloud.DomainDetails) error {
	deletedomaintUrl := fmt.Sprintf("https://cdnportal.myasiacloud.com/api/site/%v/domain/%v", ad.Vhost, ad.ID)
	asiaCloudToken := getAsiaCloudToken()
	_, err := myhttp.HttpClinet().R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+asiaCloudToken).
		Delete(deletedomaintUrl)
	if err != nil {
		return err
	}
	// log.Println(resq.Body())
	return nil
}

func (as *AsiacloudService) UpdateDomain(ad *asiacloud.DomainDetails) error {
	updatedomaintUrl := fmt.Sprintf("https://cdnportal.myasiacloud.com/api/site/%v/domain/%v", ad.Vhost, ad.ID)
	asiaCloudToken := getAsiaCloudToken()
	_, err := myhttp.HttpClinet().R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+asiaCloudToken).
		SetBody(ad).
		Put(updatedomaintUrl)
	if err != nil {
		return err
	}
	return nil
}

// 获取 亚洲云海token
func getAsiaCloudToken() string {
	var accunot asiacloud.AsiaCloudAccount
	result := global.GVA_DB.Where("id=?", 1).Find(&accunot)
	if result.Error != nil {
		log.Println(result.Error.Error())
	}
	apiskeyMd5 := MD5(accunot.Uid + accunot.Apikey)
	timestamp := time.Now().Unix()
	fullapiMd5 := MD5(apiskeyMd5 + strconv.Itoa(int(timestamp)))
	fmt.Println("当前时间戳  ", timestamp)
	// fmt.Println(fullapiMd5)
	loginUrl := "https://cdnportal.myasiacloud.com/api/user/login/token"
	var myresule asiacloud.LoginResult
	mypayload := asiacloud.LoginPayload{
		Sign: fullapiMd5,
		Time: timestamp,
		Uid:  accunot.Uid,
	}
	_, err := myhttp.HttpClinet().R().
		SetHeader("Content-Type", "application/json").
		SetResult(&myresule).
		SetBody(&mypayload).
		Post(loginUrl)
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Println(string(resp.Body()))
	asiaCloudToken := myresule.Data.Token
	return asiaCloudToken
}
func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

package dns

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/asiacloud"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dns"
	"github.com/flipped-aurora/gin-vue-admin/server/model/packageNet/im"
	"github.com/flipped-aurora/gin-vue-admin/server/service/asiacloud"
	packagenet "github.com/flipped-aurora/gin-vue-admin/server/service/packageNet"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/myhttp"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/telegram"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

type DnsService struct{}

func (ds *DnsService) GetDomainList(pageinfo request.PageInfo) *dns.Dnsresq {
	params := url.Values{}
	params.Set("groupID", "")
	params.Set("page", strconv.Itoa(pageinfo.Page))
	params.Set("pageSize", strconv.Itoa(pageinfo.PageSize))
	fill(params)
	c := resty.New()
	rsp, err := c.R().SetQueryParamsFromValues(params).Get("https://www.dns.com/api/domain/list")
	if err != nil {
		return nil
	}
	var dnsresq dns.Dnsresq
	json.Unmarshal(rsp.Body(), &dnsresq)
	return &dnsresq
}

func (ds *DnsService) GetDomainRecordById(dnsinfo dns.DnsInfo) *gjson.Result {
	params := url.Values{}
	params.Set("domainID", dnsinfo.DomainID)
	params.Set("pageSize", "20")
	fill(params)
	c := resty.New()
	rsp, err := c.R().SetQueryParamsFromValues(params).Get("https://www.dns.com/api/record/list")
	if err != nil {
		return nil
	}
	recordvalue := gjson.Get(rsp.String(), "data.data")
	return &recordvalue
}

func (ds *DnsService) AddDomainRecordById(dnsRecord dns.DnsRecord) string {
	params := url.Values{}
	params.Set("domainID", dnsRecord.DomainID)
	params.Set("type", dnsRecord.Type)
	params.Set("host", dnsRecord.Host)
	params.Set("value", dnsRecord.Value)
	params.Set("TTL", dnsRecord.TTL)
	params.Set("remark", dnsRecord.Remark)
	fill(params)
	c := resty.New()
	rsp, err := c.R().SetQueryParamsFromValues(params).Post("https://www.dns.com/api/record/create")
	if err != nil {
		return ""
	}
	return rsp.String()
}

func (ds *DnsService) AutoAddIm(obj dns.AutoRecord) error {

	fullDomain := fmt.Sprintf("%v.%v", obj.Host, obj.DomainName)
	fullhost := fmt.Sprintf("%v:%v", obj.ImIp, "9292")
	s := strings.Split(obj.Value, ".")
	fmt.Println(s[0])
	imobj := packagenet.GetIMObj(obj.ImIp)
	lineList, _ := packagenet.ReqLines(imobj.Jsonconfig)
	fmt.Println(imobj)
	fmt.Println(lineList)
	msg := fmt.Sprintf(`
		自动解析记录 %v     %v    %v   %v 
	`, fullDomain, fullhost, obj.Value, imobj.ProjectName)
	telegram.SendTgMsg(msg)
	err := addRecord(obj)
	if err != nil {
		return err
	}
	err1 := addasiacloud(fullDomain, fullhost, s[0])
	if err1 != nil {
		return err1
	}
	err2 := changeJson(lineList, fullDomain, imobj)
	if err2 != nil {
		return err2
	}
	return nil
}

func fill(params url.Values) {
	params.Set("apiKey", global.GVA_CONFIG.Dns.ApiKey)
	params.Set("timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	m := md5.New()
	rawquery, _ := url.QueryUnescape(params.Encode())
	m.Write([]byte(rawquery + global.GVA_CONFIG.Dns.ApiSecret))
	params.Set("hash", hex.EncodeToString(m.Sum(nil)))
}

// 向dns 添加解析
func addRecord(obj dns.AutoRecord) error {
	params := url.Values{}
	params.Set("domainID", obj.DomainID)
	params.Set("type", obj.Type)
	params.Set("host", obj.Host)
	params.Set("value", obj.Value)
	params.Set("TTL", obj.TTL)
	params.Set("remark", obj.Remark)
	fill(params)
	c := resty.New()
	rsp, err := c.R().SetQueryParamsFromValues(params).Post("https://www.dns.com/api/record/create")
	if err != nil {
		return err
	}
	telegram.SendTgMsg(rsp.String())
	return nil
}

// 向方能添加解析
func addasiacloud(domain string, host string, vhost string) error {
	asiaCloudToken := asiacloud.GetAsiaCloudToken()
	addDomainUrl := "https://cdnportal.myasiacloud.com/api/site/" + vhost + "/domain"
	resp, err := myhttp.HttpClinet().R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+asiaCloudToken).
		SetBody(model.DomainDetails{
			Domain: domain,
			Host:   host,
		}).
		Post(addDomainUrl)

	if err != nil {
		return err
	}
	telegram.SendTgMsg(resp.String())
	return nil
}

// 修改json
func changeJson(lines []im.Line, domain string, imobj *im.Im) error {
	var newLine im.Line
	for _, v := range lines {
		if v.Type == 0 {
			newLine = v
			break
		}
	}
	newdomain := fmt.Sprintf("http://%v", domain)
	newLine.BaseUrl = newdomain
	lines = append(lines, newLine)
	json_file_name := packagenet.Getfilename(strconv.FormatUint(uint64(imobj.ID), 10))
	err := packagenet.SaveTocos(lines, json_file_name)
	return err
}

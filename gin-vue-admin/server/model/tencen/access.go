package tencen

import "gorm.io/gorm"

type TencentAccess struct {
	gorm.Model
	AccessKeyId     string `json:"accessKeyId";gorm:"type:varchar(50);not null;comment:腾讯云apikey"`
	AccessKeySecret string `json:"accessKeySecret";gorm:"type:varchar(50);not null;comment:腾讯云api密钥"`
	Type            int    `json:"type";gorm:"type:varchar(50);not null;comment:开启状态 1 开启 0关闭"`
}

package cloudflaremodel

import "gorm.io/gorm"

type CloudFlareAccount struct {
	gorm.Model
	Name  string `gorm:"type:varchar(50);not null;comment:名称"`
	Email string `gorm:"type:varchar(50);not null;comment:邮箱"`
	Key   string `gorm:"type:varchar(50);not null;comment:密钥"`
	Type  int    `gorm:"type:bigint(20);not null;comment:状态"`
}

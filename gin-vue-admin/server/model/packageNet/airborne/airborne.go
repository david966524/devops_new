package airborne

import "gorm.io/gorm"

type AirBorne struct {
	gorm.Model
	ProjectName string `json:"projectName";form:"projectName";gorm:"type:varchar(50);not null;comment:项目名"`
	ServerIp    string `json:"serverip";form:"serverip";gorm:"type:varchar(50);not null;comment:服务器ip"`
	Groupid     string `json:"groupid";form:"groupid";gorm:"type:varchar(50);comment:方能对应组id"`
	Remark      string `json:"remark";form:"remark";gorm:"type:varchar(50);comment:备注"`
}

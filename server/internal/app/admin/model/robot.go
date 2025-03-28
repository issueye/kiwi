package model

import (
	"kiwi/internal/common/model"
	"time"
)

type Robot struct {
	model.BaseModel
	Name      string    `gorm:"size:100;not null;comment:机器人名称"`
	ApiKey    string    `gorm:"size:255;uniqueIndex;not null;comment:API访问密钥"`
	ExpiresAt time.Time `gorm:"comment:密钥过期时间"`
	IsActive  bool      `gorm:"default:true;comment:是否激活"`
	CreatedBy uint      `gorm:"comment:创建者ID"`
	Remark    string    `gorm:"size:500;comment:备注信息"`
}

func (r Robot) TableName() string { return "sys_robot" }
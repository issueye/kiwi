package model

import (
	"kiwi/internal/common/model"
)

type RobotInfo struct {
	model.BaseModel
	RobotBase
}

type EnumRobotType string

const (
	QQRobot       EnumRobotType = "qq"
	DingtalkRobot EnumRobotType = "dingtalk"
)

type RobotBase struct {
	Name       string        `gorm:"column:name;size:200;not null;comment:名称;" json:"name"`                       // 名称
	RobotType  EnumRobotType `gorm:"column:robot_type;size:200;not null;comment:机器人类型;" json:"robot_type"`        // 机器人类型
	WebhookURL string        `gorm:"column:webhook_url;size:-1;not null;comment:webhook url;" json:"webhook_url"` // webhook url
	Token      string        `gorm:"column:token;size:-1;not null;comment:token;" json:"token"`                   // token
	Secret     string        `gorm:"column:secret;size:-1;not null;comment:secret;" json:"secret"`                // secret
	Script     string        `gorm:"column:script;size:200;not null;comment:脚本;" json:"script"`                   // script
	State      bool          `gorm:"column:state;not null;comment:状态;" json:"state"`                              // 状态
}

func (p RobotInfo) TableName() string { return "biz_robot_info" }

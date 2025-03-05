package requests

import (
	"encoding/json"
	"errors"
	"kiwi/internal/app/project/model"
	commonModel "kiwi/internal/common/model"
)

type CreateRobot struct {
	model.RobotInfo
}

func (req *CreateRobot) Validate() error {
	if req.Name == "" {
		return errors.New("机器人名称不能为空")
	}
	if req.RobotType != "qq" && req.RobotType != "dingtalk" {
		return errors.New("无效的机器人类型")
	}
	return nil
}

func NewCreateRobot() *CreateRobot {
	return &CreateRobot{
		RobotInfo: model.RobotInfo{},
	}
}

type UpdateRobot struct {
	model.RobotInfo
}

func (req *UpdateRobot) Validate() error {
	if req.Name == "" {
		return errors.New("机器人名称不能为空")
	}

	if req.RobotType != "qq" && req.RobotType != "dingtalk" {
		return errors.New("无效的机器人类型")
	}
	return nil
}

func NewUpdateRobot() *UpdateRobot {
	return &UpdateRobot{
		RobotInfo: model.RobotInfo{},
	}
}

func (req *CreateRobot) ToJson() string {
	data, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(data)
}

type QueryRobot struct {
	KeyWords string `json:"keywords" form:"keywords"` // 关键词
}

func NewQueryRobot() *commonModel.PageQuery[*QueryRobot] {
	return commonModel.NewPageQuery(0, 0, &QueryRobot{})
}

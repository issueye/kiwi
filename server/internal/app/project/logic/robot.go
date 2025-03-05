package logic

import (
	"kiwi/internal/app/project/model"
	"kiwi/internal/app/project/requests"
	"kiwi/internal/app/project/service"
	commonModel "kiwi/internal/common/model"
)

func CreateRobot(req *requests.CreateRobot) error {
	srv := service.NewRobot()
	return srv.Create(&req.RobotInfo)
}

func UpdateRobot(req *requests.UpdateRobot) error {
	dataMap := make(map[string]any)
	dataMap["name"] = req.Name
	dataMap["webhook_url"] = req.WebhookURL
	dataMap["token"] = req.Token
	dataMap["secret"] = req.Secret
	dataMap["script"] = req.Script
	dataMap["robot_type"] = req.RobotType
	return service.NewRobot().UpdateByMap(req.ID, dataMap)
}

func UpdateRobotStatus(id uint, state int) error {
	return service.NewRobot().UpdateRobotState(id, state)
}

func DeleteRobot(id uint) error {
	return service.NewRobot().Delete(id)
}

func RobotList(condition *commonModel.PageQuery[*requests.QueryRobot]) (*commonModel.ResPage[model.RobotInfo], error) {
	return service.NewRobot().ListRobot(condition)
}

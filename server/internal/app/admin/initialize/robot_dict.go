package initialize

import (
	"kiwi/internal/app/admin/model"
	"kiwi/internal/app/admin/service"
)

func InitRobotTypeDict() {
	robotDict := &model.DictsInfo{
		DictsBase: model.DictsBase{
			Code:        "10000",
			Name:        "机器人类型",
			ContentType: model.ContentTypeText,
			Details: []model.DictDetail{
				{DictDetailBase: model.DictDetailBase{Key: "QQ", Value: "QQ"}},
				{DictDetailBase: model.DictDetailBase{Key: "DTalk", Value: "DTalk"}},
			},
		},
	}

	RobotTypeIsNotExistAdd(robotDict)
}

func RobotTypeIsNotExistAdd(data *model.DictsInfo) {
	service.NewDicts().NotExistCreate(data)
}

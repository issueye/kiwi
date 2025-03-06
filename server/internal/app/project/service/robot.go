package service

import (
	"kiwi/internal/app/project/model"
	"kiwi/internal/app/project/requests"
	commonModel "kiwi/internal/common/model"
	"kiwi/internal/common/service"

	"gorm.io/gorm"
)

type Robot struct {
	service.BaseService[model.RobotInfo]
}

func NewRobot(args ...any) *Robot {
	srv := &Robot{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// ListRobot
// 根据条件查询列表
func (r *Robot) ListRobot(condition *commonModel.PageQuery[*requests.QueryRobot]) (*commonModel.ResPage[model.RobotInfo], error) {
	return service.GetList[model.RobotInfo](condition, func(qu *requests.QueryRobot, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or remark like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		return d
	})
}

// UpdateRobotStatus 更新分支状态
func (r *Robot) UpdateRobotState(id uint, state int) error {
	return r.GetDB().Model(&model.RobotInfo{}).
		Where("id = ?", id).
		Update("state", state).Error
}

package service

import (
	"kiwi/internal/app/project/model"
	"kiwi/internal/app/project/requests"
	commonModel "kiwi/internal/common/model"
	"kiwi/internal/common/service"

	"gorm.io/gorm"
)

type Project struct {
	service.BaseService[model.ProjectInfo]
}

func NewProject(args ...any) *Project {
	srv := &Project{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// ListProject
// 根据条件查询列表
func (r *Project) ListProject(condition *commonModel.PageQuery[*requests.QueryProject]) (*commonModel.ResPage[model.ProjectInfo], error) {
	return service.GetList[model.ProjectInfo](condition, func(qu *requests.QueryProject, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or remark like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		// 加载分支
		d = d.Preload("Branchs")

		// 加载机器人
		d = d.Preload("ProjectRobots")

		return d
	})
}

func (r *Project) SaveProjectRobots(projectId uint, robots []uint) error {
	// 先删除
	err := r.GetDB().Where("project_id =?", projectId).Delete(&model.ProjectRobotInfo{}).Error
	if err != nil {
		return err
	}

	// 再添加
	for _, robotId := range robots {
		info := &model.ProjectRobotInfo{}
		info.ProjectId = projectId
		info.RobotId = robotId

		err = r.GetDB().Create(info).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Project) DeleteProjectRobots(projectId uint) error {
	return r.GetDB().Where("project_id = ?", projectId).Delete(&model.ProjectRobotInfo{}).Error
}

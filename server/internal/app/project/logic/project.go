package logic

import (
	"kiwi/internal/app/project/model"
	"kiwi/internal/app/project/requests"
	"kiwi/internal/app/project/service"
	commonModel "kiwi/internal/common/model"
)

func CreateProject(req *requests.CreateProject) error {
	srv := service.NewProject()
	srv.Begin()
	var err error
	defer func() {
		if err != nil {
			srv.Rollback()
			return
		}

		srv.Commit()
	}()

	err = srv.Create(&req.ProjectInfo)
	if err != nil {
		return err
	}

	err = srv.SaveProjectRobots(req.ID, req.Robots)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProject(req *requests.UpdateProject) error {
	srv := service.NewProject()
	srv.Begin()
	var err error
	defer func() {
		if err != nil {
			srv.Rollback()
			return
		}

		srv.Commit()
	}()

	srv.Update(req.ID, &req.ProjectInfo)
	if err != nil {
		return err
	}

	err = srv.SaveProjectRobots(req.ID, req.Robots)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProject(id uint) error {
	// 删除字典，并且删除对应明细数据
	srv := service.NewProject()

	info, err := srv.GetById(id)
	if err != nil {
		return err
	}

	srv.Begin()

	defer func() {
		if err != nil {
			srv.Rollback()
			return
		}

		srv.Commit()
	}()

	err = srv.Delete(id)
	if err != nil {
		return err
	}

	branchSrv := service.NewBranch(srv.GetDB(), true)
	err = branchSrv.DeleteByFields(map[string]any{"project_id": info.ID})
	if err != nil {
		return err
	}

	err = srv.DeleteProjectRobots(info.ID)
	if err != nil {
		return err
	}

	return nil
}

func ProjectList(condition *commonModel.PageQuery[*requests.QueryProject]) (*commonModel.ResPage[model.ProjectInfo], error) {
	return service.NewProject().ListProject(condition)
}

func GetProject(id uint) (*model.ProjectInfo, error) {
	return service.NewProject().GetById(id)
}

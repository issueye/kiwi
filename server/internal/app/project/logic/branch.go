package logic

import (
	"kiwi/internal/app/project/model"
	"kiwi/internal/app/project/requests"
	"kiwi/internal/app/project/service"
	commonModel "kiwi/internal/common/model"
)

func CreateBranch(req *requests.CreateBranch) error {
	srv := service.NewBranch()
	return srv.Create(&req.BranchInfo)
}

func UpdateBranch(req *requests.UpdateBranch) error {
	return service.NewBranch().Update(req.ID, &req.BranchInfo)
}

func UpdateBranchStatus(id uint, status string) error {
	return service.NewBranch().UpdateBranchStatus(id, status)
}

func DeleteBranch(id uint) error {
	return service.NewBranch().Delete(id)
}

func BranchList(condition *commonModel.PageQuery[*requests.QueryBranch]) (*commonModel.ResPage[model.BranchInfo], error) {
	return service.NewBranch().ListBranch(condition)
}

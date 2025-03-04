package logic

import (
	"kiwi/internal/app/project/model"
	"kiwi/internal/app/project/requests"
	"kiwi/internal/app/project/service"
	commonModel "kiwi/internal/common/model"
)

func CreateTag(req *requests.CreateTag) error {
	srv := service.NewTag()
	return srv.Create(&req.TagInfo)
}

func UpdateTag(req *requests.UpdateTag) error {
	return service.NewTag().Update(req.ID, &req.TagInfo)
}

func UpdateTagReleaseStatus(id uint, status string) error {
	return service.NewTag().UpdateTagReleaseStatus(id, status)
}

func DeleteTag(id uint) error {
	return service.NewTag().Delete(id)
}

func TagList(condition *commonModel.PageQuery[*requests.QueryTag]) (*commonModel.ResPage[model.TagInfo], error) {
	return service.NewTag().ListTag(condition)
}

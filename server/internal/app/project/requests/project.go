package requests

import (
	"encoding/json"
	"kiwi/internal/app/project/model"
	commonModel "kiwi/internal/common/model"
)

type CreateProject struct {
	model.ProjectInfo
	Robots []uint `json:"robots" form:"robots"` // 机器人ID列表
}

func NewCreateProject() *CreateProject {
	return &CreateProject{
		ProjectInfo: model.ProjectInfo{},
		Robots:      []uint{},
	}
}

type UpdateProject struct {
	model.ProjectInfo
	Robots []uint `json:"robots" form:"robots"` // 机器人ID列表
}

func NewUpdateProject() *UpdateProject {
	return &UpdateProject{
		ProjectInfo: model.ProjectInfo{},
		Robots:      []uint{},
	}
}

func (req *CreateProject) ToJson() string {
	data, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(data)
}

type QueryProject struct {
	KeyWords string `json:"keywords" form:"keywords"` // 关键词
}

func NewQueryProject() *commonModel.PageQuery[*QueryProject] {
	return commonModel.NewPageQuery(0, 0, &QueryProject{})
}

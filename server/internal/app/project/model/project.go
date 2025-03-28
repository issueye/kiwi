package model

import (
	"kiwi/internal/common/model"
)

type ProjectInfo struct {
	model.BaseModel
	ProjectBase
}

type ProjectBase struct {
	Name          string             `gorm:"column:name;size:200;not null;comment:名称;" json:"name"`
	RepoUrl       string             `gorm:"column:repo_url;size:200;not null;comment:git 仓库地址;" json:"repo_url"`
	Description   string             `gorm:"column:description;size:200;not null;comment:项目描述;" json:"description"`
	Branchs       []BranchInfo       `gorm:"foreignKey:ProjectId;references:id;" json:"branchs"`
	ProjectRobots []ProjectRobotInfo `gorm:"foreignKey:ProjectId;references:id;" json:"project_robots"` // 项目机器人
}

func (p ProjectInfo) TableName() string { return "biz_project_info" }

type ProjectRobotInfo struct {
	model.BaseModel
	ProjectRobotBase
}

type ProjectRobotBase struct {
	ProjectId uint `gorm:"column:project_id;size:200;not null;comment:项目ID;" json:"project_id"` // 项目ID
	RobotId   uint `gorm:"column:robot_id;size:200;not null;comment:机器人ID;" json:"robot_id"`    // 机器人ID
}

func (p ProjectRobotInfo) TableName() string { return "biz_project_robot_info" }

// 分支信息
type BranchInfo struct {
	model.BaseModel
	BranchBase
}

type BranchBase struct {
	ProjectId   uint   `gorm:"column:project_id;size:200;not null;comment:项目ID;" json:"project_id"`             // 项目ID
	Name        string `gorm:"column:name;size:200;not null;comment:分支名称;" json:"name"`                         // 分支名称
	Hash        string `gorm:"column:hash;size:200;not null;comment:Hash;" json:"hash"`                         // Hash 值
	Description string `gorm:"column:description;size:200;not null;comment:分支描述;" json:"description"`           // 分支描述
	Status      string `gorm:"column:status;size:50;not null;default:'developing';comment:分支状态;" json:"status"` // 分支状态: developing, merged, closed
}

func (p BranchInfo) TableName() string { return "biz_project_branch_info" }

package initialize

import (
	"kiwi/internal/app/admin/model"
	"kiwi/internal/app/admin/service"
	"log/slog"
)

func InitRoleMenus() {
	rms := []*model.RoleMenu{
		{RoleCode: "9001", MenuCode: "1000"},
		{RoleCode: "9001", MenuCode: "1001"},
		{RoleCode: "9001", MenuCode: "1002"},
		{RoleCode: "9001", MenuCode: "9000"},
		{RoleCode: "9001", MenuCode: "9001"},
		{RoleCode: "9001", MenuCode: "9002"},
		{RoleCode: "9001", MenuCode: "9003"},
		{RoleCode: "9001", MenuCode: "9004"},
	}

	for _, rm := range rms {
		RMIsNotExistAdd(rm)
	}
}

func RMIsNotExistAdd(rm *model.RoleMenu) {
	RoleSrv := service.NewUser()
	isHave, err := RoleSrv.CheckRoleMenu(rm.RoleCode, rm.MenuCode)
	if err != nil {
		slog.Error("查询角色菜单失败", slog.String("失败原因", err.Error()))
		return
	}

	if !isHave {
		err = RoleSrv.AddRoleMenu(rm)
		if err != nil {
			slog.Error("添加角色菜单失败", slog.String("失败原因", err.Error()))
			return
		}
	}
}

func InitRoles() {
	Roles := []*model.Role{
		model.NewRole("9001", "管理员"),
	}

	for _, Role := range Roles {
		RoleIsNotExistAdd(Role)
	}
}

func RoleIsNotExistAdd(Role *model.Role) {
	RoleSrv := service.NewUser()
	info, err := RoleSrv.GetRoleByName(Role.Name)
	if err != nil {
		slog.Error("查询角色失败", slog.String("失败原因", err.Error()))
		return
	}

	if info.ID == 0 {
		err = RoleSrv.AddRole(Role)
		if err != nil {
			slog.Error("添加角色失败", slog.String("失败原因", err.Error()))
			return
		}
	}
}

func InitUserRole() {
	userRole := []*model.UserRole{
		{UserID: 1, RoleCode: "9001"},
	}

	for _, ur := range userRole {
		URIsNotExistAdd(ur)
	}
}

func URIsNotExistAdd(ur *model.UserRole) {
	RoleSrv := service.NewUser()
	isHave, err := RoleSrv.CheckUserRole(int(ur.UserID), ur.RoleCode)
	if err != nil {
		slog.Error("查询用户角色失败", slog.String("失败原因", err.Error()))
		return
	}

	if !isHave {
		err = RoleSrv.AddUserRole(ur)
		if err != nil {
			slog.Error("添加用户角色失败", slog.String("失败原因", err.Error()))
			return
		}
	}
}

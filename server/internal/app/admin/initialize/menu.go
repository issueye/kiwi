package initialize

import (
	"kiwi/internal/app/admin/model"
	"kiwi/internal/app/admin/service"
	"log/slog"
)

func InitMenus() {
	menus := []*model.Menu{
		model.BaseNewMenu(model.MenuBase{Code: "1000", Name: "项目管理", Description: "项目管理", Frontpath: "/project", Order: 20, Visible: true, Icon: "Folder", ParentCode: ""}),
		model.BaseNewMenu(model.MenuBase{Code: "1001", Name: "项目", Description: "项目管理", Frontpath: "/project/index", Order: 20, Visible: true, Icon: "Folder", ParentCode: "1000"}),
		model.BaseNewMenu(model.MenuBase{Code: "1002", Name: "机器人管理", Description: "机器人管理", Frontpath: "/project/robot", Order: 20, Visible: true, Icon: "User", ParentCode: "1000"}),

		model.BaseNewMenu(model.MenuBase{Code: "9000", Name: "系统管理", Description: "系统管理", Frontpath: "/system", Order: 90, Visible: true, Icon: "Setting", ParentCode: ""}),
		model.BaseNewMenu(model.MenuBase{Code: "9001", Name: "用户管理", Description: "用户管理", Frontpath: "/system/user", Order: 91, Visible: true, Icon: "User", ParentCode: "9000"}),
		model.BaseNewMenu(model.MenuBase{Code: "9002", Name: "角色管理", Description: "角色管理", Frontpath: "/system/role", Order: 92, Visible: true, Icon: "Avatar", ParentCode: "9000"}),
		model.BaseNewMenu(model.MenuBase{Code: "9003", Name: "菜单管理", Description: "菜单管理", Frontpath: "/system/menu", Order: 93, Visible: true, Icon: "Menu", ParentCode: "9000"}),
		model.BaseNewMenu(model.MenuBase{Code: "9004", Name: "字典管理", Description: "字典管理", Frontpath: "/system/dict_mana", Order: 94, Visible: true, Icon: "List", ParentCode: "9000"}),
		model.BaseNewMenu(model.MenuBase{Code: "9005", Name: "系统设置", Description: "系统设置", Frontpath: "/system/setting", Order: 95, Visible: true, Icon: "Tools", ParentCode: "9000"}),
	}

	for _, menu := range menus {
		MenuIsNotExistAdd(menu)
	}
}

func MenuIsNotExistAdd(menu *model.Menu) {
	menuSrv := service.NewMenu()

	isHave, err := menuSrv.CheckMenuExist(menu)
	if err != nil {
		slog.Error("检查菜单是否存在失败", slog.String("失败原因", err.Error()))
		return
	}

	if !isHave {
		err = menuSrv.AddMenu(menu)
		if err != nil {
			slog.Error("添加菜单失败", slog.String("失败原因", err.Error()))
		}
	}
}

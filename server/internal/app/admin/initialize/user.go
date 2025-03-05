package initialize

import (
	"kiwi/internal/app/admin/logic"
	"kiwi/internal/app/admin/model"
	"kiwi/internal/app/admin/service"
	"kiwi/internal/global"
	"log/slog"
)

// 初始化管理员用户数据
func InitAdminUser() {
	// 检查是否已经存在管理员用户
	adminUser, err := service.NewUser().GetUserByName("admin")
	if err != nil {
		return
	}

	if adminUser.ID != 0 {
		return
	}

	// 创建管理员用户
	password, err := logic.MakePassword(global.DEFAULT_PWD)
	if err != nil {
		slog.Error("生成密码哈希失败", slog.String("失败原因", err.Error()))
		return
	}

	user := model.User{
		Username: "admin",
		Password: password,
		NickName: "管理员",
		Avatar:   "https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png",
	}

	err = service.NewUser().AddUser(&user)
	if err != nil {
		slog.Error("创建管理员用户失败: %s", slog.String("失败原因", err.Error()))
		return
	}
}

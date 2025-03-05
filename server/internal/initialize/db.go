package initialize

import (
	"kiwi/internal/app/admin/initialize"
	adminModel "kiwi/internal/app/admin/model"
	projectModel "kiwi/internal/app/project/model"
	"kiwi/internal/global"
	"kiwi/pkg/db"
	"log/slog"
	"path/filepath"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	"gorm.io/gorm"
)

func InitDB() {
	path := filepath.Join(global.ROOT_PATH, "data", "data.db")
	global.DB = db.InitSqlite(path)

	InitDATA(global.DB)
}

func InitDATA(db *gorm.DB) {
	db.AutoMigrate(&adminModel.User{})
	db.AutoMigrate(&adminModel.Role{})
	db.AutoMigrate(&adminModel.UserRole{})
	db.AutoMigrate(&adminModel.RoleMenu{})
	db.AutoMigrate(&adminModel.Menu{})
	db.AutoMigrate(&adminModel.DictsInfo{})
	db.AutoMigrate(&adminModel.DictDetail{})

	db.AutoMigrate(&projectModel.ProjectInfo{})
	db.AutoMigrate(&projectModel.BranchInfo{})
	db.AutoMigrate(&projectModel.TagInfo{})
	db.AutoMigrate(&projectModel.VersionInfo{})
	db.AutoMigrate(&projectModel.RobotInfo{})

	initialize.Init()
}

func FreeDB() {
	sqldb, err := global.DB.DB()
	if err != nil {
		slog.Error("close db error", slog.String("error msg", err.Error()))
	}

	if err = sqldb.Close(); err != nil {
		slog.Error("close db error", slog.String("error msg", err.Error()))
	}
}

package db

import (
	"log/slog"
	"time"

	sqlite "github.com/glebarez/sqlite"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

func InitSqlite(path string) *gorm.DB {
	newLogger := glogger.New(
		Writer{
			BPrint: true,
		},
		glogger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  glogger.Info,           // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)

	l := newLogger.LogMode(glogger.Silent)
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		Logger: l,
	})

	if err != nil {
		slog.Error("连接数据库异常", slog.String("错误信息", err.Error()))
		panic(err)
	}

	db = db.Debug()
	slog.Info("初始化sqlite数据库完成", slog.String("dsn", path))
	return db
}

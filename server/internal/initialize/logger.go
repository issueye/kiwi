package initialize

import (
	"fmt"
	"kiwi/internal/common/config"

	"log/slog"

	"gopkg.in/natefinch/lumberjack.v2"
)

var LevelMap = map[int]slog.Level{
	-1: slog.LevelDebug,
	0:  slog.LevelInfo,
	1:  slog.LevelWarn,
	2:  slog.LevelError,
}

func InitLogger() {
	path := config.GetParam(config.LOG, "path", "logs").String()
	lWriter := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s.log", path),
		MaxSize:    config.GetParam(config.LOG, "max-size", 100).Int(),
		MaxBackups: config.GetParam(config.LOG, "max-backups", 100).Int(),
		MaxAge:     config.GetParam(config.LOG, "max-age", 100).Int(),
		Compress:   config.GetParam(config.LOG, "compress", true).Bool(),
	}

	lel := config.GetParam(config.LOG, "level", 0).Int()
	var level = new(slog.LevelVar)
	level.Set(LevelMap[lel])

	fileHandler := slog.NewTextHandler(lWriter, &slog.HandlerOptions{Level: level})
	slog.SetDefault(slog.New(fileHandler))
}

package initialize

import (
	"kiwi/internal/common/config"
	"kiwi/internal/global"
	"os"
	"path/filepath"

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
	path = filepath.Join(global.ROOT_PATH, path, "server.log")
	lWriter := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    config.GetParam(config.LOG, "max-size", 100).Int(),
		MaxBackups: config.GetParam(config.LOG, "max-backups", 100).Int(),
		MaxAge:     config.GetParam(config.LOG, "max-age", 100).Int(),
		Compress:   config.GetParam(config.LOG, "compress", true).Bool(),
	}

	lel := config.GetParam(config.LOG, "level", 0).Int()
	var level = new(slog.LevelVar)
	level.Set(LevelMap[lel])

	w := &Writer{
		lWriter: lWriter,
	}

	if lel == -1 {
		w.isDebug = true
	}

	fileHandler := slog.NewTextHandler(w, &slog.HandlerOptions{Level: level})
	slog.SetDefault(slog.New(fileHandler))
}

type Writer struct {
	isDebug bool
	lWriter *lumberjack.Logger
}

func (w *Writer) Write(p []byte) (n int, err error) {
	// 如果是调试模式，则内容也需要输出到控制台
	if w.isDebug {
		os.Stdout.Write(p)
	}

	return w.lWriter.Write(p)
}

package pages

import (
	"kiwi/internal/global"
	"log/slog"
)

func WriteLog(log string) {
	slog.Info(log)
	global.MsgChannel <- log
}

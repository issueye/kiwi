package db

import (
	"log/slog"
	"strings"

	"github.com/spf13/cast"
)

type Config struct {
	Username string `json:"user"`     // 用户名称
	Password string `json:"password"` // 密码
	Host     string `json:"host"`     // 服务器地址
	Database string `json:"name"`     // 数据库
	Port     int    `json:"port"`     // 端口号
	LogMode  bool   `json:"logMode"`  // 日志模式
}

// Writer 封装的SQL打印
type Writer struct {
	BPrint bool
}

func (w Writer) Printf(format string, args ...interface{}) {
	if !w.BPrint {
		return
	}

	funcPath := args[0].(string)
	// 只保留最后两层代码目录
	fData := strings.Split(funcPath, "/")
	if len(fData) > 2 {
		funcPath = fData[len(fData)-2] + "/" + fData[len(fData)-1]
	}

	switch len(args) {
	case 3:
	case 4:
		{
			if args[2] == "-" {
				slog.Debug("SQL语句",
					slog.Float64("time", cast.ToFloat64(args[1])),
					slog.String("sql", cast.ToString(args[3])),
					slog.String("row", cast.ToString(args[2])),
					slog.String("file", funcPath),
				)
			} else {
				slog.Debug("SQL语句",
					slog.Float64("time", cast.ToFloat64(args[1])),
					slog.String("sql", cast.ToString(args[3])),
					slog.String("row", cast.ToString(args[2])),
					slog.String("file", funcPath),
				)
			}
		}
	case 5: // 错误SQL语句
		{
			// 判断如果是 SLOW SQL 则使用 warn 级别
			if cast.ToInt64(args[2]) > 200 {
				slog.Warn("[SLOW SQL]语句",
					slog.String("sql", cast.ToString(args[4])),
					slog.Float64("time", cast.ToFloat64(args[2])),
					slog.String("row", cast.ToString(args[3])),
					slog.String("file", funcPath),
				)
			} else {
				slog.Error("SQL语句",
					slog.String("sql", cast.ToString(args[4])),
					slog.String("msg", cast.ToString(args[1])),
					slog.Float64("time", cast.ToFloat64(args[2])),
					slog.String("row", cast.ToString(args[3])),
					slog.String("file", funcPath),
				)
			}
		}
	}
}

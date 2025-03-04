package global

import (
	"io/fs"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	MsgChannel = make(chan string, 10)
	HttpEngine *gin.Engine
	HttpServer *http.Server
	DB         *gorm.DB
	S_WEB      fs.FS
)

const (
	TOPIC_CONSOLE_LOG = "TOPIC_CONSOLE_LOG"
	ROOT_PATH         = "root"
	DEFAULT_PWD       = "123456"
	DB_Key            = "data_base:info"
)

var (
	STATIC_PATH = filepath.Join(ROOT_PATH, "static")
	TMP_PATH    = filepath.Join(ROOT_PATH, "tmp")
	PKG_PATH    = filepath.Join(ROOT_PATH, "packages")
	DATA_PATH   = filepath.Join(ROOT_PATH, "data")
)

func WriteLog(msg string) {
	MsgChannel <- msg
}

var (
	APP_NAME = "消息推送服务"
	VERSION  = "v1.0.0.1"
)

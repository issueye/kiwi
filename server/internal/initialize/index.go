package initialize

import (
	"context"
)

func RunServer(ctx context.Context) {
	InitRuntime()
	InitConfig()
	InitLogger()
	InitDB()
	InitHttpServer(ctx)
}

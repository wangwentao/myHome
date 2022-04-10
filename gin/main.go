package main

import (
	"context"
	"myHome/gin/configs"
	"myHome/gin/router"
	"myHome/gin/utils/logs"
)

func main() {

	ctx := context.Background()

	configs.InitSettings(ctx)
	defer configs.ReleaseSettings(ctx)

	engine := router.SetupRouter()
	err := engine.Run(":" + configs.Server.HttpPort)
	logs.Error(err).Msg("Start Server occure error!")
}

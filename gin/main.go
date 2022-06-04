package main

import (
	"myHome/gin/configs"
	"myHome/gin/router"
	"myHome/gin/utils/logs"
)

// @title Interface docs
// @version v1.0
// @description myHome project api docs
// @contact.name wangwentao
// @contact.email 37362762@qq.com
// @host localhost:8880
func main() {

	configs.InitSettings()
	defer configs.ReleaseSettings()

	engine := router.SetupRouter()
	err := engine.Run(":" + configs.Server.HttpPort)
	logs.Error(err).Msg("Start Server occure error!")
}

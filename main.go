package main

import (
	"ginchat/models"
	"ginchat/router"
	"ginchat/utils"
	"time"

	"github.com/spf13/viper"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	InitTimer()

	r := router.Router()
	r.Run(viper.GetString("port.server")) // listen and serve on 0.0.0.0:8081 (for windows "localhost:8081")
}

func InitTimer() {
	utils.Timer(time.Duration(viper.GetInt("timeout.DelayHeartbeat"))*time.Second, time.Duration(viper.GetInt("timeout.HeartbeatHz"))*time.Second, models.CleanConnection, "")
}

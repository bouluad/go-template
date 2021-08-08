package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-template/config"
	"github.com/go-template/internal/router"
)

func setConfiguration(configPath string) {
	config.Setup(configPath)
	gin.SetMode(config.GetConfig().Server.Mode)
}
func Run(configPath string) {
	if configPath == "" {
		configPath = "config/config.yml"
	}
	setConfiguration(configPath)
	conf := config.GetConfig()
	web := router.Setup()
	fmt.Println("Go API template Running on port " + conf.Server.Port)
	_ = web.Run(":" + conf.Server.Port)
}

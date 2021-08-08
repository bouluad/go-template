package main

import (
	_ "github.com/go-template/docs"
	"github.com/go-template/internal/api"
)

// @Golang API REST
// @version 1.0
// @description API REST in Golang with Gin Framework

// @contact.name BOULUAD Mohammed
// @contact.url https://bouluad.com
// @contact.email contact@bouluad.com

// @license.name MIT

// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	api.Run("")
}

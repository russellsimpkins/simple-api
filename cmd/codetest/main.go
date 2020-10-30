package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/russellsimpkins/simple-api/cmd/codetest/apis"
	"github.com/russellsimpkins/simple-api/cmd/codetest/config"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func main() {
	// load application configurations
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	// fmt.Println(config.Config.ConfigVar)

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/hello/:who", apis.GetHello)
	}

	r.Run(fmt.Sprintf(":%v", config.Config.ListenPort))
}

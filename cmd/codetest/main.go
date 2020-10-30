package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/russellsimpkins/simple-api/cmd/codetest/apis"
	"github.com/russellsimpkins/simple-api/cmd/codetest/config"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Simple Swagger API.
// @version 1.0
// @description Codetest API for Joe.
// @termsOfService http://swagger.io/terms

// @contact.name API Support
// @contact.email russellsimpkins@gmail.com

// @license.name MIT
// @license.url https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

// @BasePath /api/v1
func main() {

	// Loads the application configurations.
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	// Creates a router without any middleware by default.
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// enables swagger.
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// adds our apis.
	v1 := r.Group("/api/v1")
	{
		v1.GET("/hello/:who", apis.GetHello)
	}

	// starts the service.
	r.Run(fmt.Sprintf(":%v", config.Config.ListenPort))
}

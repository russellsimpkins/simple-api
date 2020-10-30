package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/russellsimpkins/simple-api/cmd/codetest/apis"
	"github.com/russellsimpkins/simple-api/cmd/codetest/config"
	_ "github.com/russellsimpkins/simple-api/cmd/codetest/docs"
	"github.com/russellsimpkins/simple-api/cmd/codetest/httputil"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
)

// @title Simple Swagger API.
// @version 1.0
// @description Codetest API for Joe.
// @termsOfService http://swagger.io/terms

// @contact.name API Support
// @contact.email russellsimpkins@gmail.com

// @license.name MIT
// @license.url https://github.com/russellsimpkins/simple-api/blob/master/LICENSE

// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
		v1.Use(auth())
		v1.GET("/hello/:who", apis.GetHello)
	}

	// starts the service.
	r.Run(fmt.Sprintf(":%v", config.Config.ListenPort))
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 {
			httputil.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			c.Abort()
		}
		if authHeader != config.Config.ApiKey {
			httputil.NewError(c, http.StatusUnauthorized, fmt.Errorf("this user isn't authorized to this operation: auth_key is: %s.", authHeader, config.Config.ApiKey))
			c.Abort()
		}
		c.Next()
	}
}

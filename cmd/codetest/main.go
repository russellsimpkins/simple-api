package main

import (
	"fmt"
	"github.com/russellsimpkins/simple-api/cmd/codtest/config"
)

func main() {
	// load application configurations
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	fmt.Println(config.Config.ConfigVar)
}
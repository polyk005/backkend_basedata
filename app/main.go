package main

import (
	"fmt"
	"src/internal/config"
)

func main() {
	config.LoadEnvironment()

	appConfig := config.NewConfig()

	fmt.Println(fmt.Sprintf("User: %s", appConfig.Database.User))
	fmt.Println(fmt.Sprintf("Host: %s", appConfig.Database.Host))
	fmt.Println(fmt.Sprintf("Password: %s", appConfig.Database.Password))
	fmt.Println(fmt.Sprintf("DbName: %s", appConfig.Database.DbName))
	fmt.Println(fmt.Sprintf("Ports: %d", appConfig.Database.Port))
}

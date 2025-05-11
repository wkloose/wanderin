package main

import (
	"fmt"
	"wanderin/config"
	"wanderin/cmd"
)

func main() {
	config.InitDB()
	config.InitEnv()
	router := cmd.SetupRouter()
	port := ":3000"
	fmt.Println("Server is running on port", port)
	router.Run(port)
}

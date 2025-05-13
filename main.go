package main

import (
	"fmt"
	"wanderin/config"
	"wanderin/cmd"
	"os"
)

func main() {
	config.InitDB()
	config.InitEnv()
	router := cmd.SetupRouter()
	port := os.Getenv("PORT")
	fmt.Println("Server is running on port", port)
	router.Run(port)
}

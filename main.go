package main

import (
	"os"
	"prakerja_b4/configs"
	"prakerja_b4/routes"
)

func init() {
	configs.LoadEnv()
	configs.ConnectDatabase()
}

func main() {
	e := routes.Init()
	e.Start(":" + getPort())
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	return port
}

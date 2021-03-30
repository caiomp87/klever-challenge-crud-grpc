package main

import (
	"backend/app"
	"backend/config"
)

func main() {
	config.LoadEnv()
	app.StartGrpcServer()
}

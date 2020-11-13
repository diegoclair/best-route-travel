package main

import (
	"os"

	"github.com/diegoclair/best-route-travel/server"
	"github.com/diegoclair/best-route-travel/service"
	"github.com/diegoclair/go_utils-lib/logger"
)

// PORT is the default port to start the application
const PORT string = "3000"

func main() {
	logger.Info("Reading the initial configs...")

	svc := service.New()
	svm := service.NewServiceManager()

	if len(os.Args) > 1 {
		defer os.Exit(0)
		if os.Args[1] == "cli" {
			svm.CommandLineService(svc, svm.TravelService(svc)).RunCLI()
			return
		}
	}

	server := server.InitServer(svc, svm)
	logger.Info("About to start the application...")

	port := os.Getenv("PORT")

	if port == "" {
		port = PORT
	}

	if err := server.Start(":" + port); err != nil {
		panic(err)
	}
}

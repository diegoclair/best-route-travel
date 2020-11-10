package main

// PORT is the default port to start the application
const PORT string = "3000"

func main() {
	// logger.Info("Reading the initial configs...")

	// svc := service.New()
	// server := server.InitServer(svc)
	// logger.Info("About to start the application...")

	// port := os.Getenv("PORT")

	// if port == "" {
	// 	port = PORT
	// }

	// if err := server.Start(":" + port); err != nil {
	// 	panic(err)
	// }
	comandLineInterface()
}

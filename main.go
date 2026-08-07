package main

import (
	"log"
	"os"

	"github.com/klass-lk/ginboot"
	"github.com/klass-lk/ginboot/runtime/lambda"
	"github.com/klass-lk/test/internal/di"
)

func main() {
	// Initialize Ginboot app (automatically loads .env, ginboot.yml, and configures telemetry)
	app := ginboot.New()

	// Initialize Lambda runner if running on AWS Lambda
	if os.Getenv("LAMBDA_TASK_ROOT") != "" {
		app.SetRunner(lambda.NewRunner())
	}

	app.SetBasePath("/api")
	di.NewContainer(app)

	// Start server
	if err := app.Start(8081); err != nil {
		log.Fatal(err)
	}
}

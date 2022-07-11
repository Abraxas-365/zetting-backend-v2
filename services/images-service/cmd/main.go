package main

import (
	"fmt"
	"os"
	internal_S3 "service-images/internal/s3"
	"service-images/pkg/application"
	"service-images/pkg/infrastructure/repository"
	handler "service-images/pkg/infrastructure/rest/handlers"
	"service-images/pkg/infrastructure/rest/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	mongoUri := os.Getenv("MONGODB_URI")
	awsId := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecret := os.Getenv("AWS_SECRET_ACCESS_KEY")
	// Initialize a session in us-west-2 that the SDK will use to load
	s3, err := internal_S3.NewS3Client(awsId, awsSecret, "zetting", "us-east-1")
	if err != nil {
		panic(err)
	}
	repo, _ := repository.NewMongoRepository(mongoUri, "Zetting", 10, "Images")
	application := application.NewApplication(repo, s3)

	app := fiber.New()
	app.Use(logger.New())
	handler := handler.NewHandler(application)
	routes.ServeRoute(app, handler)
	fmt.Println("inicando en puerto 3010")

	app.Listen(":3010")
}

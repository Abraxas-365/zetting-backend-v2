package main

import (
	"fmt"
	"os"
	"service-users/internal/rabbit"
	"service-users/pkg/user/application"
	"service-users/pkg/user/domain/service"
	"service-users/pkg/user/infrastructure/mqueue/publisher"
	"service-users/pkg/user/infrastructure/repository"
	"service-users/pkg/user/infrastructure/rest/handlers"
	"service-users/pkg/user/infrastructure/rest/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	mongoUri := os.Getenv("MONGODB_URI")

	mqUri := os.Getenv("MQ_URI")
	repo, _ := repository.NewMongoRepository(mongoUri, "Zetting", 10, "Users")
	service := service.NewUserService(repo)
	mq, err := rabbit.NewMQueueConection(mqUri)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	mqpublisher := mqpublisher.NewMQPublisher(mq)
	application := application.NewApp(repo, service, mqpublisher)
	handler := handlers.NewHandler(application)
	app := fiber.New()
	app.Use(logger.New())
	routes.UserRoutes(app, handler) //User routes

	app.Listen(":3000")
}

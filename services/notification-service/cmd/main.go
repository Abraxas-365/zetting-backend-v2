package main

import (
	"notifications/pkg/notification/application"
	"notifications/pkg/notification/infrastructure/repository"
	"notifications/pkg/notification/infrastructure/rest/handlers"
	"notifications/pkg/notification/infrastructure/rest/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	mongoUri := os.Getenv("MONGODB_URI")

	// mqUri := os.Getenv("MQ_URI")
	repo, _ := repository.NewMongoRepository(mongoUri, "Zetting", 10, "Notifications")
	// mq, err := rabbit.NewMQueueConection(mqUri)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	os.Exit(1)
	// }

	// mqpublisher := mqpublisher.NewMQPublisher(mq)
	application := application.NewNotificationApp(repo)
	handler := handlers.NewNoticationHandler(application)
	app := fiber.New()
	app.Use(logger.New())
	routes.NotificationsRoutes(app, handler) //User routes

	app.Listen(":3020")

}

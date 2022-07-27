package main

import (
	"fmt"
	"notifications/internal/rabbit"
	"notifications/pkg/notification/application"
	"notifications/pkg/notification/infrastructure/repository"
	"notifications/pkg/notification/infrastructure/rest/handlers"
	"notifications/pkg/notification/infrastructure/rest/routes"
	user_app "notifications/pkg/user/application"
	userMQHandler "notifications/pkg/user/infrastructure/mqueue/consumer/handlers"
	userMQRoutes "notifications/pkg/user/infrastructure/mqueue/consumer/routes"
	user_repo "notifications/pkg/user/infrastructure/repository"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	//Infrastructure
	mongoUri := os.Getenv("MONGODB_URI")
	mqUri := os.Getenv("MQ_URI")
	mq, err := rabbit.NewMQueueConection(mqUri)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	//repository
	repo, _ := repository.NewMongoRepository(mongoUri, "Zetting", 10, "Notifications")
	UserRepo, _ := user_repo.NewMongoRepository(mongoUri, "Zetting", 10, "NotificationUsers")

	//Application
	application := application.NewNotificationApp(repo)
	UserApp := user_app.NewUserApplication(UserRepo)

	//MQ Handlers
	UserMQHandler := userMQHandler.NewMQHandler(UserApp)

	// MQ Routes
	userMQRoutes.ConsumerRoutes(mq, UserMQHandler)

	handler := handlers.NewNoticationHandler(application)
	app := fiber.New()
	app.Use(logger.New())
	routes.NotificationsRoutes(app, handler) //User routes

	app.Listen(":3020")

}

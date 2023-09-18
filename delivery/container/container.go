package container

import (
	"eko-car/config"
	carFeature "eko-car/domain/car/feature"
	carRepository "eko-car/domain/car/repository"
	userFeature "eko-car/domain/user/feature"
	userRepository "eko-car/domain/user/repository"
	"eko-car/infrastructure/broker/rabbitmq"
	"eko-car/infrastructure/database"
	"eko-car/infrastructure/logger"
	"eko-car/infrastructure/service/queue"
	"eko-car/infrastructure/shared/constant"
	"fmt"
	"log"
)

type Container struct {
	EnvironmentConfig config.EnvironmentConfig
	RabbitMQ          rabbitmq.RabbitMQ
	QueueServices     queue.QueueService
	UserFeature       userFeature.UserFeature
	CarFeature        carFeature.CarFeature
}

func SetupContainer() Container {
	fmt.Println("Starting new container...")

	fmt.Println("Loading config...")
	config, err := config.LoadENVConfig()
	if err != nil {
		log.Panic(err)
	}

	logger.InitializeLogger(constant.LOGRUS) // choose which log, ZAP or LOGRUS. Default: LOGRUS

	fmt.Println("Loading database...")
	db, err := database.LoadDatabase(config.Database)
	if err != nil {
		log.Panic(err)
	}

	// fmt.Println("Loading message broker...")
	// rmq := rabbitmq.NewConnection(config.RabbitMq)
	// // Connect RabbitMQ
	// err = rmq.Connect()
	// if err != nil {
	// 	log.Panic(err)
	// }

	// fmt.Println("Loading service's...")
	// queueService := queue.NewQueueService(rmq, config.RabbitMq)

	fmt.Println("Loading repository's...")
	userRepository := userRepository.NewUserRepository(db)
	carRepository := carRepository.NewCarRepository(db)

	fmt.Println("Loading feature's...")
	userFeature := userFeature.NewUserFeature(userRepository)
	carFeature := carFeature.NewCarFeature(carRepository)

	return Container{
		EnvironmentConfig: config,
		UserFeature:       userFeature,
		CarFeature:        carFeature,
	}
}

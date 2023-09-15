package container

import (
	"eko-car/config"
	health_feature "eko-car/domain/health/feature"
	health_repository "eko-car/domain/health/repository"
	logistik_feature "eko-car/domain/logistik/feature"
	logistik_repository "eko-car/domain/logistik/repository"
	sales_feature "eko-car/domain/sales/feature"
	sales_repository "eko-car/domain/sales/repository"
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
	HealthFeature     health_feature.HealthFeature
	LogistikFeature   logistik_feature.LogistikFeature
	SalesFeature      sales_feature.SalesFeature
	QueueServices     queue.QueueService
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
	healthRepository := health_repository.NewHealthFeature(db)
	logistikRepository := logistik_repository.NewLogistikRepository(db)
	salesRepository := sales_repository.NewSalesRepository(db)

	fmt.Println("Loading feature's...")
	healthFeature := health_feature.NewHealthFeature(config, healthRepository)
	logistikFeature := logistik_feature.NewLogistikFeature(logistikRepository)
	salesFeature := sales_feature.NewSalesFeature(salesRepository)

	return Container{
		EnvironmentConfig: config,
		HealthFeature:     healthFeature,
		LogistikFeature:   logistikFeature,
		SalesFeature:      salesFeature,
	}
}

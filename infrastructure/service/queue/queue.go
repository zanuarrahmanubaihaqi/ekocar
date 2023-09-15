package queue

import (
	"eko-car/infrastructure/broker/rabbitmq"
	"context"
)

type QueueService interface {
	PublishData(ctx context.Context, topic string, msg interface{}) (err error)
	ConsumeData(ctx context.Context, topic string) (err error)
}

type queueService struct {
	rabbitmq rabbitmq.RabbitMQ
	cfg      rabbitmq.RabbitmqConfig
}

func NewQueueService(rabbitmq rabbitmq.RabbitMQ, cfg rabbitmq.RabbitmqConfig) QueueService {
	return &queueService{
		rabbitmq: rabbitmq,
		cfg:      cfg,
	}
}

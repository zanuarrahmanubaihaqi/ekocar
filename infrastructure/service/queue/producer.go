package queue

import (
	Error "eko-car/domain/shared/error"
	"eko-car/infrastructure/logger"
	"eko-car/infrastructure/shared/constant"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func (q queueService) PublishData(ctx context.Context, topic string, msg interface{}) (err error) {

	cfg := q.rabbitmq.GetConfig()

	select {
	case err := <-cfg.Err:
		if err != nil {
			q.rabbitmq.Reconnect()
		}
	default:
	}

	ch, err := cfg.Conn.Channel()
	if err != nil {
		err = Error.New(constant.PUBLISHER_RABBITMQ, constant.ErrDefineChannelToBroker, err)
		return
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		topic,                           // name
		constant.RABBITMQ_EXCHANGE_TYPE, // type
		true,                            // durable
		false,                           // auto-deleted
		false,                           // internal
		false,                           // no-wait
		nil,                             // arguments
	)
	if err != nil {
		err = Error.New(constant.PUBLISHER_RABBITMQ, constant.ErrDeclareExhangeToBroker, err)
		return
	}

	body, err := json.Marshal(msg)
	if err != nil {
		err = Error.New(constant.PUBLISHER_RABBITMQ, constant.ErrToMarshalJSON, err)
		return
	}

	err = cfg.Channel.Publish(
		topic, // exchange
		"",    // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,    // keeping message if broker restart
			ContentType:  "application/json", // XXX: We will revisit this in future episodes
			Body:         body,
			Timestamp:    time.Now(),
		})
	if err != nil {
		err = Error.New(constant.PUBLISHER_RABBITMQ, constant.ErrPublishQueueToBroker, err)
		return
	}

	fmt.Println(fmt.Sprintf(constant.SUCCESS_PUBLISH_TO_BROKER, topic))
	logger.LogInfo(constant.PUBLISHER_RABBITMQ, fmt.Sprintf(constant.SUCCESS_PUBLISH_TO_BROKER, topic))

	return

}

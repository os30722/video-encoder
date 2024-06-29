package mom

import (
	"context"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	timeOut  = 5 * time.Second
	parallel = 8
)

var (
	chanel   *amqp.Channel
	taskName string = "tasks"
)

func init() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	err = ch.Qos(
		parallel,
		0,
		false,
	)

	if err != nil {
		panic(err)
	}
	chanel = ch

	CreateTaskQueue()
}

func CreateTaskQueue() {
	_, err := chanel.QueueDeclare(
		taskName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

}

func GetTaskMsg() (<-chan amqp.Delivery, error) {
	msgChan, err := chanel.Consume(
		taskName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return msgChan, nil
}

func PublishTask(ctx context.Context, input interface{}) error {
	return PublishJson(ctx, taskName, input)
}

func PublishJson(ctx context.Context, queeName string, input interface{}) error {
	buffer, err := json.Marshal(input)
	if err != nil {
		return err
	}
	ctx, _ = context.WithTimeout(ctx, timeOut)
	err = chanel.PublishWithContext(ctx,
		"",
		queeName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        buffer,
		})
	if err != nil {
		return err
	}
	return err
}

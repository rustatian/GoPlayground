package transport

import (
	"github.com/48d90782/GoPlayground/secret/pkg/shared_data"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func NewRabbitMQHandler(conn *amqp.Connection, logger *zap.Logger, shared shared_data.Shared) error {
	return nil
}

package consumer

import (
	"context"
	"errors"
	"strings"

	sarama "github.com/IBM/sarama"

	logger "github.com/MGomed/common/logger"
)

type consumer struct {
	log                  logger.Interface
	consumerGroup        sarama.ConsumerGroup
	consumerGroupHandler *GroupHandler
}

// NewConsumer is consumer struct constructor
func NewConsumer(
	log logger.Interface,
	consumerGroup sarama.ConsumerGroup,
	consumerGroupHandler *GroupHandler,
) *consumer {
	return &consumer{
		log:                  log,
		consumerGroup:        consumerGroup,
		consumerGroupHandler: consumerGroupHandler,
	}
}

// Consume consumes messages from topic and handle it
func (c *consumer) Consume(ctx context.Context, topic string, handler Handler) error {
	c.consumerGroupHandler.msgHandler = handler

	return c.consume(ctx, topic)
}

// Close closes consumer connection
func (c *consumer) Close() error {
	return c.consumerGroup.Close()
}

func (c *consumer) consume(ctx context.Context, topic string) error {
	for {
		err := c.consumerGroup.Consume(ctx, strings.Split(topic, ","), c.consumerGroupHandler)
		if err != nil {
			if errors.Is(err, sarama.ErrClosedConsumerGroup) {
				return nil
			}

			return err
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}

		c.log.Info("Rebalancing...")
	}
}

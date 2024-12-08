package consumer

import (
	"context"

	sarama "github.com/IBM/sarama"

	logger "github.com/MGomed/common/logger"
)

// Handler defines message handling func
type Handler func(ctx context.Context, msg *sarama.ConsumerMessage) error

// GroupHandler is consumer group handler
type GroupHandler struct {
	log        logger.Interface
	msgHandler Handler
}

// NewGroupHandler is GroupHandler struct constructor
func NewGroupHandler(logger logger.Interface) *GroupHandler {
	return &GroupHandler{
		log: logger,
	}
}

// Setup is started at the beginning of a new session before ConsumeClaim is called
func (c *GroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of the session life after all ConsumeClaim goroutines are completed
func (c *GroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim should start a consumer message loop ConsumerGroupClaim().
// After the Messages() channel is closed, the handler must complete the processing
func (c *GroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				c.log.Warn("Message channel was closed")

				return nil
			}

			c.log.Debug("Message claimed: value = %s, timestamp = %v, topic = %s",
				string(message.Value), message.Timestamp, message.Topic)

			err := c.msgHandler(session.Context(), message)
			if err != nil {
				c.log.Error("Error handling message: %v", err)

				continue
			}

			session.MarkMessage(message, "")

		case <-session.Context().Done():
			c.log.Warn("Session context done")

			return nil
		}
	}
}

package playercommandproducer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Denisius664/room-service/internal/models"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)

type PlayerCommandProducer struct {
	writer *kafka.Writer
	topic  string
}

// NewRoomEventProducer creates a producer that writes to the provided brokers and topic.
func NewPlayerCommandProducer(brokers []string, topic string) *PlayerCommandProducer {
	w := &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.Hash{},
		// reasonable defaults
		Async: false,
	}
	return &PlayerCommandProducer{writer: w, topic: topic}
}

// Produce sends the RoomEvent to Kafka as a JSON-encoded message. The event Name is used as the message key.
func (p *PlayerCommandProducer) Produce(ctx context.Context, cmd *models.PlayerCommand) error {
	if cmd == nil {
		return errors.New("cmd is nil")
	}
	b, err := json.Marshal(cmd)
	if err != nil {
		return errors.Wrap(err, "marshal cmd")
	}

	msg := kafka.Message{
		Key:   []byte(cmd.Name),
		Value: b,
		Time:  time.Now(),
	}

	// WriteMessage is blocking and returns when message is accepted by broker
	if err := p.writer.WriteMessages(ctx, msg); err != nil {
		return errors.Wrap(err, "write message to kafka")
	}
	return nil
}

// Close closes the underlying writer.
func (p *PlayerCommandProducer) Close(ctx context.Context) error {
	if p.writer == nil {
		return nil
	}
	return p.writer.Close()
}

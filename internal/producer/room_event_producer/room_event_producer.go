package roomeventproducer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"

	"github.com/Denisius664/room-service/internal/models"
)

// RoomEventProducer publishes RoomEvent messages to Kafka.
type RoomEventProducer struct {
	writer *kafka.Writer
	topic  string
}

// NewRoomEventProducer creates a producer that writes to the provided brokers and topic.
func NewRoomEventProducer(brokers []string, topic string) *RoomEventProducer {
	w := &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.Hash{},
		// reasonable defaults
		Async: false,
	}
	return &RoomEventProducer{writer: w, topic: topic}
}

// Produce sends the RoomEvent to Kafka as a JSON-encoded message. The event Name is used as the message key.
func (p *RoomEventProducer) Produce(ctx context.Context, event *models.RoomEvent) error {
	if event == nil {
		return errors.New("event is nil")
	}
	b, err := json.Marshal(event)
	if err != nil {
		return errors.Wrap(err, "marshal event")
	}

	msg := kafka.Message{
		Key:   []byte(event.Name),
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
func (p *RoomEventProducer) Close(ctx context.Context) error {
	if p.writer == nil {
		return nil
	}
	return p.writer.Close()
}

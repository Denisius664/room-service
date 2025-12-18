package roomeventproducer_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/segmentio/kafka-go"

	"github.com/Denisius664/room-service/config"
)

// Тест проверяет:
// 1) Можно подключиться к брокеру
// 2) Можно отправить сообщение
// 3) Можно прочитать сообщение обратно
func TestKafkaConnection(t *testing.T) {
	cfg, err := config.LoadConfig(os.Getenv("configPath"))
	if os.Getenv("configPath") == "" {
		t.Skip("integration test skipped; set configPath to run")
	}
	if err != nil {
		panic(fmt.Sprintf("ошибка парсинга конфига, %v", err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	kafkaBrockers := []string{fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)}

	// Создаем writer
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  kafkaBrockers,
		Topic:    cfg.Kafka.RoomEventsTopic,
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	// Пытаемся отправить сообщение
	err = writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte("ping"),
		Value: []byte("pong"),
	})
	if err != nil {
		t.Fatalf("❌ Не удалось отправить сообщение в Kafka: %v", err)
	}

	// Создаем reader
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   kafkaBrockers,
		Topic:     cfg.Kafka.RoomEventsTopic,
		Partition: 0,
		MinBytes:  1,
		MaxBytes:  10e6,
	})
	defer reader.Close()

	// Читаем сообщение
	msg, err := reader.ReadMessage(ctx)
	if err != nil {
		t.Fatalf("❌ Не удалось прочитать сообщение из Kafka: %v", err)
	}

	// Проверяем, что получили то, что отправили
	if string(msg.Value) != "pong" {
		t.Fatalf("❌ Неверное значение: ожидалось 'pong', получено '%s'", string(msg.Value))
	}

	t.Log("✅ Kafka connection OK — сообщение успешно отправлено и получено")
}

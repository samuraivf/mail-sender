package kafka

import (
	"context"

	kafkago "github.com/segmentio/kafka-go"

	"github.com/samuraivf/mail-sender/internal/app/mail-sender/log"
)

type KafkaReader struct {
	reader *kafkago.Reader
	log    log.Log
}

type Kafka interface {
	Close() error
	Read(ctx context.Context) (kafkago.Message, error)
}

func NewKafkaReader(config kafkago.ReaderConfig, log log.Log) Kafka {
	return &KafkaReader{
		reader: kafkago.NewReader(config),
		log:    log,
	}
}

func (r *KafkaReader) Read(ctx context.Context) (kafkago.Message, error) {
	m, err := r.reader.ReadMessage(ctx)
	if err != nil {
		r.log.Error(err)
		return kafkago.Message{}, err
	}
	return m, nil
}

func (r *KafkaReader) Close() error {
	err := r.reader.Close()
	if err != nil {
		r.log.Error(err)
	}
	r.log.Info("Closed Kafka connection")
	return nil
}

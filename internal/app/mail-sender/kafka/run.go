package kafka

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/samuraivf/mail-sender/configs"
	"github.com/samuraivf/mail-sender/internal/app/mail-sender/log"
)

func Run() {
	logger := log.New()
	reader := NewKafkaReader(configs.KafkaConfig(), logger)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	go func() {
		for {
			m, err := reader.Read(ctx)
			if err != nil {
				logger.Error(err)
				break
			}
			logger.Infof("Received message from %s-%d [%d]: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		}
	}()

	<-done

	logger.Info("Server stopped")
	reader.Close()
	cancel()
}

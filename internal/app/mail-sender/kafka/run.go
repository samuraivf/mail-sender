package kafka

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/samuraivf/mail-sender/configs"
	"github.com/samuraivf/mail-sender/internal/app/mail-sender/log"
	"github.com/samuraivf/mail-sender/internal/app/mail-sender/mail"
)

func Run() {
	logger := log.New()
	reader := NewKafkaReader(configs.KafkaConfig(), logger)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	ctx := context.Background()

	go func() {
		for {
			m, err := reader.Read(ctx)
			if err != nil {
				logger.Error(err)
				break
			}
			logger.Infof("Received message from %s-%d [%d]: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

			go func(msg string) {
				sender := mail.NewSender(configs.MailSenderConfig())
				if err := sender.Send(msg); err != nil {
					logger.Error(err)
				}
			}(string(m.Value))
		}
	}()

	<-done

	logger.Info("Server stopped")
	reader.Close()
}

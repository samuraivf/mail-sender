package configs

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/samuraivf/mail-sender/internal/app/mail-sender/mail"
	kafkago "github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func init() {
	if err := initConfig(); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("")
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal().Timestamp().Err(err).Msg("")
	}
}

func MailSenderConfig() mail.SenderConfig {
	return mail.SenderConfig{
		FromEmail:    os.Getenv("MAIL_FROM"),
		FromAppPassword: os.Getenv("MAIL_FROM_APP_PASSWORD"),
		SMTPHost:     os.Getenv("SMTP_HOST"),
		SMTPPort:     os.Getenv("SMTP_PORT"),
	}
}

func KafkaConfig() kafkago.ReaderConfig {
	return kafkago.ReaderConfig{
		Brokers:  []string{viper.GetString("kafka.brokers")},
		Topic:    viper.GetString("kafka.topic"),
		GroupID:  viper.GetString("kafka.group"),
		MaxWait:  500 * time.Millisecond,
		MinBytes: 1,
		MaxBytes: 1e6,
	}
}

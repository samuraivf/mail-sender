package configs

import (
	"time"

	"github.com/rs/zerolog/log"
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

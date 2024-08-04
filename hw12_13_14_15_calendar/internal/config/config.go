package config

import (
	"encoding/json"
	"os"

	"github.com/imega/mt"
)

type Config struct {
	Logger   LoggerConf
	Server   ServerConf
	DB       DBConf
	Grpc     GrpcConf
	AMQP     AMQPConf
	Sheduler ShedulerConf

	Error error `json:"-"`
}

type LoggerConf struct {
	Level string
}

type ServerConf struct {
	Host string
	Port int64
}

type GrpcConf struct {
	Host string
}

type DBConf struct {
	Use      bool
	Host     string
	Port     int64
	User     string
	Password string
}

type AMQPConf struct {
	DSN         string `json:"dsn"`
	ServiceName string
	MtConfig    mt.Config `json:"mt_config"`
}

type ShedulerConf struct {
	Interval int
}

func NewConfig(filepath string) Config {
	var config Config

	bb, err := os.ReadFile(filepath)
	if err != nil {
		config.Error = err
		return config
	}

	err = json.Unmarshal(bb, &config)
	config.Error = err

	return config
}

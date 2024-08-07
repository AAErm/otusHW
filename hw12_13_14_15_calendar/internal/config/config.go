package config

import (
	"encoding/json"
	"os"

	"github.com/imega/mt"
)

type Config struct {
	Logger    LoggerConf
	Server    ServerConf
	DB        DBConf
	Grpc      GrpcConf
	AMQP      AMQPConf
	Scheduler SchedulerConf

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
	DbName   string `json:"db_name"`
}

type AMQPConf struct {
	DSN         string    `json:"dsn"`
	ServiceName string    `json:"service_name"`
	MtConfig    mt.Config `json:"mt_config"`
}

type SchedulerConf struct {
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

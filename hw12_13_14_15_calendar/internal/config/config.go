package config

import "github.com/spf13/viper"

type Config struct {
	Logger LoggerConf
	Server ServerConf
	Sql    SqlConf

	Error error `json:"-"`
}

type LoggerConf struct {
	Level string
}

type ServerConf struct {
	Host string
	Port int64
}

type SqlConf struct {
	Use      bool
	Host     string
	Port     int64
	User     string
	Password string
}

func NewConfig(filepath string) Config {
	config := Config{}

	viper.SetConfigFile(filepath)

	err := viper.ReadInConfig()
	if err != nil {
		config.Error = err
		return config
	}

	err = viper.Unmarshal(&config)
	config.Error = err

	return config
}

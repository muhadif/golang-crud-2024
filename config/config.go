package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	AppSecretKey []byte `env:"APP_SECRET_KEY"`

	MySqlDatabase string `env:"MYSQL_DATABASE"`
	MySqlUsername string `env:"MYSQL_USERNAME"`
	MySqlPassword string `env:"MYSQL_PASSWORD"`
	MySqlPort     string `env:"MYSQL_PORT"`
	MySqlHost     string `env:"MYSQL_HOST"`

	ServiceURL string `env:"SERVICE_URL"`

	UserServiceURL string `env:"USER_SERVICE_URL"`
	HttpPort       string `env:"HTTP_PORT"`
	JWTExpiredTime int    `env:"JWT_EXPIRED_TIME"`
}

func LoadConfig() *Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(cfg.HttpPort)
	return &cfg
}

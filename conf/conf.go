package conf

import (
	"github.com/jinzhu/configor"
	"github.com/rs/zerolog/log"
)

type Config struct {
	ListenAddr string `default:":8000" env:"LISTEN_ADDR"`
	DB         struct {
		User     string `required:"true" env:"DB_USER"`
		Password string `required:"true" env:"DB_PASSWORD"`
		Port     int    `default:"5432"  env:"DB_PORT"`
		Host     string `required:"true" env:"DB_ADDR"`
		Name     string `required:"true" env:"DB_NAME"`
		PoolSize int    `default:"10"    env:"DB_POOL_SIZE"`
	}
	Debug bool `default:"false" env:"DEBUG"`
}

func New() *Config {
	c := &Config{}
	if err := configor.New(&configor.Config{}).Load(c); err != nil {
		log.Fatal().Err(err).Msg("App config validation error")
	}

	return c
}

package configs

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/shabohin/holiday.git/internal/domain/errs"
)

type database struct {
	URI string `env:"DATABASE_URI" toml:"uri"`
}

type Config struct {
	BindAddr string   `env:"BIND_ADDR" toml:"bind_addr" env-default:":8000"`
	LogLevel string   `env:"LOG_LEVEL" toml:"log_level" env-default:"debug"`
	Database database `toml:"database1"`
}

func ParseConfig(configPath string) (*Config, error) {
	config := &Config{}
	if configPath != "" {
		if err := cleanenv.ReadConfig(configPath, config); err != nil {
			return nil, errs.NewUnexpectedBehaviorError(err.Error())
		}
	} else {
		if err := cleanenv.ReadEnv(config); err != nil {
			return nil, errs.NewUnexpectedBehaviorError(err.Error())
		}
	}
	return config, nil
}

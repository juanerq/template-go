package config

import env "github.com/caarlos0/env/v11"

type EnvConfig struct {
	ControllerPath string `env:"CONTROLLER_PATH,required"`
	Port           string `env:"PORT,required"`
	ServiceName    string `env:"SERVICE_NAME,required"`
	ServicePrefix  string `env:"SERVICE_PREFIX,required"`
	Timeout        int    `env:"TIMEOUT,required"`
}

func LoadConfig() (*EnvConfig, error) {
	cfg := EnvConfig{}

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

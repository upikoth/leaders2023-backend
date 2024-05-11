package store

import "github.com/kelseyhightower/envconfig"

type Config struct {
	YdbDsn          string `envconfig:"YDB_DSN" required:"true"`
	YdbAuthInfo     []byte `envconfig:"YDB_AUTH_INFO" required:"true"`
	YdbAuthFileName string `envconfig:"YDB_AUTH_FILE_NAME" required:"true"`
}

func NewConfig() (*Config, error) {
	config := &Config{}

	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}

package store

import "github.com/kelseyhightower/envconfig"

type Config struct {
	YdbDsn          string `envconfig:"YDB_DSN" required:"true"`
	YdbAuthFileName string `envconfig:"YDB_AUTH_FILE_NAME" required:"true"`
	YdbAuthInfo     []byte `envconfig:"YDB_AUTH_INFO"`
}

func NewConfig() (*Config, error) {
	config := &Config{}

	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}

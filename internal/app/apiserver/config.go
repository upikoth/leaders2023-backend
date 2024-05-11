package apiserver

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port              string `envconfig:"APP_PORT" required:"true"`
	JwtSecret         []byte `envconfig:"JWT_SECRET" required:"true"`
	S3Region          string `envconfig:"S3_REGION" required:"true"`
	S3Endpoint        string `envconfig:"S3_ENDPOINT" required:"true"`
	S3AccessKeyID     string `envconfig:"S3_ACCESS_KEY_ID" required:"true"`
	S3SecretAccessKey string `envconfig:"S3_SECRET_ACCESS_KEY" required:"true"`
}

func NewConfig() (*Config, error) {
	config := &Config{}

	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}

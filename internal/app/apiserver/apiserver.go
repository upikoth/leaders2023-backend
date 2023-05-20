package apiserver

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/gin-gonic/gin"
	"github.com/upikoth/leaders2023-backend/internal/app/handler"
	"github.com/upikoth/leaders2023-backend/internal/app/store"
)

type ApiServer struct {
	config           *Config
	router           *gin.Engine
	handler          *handler.Handler
	store            *store.Store
	dadataSuggestApi *suggest.Api
	s3               *s3.S3
}

func New(config *Config) *ApiServer {
	store := store.New()
	dadataSuggestApi := dadata.NewSuggestApi()

	session := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:   aws.String(config.S3Region),
			Endpoint: aws.String(config.S3Endpoint),
			Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
				AccessKeyID:     config.S3AccessKeyId,
				SecretAccessKey: config.S3SecretAccessKey,
			}),
		},
	}))

	s3 := s3.New(session)
	handlerEnv := &handler.HandlerEnv{
		JwtSecret:          config.JwtSecret,
		S3AccessDomainName: config.S3AccessDomainName,
	}

	handler := handler.New(store, handlerEnv, dadataSuggestApi, s3)

	return &ApiServer{
		config:           config,
		router:           gin.Default(),
		handler:          handler,
		store:            store,
		dadataSuggestApi: dadataSuggestApi,
		s3:               s3,
	}
}

func (s *ApiServer) Start() error {
	s.initRoutes()
	err := s.store.Connect()

	defer func() {
		disconnectErr := s.store.Disconnect()

		if disconnectErr != nil {
			log.Println(disconnectErr)
		}
	}()

	if err != nil {
		return err
	}

	return s.router.Run(":" + s.config.Port)
}

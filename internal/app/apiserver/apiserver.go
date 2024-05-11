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

type APIServer struct {
	config           *Config
	router           *gin.Engine
	handler          *handler.Handler
	store            *store.Store
	dadataSuggestAPI *suggest.Api
	s3               *s3.S3
}

func New(config *Config) *APIServer {
	store := store.New()
	dadataSuggestAPI := dadata.NewSuggestApi()

	session := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:   aws.String(config.S3Region),
			Endpoint: aws.String(config.S3Endpoint),
			Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
				AccessKeyID:     config.S3AccessKeyID,
				SecretAccessKey: config.S3SecretAccessKey,
			}),
		},
	}))

	s3 := s3.New(session)
	handlerEnv := &handler.HandlerEnv{
		JwtSecret: config.JwtSecret,
	}

	handler := handler.New(store, handlerEnv, dadataSuggestAPI, s3)

	return &APIServer{
		config:           config,
		router:           gin.Default(),
		handler:          handler,
		store:            store,
		dadataSuggestAPI: dadataSuggestAPI,
		s3:               s3,
	}
}

func (s *APIServer) Start() error {
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

	err = s.store.AutoMigrate()

	if err != nil {
		return err
	}

	return s.router.Run(":" + s.config.Port)
}

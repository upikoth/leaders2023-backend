package apiserver

import (
	"log"

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
}

func New(config *Config) *ApiServer {
	store := store.New()
	dadataSuggestApi := dadata.NewSuggestApi()
	handler := handler.New(store, config.JwtSecret, dadataSuggestApi)

	return &ApiServer{
		config:           config,
		router:           gin.Default(),
		handler:          handler,
		store:            store,
		dadataSuggestApi: dadataSuggestApi,
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

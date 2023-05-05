package apiserver

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/upikoth/leaders2023-backend/internal/app/handler"
	"github.com/upikoth/leaders2023-backend/internal/app/store"
)

type ApiServer struct {
	config  *Config
	router  *gin.Engine
	handler *handler.Handler
	store   *store.Store
}

func New(config *Config) *ApiServer {
	store := store.New()
	handler := handler.New(store, config.JwtSecret)

	return &ApiServer{
		config:  config,
		router:  gin.Default(),
		handler: handler,
		store:   store,
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

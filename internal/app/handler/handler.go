package handler

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ekomobile/dadata/v2/api/suggest"
	v1 "github.com/upikoth/leaders2023-backend/internal/app/handler/v1"
	"github.com/upikoth/leaders2023-backend/internal/app/store"
)

type Handler struct {
	V1 *v1.HandlerV1
}

type HandlerEnv struct {
	JwtSecret []byte `envconfig:"JWT_SECRET" required:"true"`
}

func New(store *store.Store, env *HandlerEnv, dadataSuggestAPI *suggest.Api, s3 *s3.S3) *Handler {
	handlerV1Env := &v1.HandlerV1Env{
		JwtSecret: env.JwtSecret,
	}

	return &Handler{
		V1: v1.New(store, handlerV1Env, dadataSuggestAPI, s3),
	}
}

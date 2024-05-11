package v1

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/upikoth/leaders2023-backend/internal/app/store"
)

type HandlerV1 struct {
	store            *store.Store
	env              *HandlerV1Env
	dadataSuggestAPI *suggest.Api
	s3               *s3.S3
}

type HandlerV1Env struct {
	JwtSecret []byte `envconfig:"JWT_SECRET" required:"true"`
}

func New(store *store.Store, env *HandlerV1Env, dadataSuggestAPI *suggest.Api, s3 *s3.S3) *HandlerV1 {
	return &HandlerV1{
		store:            store,
		env:              env,
		dadataSuggestAPI: dadataSuggestAPI,
		s3:               s3,
	}
}

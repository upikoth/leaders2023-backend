package v1

import (
	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/upikoth/leaders2023-backend/internal/app/store"
)

type HandlerV1 struct {
	store            *store.Store
	jwtSecret        []byte
	dadataSuggestApi *suggest.Api
}

func New(store *store.Store, jwtSecret []byte, dadataSuggestApi *suggest.Api) *HandlerV1 {
	return &HandlerV1{
		store:            store,
		jwtSecret:        jwtSecret,
		dadataSuggestApi: dadataSuggestApi,
	}
}

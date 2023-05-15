package handler

import (
	"github.com/ekomobile/dadata/v2/api/suggest"
	v1 "github.com/upikoth/leaders2023-backend/internal/app/handler/v1"
	"github.com/upikoth/leaders2023-backend/internal/app/store"
)

type Handler struct {
	V1 *v1.HandlerV1
}

func New(store *store.Store, jwtSecret []byte, dadataSuggestApi *suggest.Api) *Handler {
	return &Handler{
		V1: v1.New(store, jwtSecret, dadataSuggestApi),
	}
}

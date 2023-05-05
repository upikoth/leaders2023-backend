package v1

import "github.com/upikoth/leaders2023-backend/internal/app/store"

type HandlerV1 struct {
	store     *store.Store
	jwtSecret []byte
}

func New(store *store.Store, jwtSecret []byte) *HandlerV1 {
	return &HandlerV1{
		store:     store,
		jwtSecret: jwtSecret,
	}
}

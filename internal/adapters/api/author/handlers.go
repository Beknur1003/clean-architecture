package author

import (
	"github.com/Beknur1003/clean-architecture.git/internal/adapters/api"
	"github.com/julienschmidt/httprouter"
)

// handlers - controllers

type handler struct {
	authorService Service
}

func NewHandler(service Service) api.Handler {
	return &handler{authorService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	panic("implement me")
}

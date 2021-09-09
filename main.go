package main

import (
	"github.com/go-chi/chi/v5"
)

type RootData struct {
	Ff int `json:"ff" validate:"required,numeric,min=10,max=100"`
}

func main() {
	controller := NewController()
	SetActions(controller)

	_ = ListenPort(1111).ModifyRouter(func(router *chi.Mux) {
		router.Group(func(r chi.Router) {
			r.With(
				jsonInCtxStruct(func() interface{} { return new(RootData) }),
				RootRequest,
			).Post("/", controller.GetAction("POST."))
		})
		router.Get("/", controller.GetAction("."))
		router.Get("/home", controller.GetAction(".home"))
	}).Start()
}

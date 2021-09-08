package main

import (
	"github.com/go-chi/chi/v5"
)

func main() {
	controller := NewController()
	SetActions(controller)

	ListenPort(1111).ModifyRouter(func(router *chi.Mux) {
		router.Group(func(r chi.Router) {
			r.Use(jsonInContext)
			r.Post("/", controller.GetAction("POST."))
		})
		router.Get("/", controller.GetAction("."))
		router.Get("/home", controller.GetAction(".home"))
	}).Start()
}

package main

import "github.com/go-chi/chi/v5"

func GetRoutes(c *controller) func(router *chi.Mux) {
	return func(router *chi.Mux) {
		router.Group(func(r chi.Router) {
			r.With(
				jsonInCtxStruct(c.GetStructFunc("POST.")),
				RootRequest,
			).Post("/", c.GetAction("POST."))
		})

		router.Get("/", c.GetAction("."))
		router.Get("/home", c.GetAction(".home"))
	}
}

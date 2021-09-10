package main

import (
	"fmt"
	"net/http"
)

func SetActions(c *controller) {
	c.AddAction(".", func(writer http.ResponseWriter, request *http.Request) {
		JsonResponse(Response{
			Success: true,
			Data:    "Root page",
		}, writer)
	}).AddStructFunc(".", func() interface{} { return new(RootData) })

	c.AddAction("POST.", func(writer http.ResponseWriter, request *http.Request) {

		JsonResponse(Response{
			Success: true,
			Data:    request.Context().Value("json"),
		},
			writer)
	})
	c.AddAction(".home", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, ".home")
	})
}

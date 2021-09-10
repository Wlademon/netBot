package main

import (
	"fmt"
	"net/http"
)

func SetActions(c *controller) {
	c.SetAction(".", func(writer http.ResponseWriter, request *http.Request) {
		JsonResponse(Response{
			Success: true,
			Data:    "Root page",
		}, writer)
	}, func() interface{} { return new(RootData) })

	c.SetAction("POST.", func(writer http.ResponseWriter, request *http.Request) {
		JsonResponse(Response{
			Success: true,
			Data:    request.Context().Value("json"),
		},
			writer)
	}, func() interface{} { return new(RootData) })
	c.AddAction(".home", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, ".home")
	})
}

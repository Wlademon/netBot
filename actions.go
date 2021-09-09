package main

import (
	"fmt"
	"net/http"
)

func SetActions(c *controller) {
	c.AddAction(".", func(writer http.ResponseWriter, request *http.Request) {
		JsonResponse(Response{
			Success: true,
			Data: struct {
				Test string `json:"test"`
			}{Test: "1112"},
		}, writer)
	})
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

package main

import (
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type RootData struct {
	Ff int `json:"ff"`
}

func RootRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		data := request.Context().Value("json").(RootData)
		err := validator.New().Struct(data)
		if err != nil {
			return
		}

		handler.ServeHTTP(writer, request)
	})
}
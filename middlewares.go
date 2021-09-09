package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func jsonInCtxStruct(f func() interface{}) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			var obj = f()
			var jsonCtx = request.Context()
			err := json.NewDecoder(request.Body).Decode(&obj)
			if err == nil {
				jsonCtx = context.WithValue(request.Context(), "json", obj)
			} else {
				jsonCtx = context.WithValue(request.Context(), "json", nil)
			}

			request = request.WithContext(jsonCtx)

			handler.ServeHTTP(writer, request)
		})
	}
}

func jsonInContext(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var p interface{}
		err := json.NewDecoder(request.Body).Decode(&p)
		if err == nil {
			jsonCtx := context.WithValue(request.Context(), "json", p)
			request = request.WithContext(jsonCtx)
		}

		handler.ServeHTTP(writer, request)
	})
}

func formInContext(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		err := request.ParseForm()
		fmt.Print(request.Form)
		if err == nil {
			formCtx := context.WithValue(request.Context(), "form", request.Form)
			request = request.WithContext(formCtx)
		}

		handler.ServeHTTP(writer, request)
	})
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

func JsonResponse(data Response, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(data)
	if err != nil {
		m := fmt.Sprintf("{\"success\": false, \"error\": \"%s\"}", err)
		_, _ = fmt.Fprint(writer, m)
	}
	_, _ = fmt.Fprint(writer, string(bytes))
}

func ErrorResponse(writer http.ResponseWriter, response Response, code int) {
	writer.Header().Set("Content-Type", "application/json")
	bytes, _ := json.Marshal(response)
	writer.Header().Set("X-Content-Type-Options", "nosniff")
	writer.WriteHeader(code)
	_, _ = fmt.Fprintln(writer, string(bytes))
}

func ForbiddenResponse(writer http.ResponseWriter) {
	ExceptionResponse(writer, http.StatusForbidden)
}

func InvalidDataResponse(writer http.ResponseWriter, err error) {
	ErrorResponse(writer, Response{
		Success: false,
		Data:    nil,
		Error:   err.Error(),
	}, http.StatusUnprocessableEntity)
}

func ExceptionResponse(writer http.ResponseWriter, status int) {
	ErrorResponse(writer, Response{
		Success: false,
		Data:    nil,
		Error:   http.StatusText(status),
	}, status)
}

func SetErrorResponse() func(router *chi.Mux) {
	return func(router *chi.Mux) {
		router.NotFound(func(writer http.ResponseWriter, request *http.Request) {
			ExceptionResponse(writer, http.StatusNotFound)
		})
		router.MethodNotAllowed(func(writer http.ResponseWriter, request *http.Request) {
			ExceptionResponse(writer, http.StatusMethodNotAllowed)
		})
	}
}

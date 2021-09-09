package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
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
	http.Error(writer, string(bytes), code)
}

func ForbiddenResponse(writer http.ResponseWriter) {
	ErrorResponse(writer, Response{
		Success: false,
		Data:    nil,
		Error:   http.StatusText(http.StatusForbidden),
	}, http.StatusForbidden)
}

func InvalidDataResponse(writer http.ResponseWriter, err error) {
	ErrorResponse(writer, Response{
		Success: false,
		Data:    nil,
		Error:   err.Error(),
	}, http.StatusUnprocessableEntity)
}

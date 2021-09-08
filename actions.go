package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Success bool `json:"success"`
	Data interface{} `json:"data"`
	Context interface{} `json:"context"`
}

func SetActions (c *controller) {
	c.AddAction(".", func(writer http.ResponseWriter, request *http.Request) {
		jsonResponse(Response{
			Success: true,
			Data: struct {
				Test string `json:"test"`
			}{Test: "1112"},
		}, writer)
	})
	c.AddAction("POST.", func(writer http.ResponseWriter, request *http.Request) {

		jsonResponse(Response{
			Success: true,
			Data: request.Context().Value("json"),
			Context: request.Context().Value("form"),
		}, writer)
	})
	c.AddAction(".home", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, ".home")
	})
}

func jsonResponse(data Response, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(data)
	if err != nil {
		m := fmt.Sprintf("{\"success\": false, \"error\": \"%s\"}", err)
		_, _ = fmt.Fprint(writer, m)
	}
	_, _ = fmt.Fprint(writer, string(bytes))
}
package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type server struct {
	router    *chi.Mux
	serverUrl string
}

func (s *server) GetRouter() *chi.Mux {
	return s.router
}

func (s *server) Start() error {
	return http.ListenAndServe(s.serverUrl, s.router)
}

func (s *server) ModifyRouter(modifier func(router *chi.Mux)) *server {
	modifier(s.router)
	return s
}

func ListenServer(serverName string, port uint) *server {
	return &server{
		router:    chi.NewRouter(),
		serverUrl: serverName + ":" + strconv.Itoa(int(port)),
	}
}

func ListenPort(port uint) *server {
	return ListenServer("", port)
}

package server

import (
	"day04/ex01/candy-server/internal/handler"
	"log"
	"net/http"
)

type Server struct {
	httpServer http.Server
}

func NewServer() *Server {
	return &Server{
		httpServer: http.Server{
			Addr:    "127.0.0.1:8080",
			Handler: &handler.Handler{},
		},
	}
}

func (s *Server) Run() error {
	log.Println("server is running")
	return s.httpServer.ListenAndServeTLS("./tls-cfg/localhost/", "")
}

func (s *Server) Stop() error {
	return s.httpServer.Close()
}

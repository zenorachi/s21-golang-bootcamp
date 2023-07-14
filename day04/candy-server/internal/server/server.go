package server

import (
	"day04/candy-server/internal/handler"
	"log"
	"net/http"
)

type Server struct {
	httpServer http.Server
}

func NewServer() *Server {
	return &Server{
		httpServer: http.Server{
			Addr:    "localhost:8080",
			Handler: &handler.Handler{},
		},
	}
}

func (s *Server) Run() error {
	log.Println("start")
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.httpServer.Close()
}

package server

import (
	"day04/ex00/candy-server/internal/handler"
	"log"
	"net/http"
)

type Server struct {
	httpServer http.Server
}

// NewServer TODO: add flag to change port
func NewServer() *Server {
	return &Server{
		httpServer: http.Server{
			Addr:    "127.0.0.1:3333",
			Handler: &handler.Handler{},
		},
	}
}

func (s *Server) Run() error {
	log.Println("server is running")
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.httpServer.Close()
}

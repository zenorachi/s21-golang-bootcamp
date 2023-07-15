package server

import (
	"crypto/x509"
	"day04/ex02/candy-server/internal/handler"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	CertFile = "./cert/localhost/cert.pem"
	KeyFile  = "./cert/localhost/key.pem"
	MiniCA   = "./cert/minica.pem"
)

type Server struct {
	httpServer http.Server
}

func NewServer() *Server {
	data, err := ioutil.ReadFile(MiniCA)
	if err != nil {
		log.Fatalln(err)
	}

	cp, err := x509.SystemCertPool()
	if err != nil {
		log.Println(err)
	}

	cp.AppendCertsFromPEM(data)

	return &Server{
		httpServer: http.Server{
			Addr:    ":8080",
			Handler: &handler.Handler{},
		},
	}
}

func (s *Server) Run() error {
	log.Println("server is running")
	return s.httpServer.ListenAndServeTLS(CertFile, KeyFile)
}

func (s *Server) Stop() error {
	return s.httpServer.Close()
}

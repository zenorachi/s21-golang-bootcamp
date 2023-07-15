package server

import (
	"crypto/tls"
	"crypto/x509"
	"day04/ex01/candy-server/internal/handler"
	"day04/ex01/utils"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	CertFile = "./tls-cfg/localhost/cert.pem"
	KeyFile  = "./tls-cfg/localhost/key.pem"
	MiniCA   = "./tls-cfg/minica.pem"
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

	tlsCfg := &tls.Config{
		ClientCAs:      cp,
		ClientAuth:     tls.RequireAndVerifyClientCert,
		GetCertificate: utils.CertReqFunc(CertFile, KeyFile),
	}

	return &Server{
		httpServer: http.Server{
			Addr:      ":8080",
			TLSConfig: tlsCfg,
			Handler:   &handler.Handler{},
		},
	}
}

func (s *Server) Run() error {
	log.Println("server is running")
	return s.httpServer.ListenAndServeTLS("", "")
}

func (s *Server) Stop() error {
	return s.httpServer.Close()
}

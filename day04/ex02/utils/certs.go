package utils

import (
	"crypto/tls"
	"fmt"
	"log"
)

func getCert(certFile, keyfile string) (c tls.Certificate, err error) {
	if certFile != "" && keyfile != "" {
		c, err = tls.LoadX509KeyPair(certFile, keyfile)
		if err != nil {
			log.Fatalf("error loading key pair: %v\n", err)
		}
	} else {
		err = fmt.Errorf("no certificate")
	}
	return
}

func ClientCertReqFunc(certFile, keyfile string) func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
	c, err := getCert(certFile, keyfile)

	return func(certReq *tls.CertificateRequestInfo) (*tls.Certificate, error) {
		if err != nil || certFile == "" {
			log.Fatalln("no certificate", err)
		}
		return &c, nil
	}
}

func CertReqFunc(certFile, keyfile string) func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
	c, err := getCert(certFile, keyfile)

	return func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		if err != nil || certFile == "" {
			log.Fatalln("no certificate", err)
		}
		return &c, nil
	}
}

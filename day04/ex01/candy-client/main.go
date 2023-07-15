package main

import (
	"crypto/tls"
	"crypto/x509"
	"day04/ex01/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	CertFile = "./tls-cfg/client/cert.pem"
	KeyFile  = "./tls-cfg/client/key.pem"
)

func main() {
	client := getClient()
	requestBody := fmt.Sprintf(`{"money": %d, "candyType": "%s", "candyCount": %d}`, 10, "AA", 1)

	resp, err := client.Post("https://localhost:8080/buy_candy", "application/json", strings.NewReader(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s", respBody)
}

func getClient() *http.Client {
	data, err := os.ReadFile("./tls-cfg/minica.pem")
	if err != nil {
		log.Fatalln(err)
	}

	cp, err := x509.SystemCertPool()
	if err != nil {
		log.Fatalln(err)
	}

	cp.AppendCertsFromPEM(data)

	tlsCfg := &tls.Config{
		RootCAs:              cp,
		GetClientCertificate: utils.ClientCertReqFunc(CertFile, KeyFile),
	}

	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsCfg,
		},
	}
}

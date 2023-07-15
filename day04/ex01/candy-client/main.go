package main

import (
	"crypto/tls"
	"crypto/x509"
	"day04/ex01/utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	CertFile = "./tls-cfg/client/cert.pem"
	KeyFile  = "./tls-cfg/client/key.pem"
)

func main() {
	client := getClient()
	kek := strings.NewReader("`{\"money\": 1, \"candyType\": AA, \"candyCount\": 1}`")
	resp, err := client.Post("https://localhost:8080/buy_candy", "application/json", kek)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
}

func getClient() *http.Client {
	data, err := ioutil.ReadFile("./tls-cfg/minica.pem")
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

package main

import (
	"crypto/tls"
	"crypto/x509"
	"day04/ex01/utils"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	CertFile = "./cert/client/cert.pem"
	KeyFile  = "./cert/client/key.pem"
)

type RequestFlags struct {
	k string
	c int64
	m int64
}

var fl RequestFlags

func init() {
	flag.StringVar(&fl.k, "k", "", "candy type")
	flag.Int64Var(&fl.c, "c", 0, "candy count")
	flag.Int64Var(&fl.m, "m", 0, "money given to the machine")
}

func main() {
	flag.Parse()
	client := getClient()
	requestBody := fmt.Sprintf(`{"money": %d, "candyType": "%s", "candyCount": %d}`, fl.m, fl.k, fl.c)

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
	data, err := os.ReadFile("./cert/minica.pem")
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

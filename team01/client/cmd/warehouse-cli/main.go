package main

import (
	"flag"
	"fmt"
	"log"
	"warehouse/client/services"
)

type flags struct {
	h string
	p uint
}

var f flags

func init() {
	flag.StringVar(&f.h, "H", "", "hostname")
	flag.UintVar(&f.p, "P", 8080, "port")
}

func main() {
	addr := fmt.Sprintf("%s:%d", f.h, f.p)

	resp, err := services.GetHeartBeat(addr)
	if err == nil {
		log.Println("error getting heartbeat")
	}

	fmt.Println("Connected to a database of Warehouse 13 at, ", addr)
	fmt.Println("Known nodes:")
	for _, node := range resp.Nodes {
		fmt.Printf("%s:%s\n", node.Host, node.Port)
	}
}

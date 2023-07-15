package main

import (
	"day04/ex02/candy-server/internal/server"
	"log"
)

func main() {
	s := server.NewServer()
	if err := s.Run(); err != nil {
		log.Fatalln(err)
	}
}

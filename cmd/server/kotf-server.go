package main

import (
	"fmt"
	"github.com/KubeOperator/kotf/pkg/config"
	"log"
)

func main() {
	config.Init()
	//constant.Init()
	if err := prepareStart(); err != nil {
		log.Fatal(err)
	}
	host := "0.0.0.0"
	port := 8080
	address := fmt.Sprintf("%s:%d", host, port)
	lis, err := newTcpListener(address)
	if err != nil {
		log.Fatal(err)
	}
	server := newServer()
	log.Printf("kobe server lisen at: %s", address)
	if err := server.Serve(*lis); err != nil {
		log.Fatal(err)
	}
}

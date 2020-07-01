package main

import (
	"fmt"
	"github.com/kotf/pkg/config"
	"github.com/spf13/viper"
	"log"
)

func main() {
	config.Init()
	//constant.Init()
	if err := prepareStart(); err != nil {
		log.Fatal(err)
	}
	host := viper.GetString("server.host")
	port := viper.GetInt("server.port")
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

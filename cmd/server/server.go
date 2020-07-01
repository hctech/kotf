package main

import (
	"google.golang.org/grpc"
	"net"
)

func newTcpListener(address string) (*net.Listener, error) {
	s, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
func newServer() *grpc.Server {
	gs := grpc.NewServer()
	//kotf := server.NewKotf()
	//api.(gs, kobe)
	return gs
}

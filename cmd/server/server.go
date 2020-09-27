package main

import (
	"github.com/KubeOperator/kotf/api"
	"github.com/KubeOperator/kotf/pkg/server"
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
	var options = []grpc.ServerOption{
		grpc.MaxRecvMsgSize(100 * 1024 * 1024),
		grpc.MaxSendMsgSize(100 * 1024 * 1024),
	}
	gs := grpc.NewServer(options...)
	kotf := server.NewKotf()
	api.RegisterKotfApiServer(gs, kotf)
	return gs
}

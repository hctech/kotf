package client

import (
	"fmt"
	"google.golang.org/grpc"
)

func NewKotfClient(host string, port int) *KotfClient {
	return &KotfClient{
		host: host,
		port: port,
	}
}

type KotfClient struct {
	host string
	port int
}

func (k *KotfClient) createConnection() (*grpc.ClientConn, error) {
	address := fmt.Sprintf("%s:%d", k.host, k.port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

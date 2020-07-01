package client

import (
	"fmt"
	"github.com/kotf/api"
	"golang.org/x/net/context"
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

func (c *KotfClient) createConnection() (*grpc.ClientConn, error) {
	address := fmt.Sprintf("%s:%d", c.host, c.port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (c *KotfClient) Init(clusterName string, cloudType string, provider string, cloudRegion string, hosts string) (result *api.Result, err error) {
	conn, err := c.createConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := api.NewKotfApiClient(conn)
	req := api.TerraformInitRequest{
		ClusterName: clusterName,
		Provider:    provider,
		Type:        cloudType,
		CloudRegion: cloudRegion,
		Hosts:       hosts,
	}
	result, err = client.Init(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

package client

import (
	"fmt"
	"github.com/KubeOperator/kotf/api"
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
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(100*1024*1024)))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (c *KotfClient) Init(clusterName string, cloudType string, provider string, cloudRegion string, hosts string) (*api.Result, error) {
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
	result, err := client.Init(context.Background(), &req)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (c *KotfClient) Apply(clusterName string) (*api.Result, error) {
	conn, err := c.createConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := api.NewKotfApiClient(conn)
	req := api.TerraformApplyRequest{
		ClusterName: clusterName,
	}
	result, err := client.Apply(context.Background(), &req)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (c *KotfClient) Destroy(clusterName string) (*api.Result, error) {
	conn, err := c.createConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := api.NewKotfApiClient(conn)
	req := api.TerraformDestroyRequest{
		ClusterName: clusterName,
	}
	result, err := client.Destroy(context.Background(), &req)
	if err != nil {
		return result, err
	}
	return result, nil
}

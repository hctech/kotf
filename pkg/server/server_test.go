package server

import (
	"context"
	"fmt"
	api "github.com/KubeOperator/kotf/api"
	"testing"
)

func TestServer(t *testing.T) {
	provider := `{
    "name": "vSphere",
    "username": "administrator@vsphere.local",
    "password": "",
    "host": ""
  }`
	cloudRegion := `{
    "datacenter": "Datacenter",
    "zones": [
      {
		"cluster": "vSAN-Cluster",
        "key": "z8928",
        "name": "Resources",
        "network": "VM Network",
        "datastore": "vsanDatastore",
        "imageName": "Centos7.6-template",
        "guestId": "centos7_64Guest"
      }
    ]
  }`
	hosts := `[
    {
      "shortName": "worker1",
      "name": "worker1.clsuter2.fit2cloud.com",
      "cpu": 1,
      "memory": 2048,
      "domain": "clsuter2.fit2cloud.com",
      "ip": "172.16.10.222",
      "zone": {
        "key": "z8928",
        "name": "Resources",
        "network": "VM Network",
        "datastore": "vsanDatastore",
        "imageName": "Centos7.6-template",
        "guestId": "centos7_64Guest",
      	"netMask": 24,
      	"gateway": "172.16.10.254"
      }
    }
  ]`

	req := api.TerraformInitRequest{
		ClusterName: "cluster2",
		Type:        "vSphere",
		CloudRegion: cloudRegion,
		Hosts:       hosts,
		Provider:    provider,
	}
	result, err := NewKotf().Init(context.Background(), &req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	req2 := api.TerraformApplyRequest{
		ClusterName: "cluster2",
		Type:        "vSphere",
	}
	result2, err := NewKotf().Apply(context.Background(), &req2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result2)
}

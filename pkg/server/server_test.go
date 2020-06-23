package server

import (
	"fmt"
	"github.com/kotf/api"
	"golang.org/x/net/context"
	"testing"
)

func TestKotfInit(t *testing.T) {
	vars := make(map[string]string)
	req := api.TerraformInitRequest{
		ClusterName: "test",
		Type:        "vSphere",
		Vars:        vars,
	}
	ctx := context.Background()
	result, err := NewKotf().Init(ctx, &req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

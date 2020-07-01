package server

import (
	"encoding/json"
	"github.com/KubeOperator/kotf/api"
	"github.com/KubeOperator/kotf/pkg/terraform"
	"golang.org/x/net/context"
)

type kotf struct {
}

func NewKotf() *kotf {
	return &kotf{}
}

func (k kotf) Init(ctx context.Context, req *api.TerraformInitRequest) (*api.Result, error) {
	t := terraform.NewTerraform()

	resp := &api.Result{
		Success: false,
	}

	var provider map[string]interface{}
	if err := json.Unmarshal([]byte(req.Provider), &provider); err == nil {
		return resp, nil
	}
	var cloudRegion map[string]interface{}
	if err := json.Unmarshal([]byte(req.CloudRegion), &cloudRegion); err == nil {
		return resp, nil
	}
	var hosts map[string]interface{}
	if err := json.Unmarshal([]byte(req.Hosts), &hosts); err == nil {
		return resp, nil
	}
	vars := map[string]interface{}{
		"provider":    provider,
		"cloudRegion": cloudRegion,
		"hosts":       hosts,
	}

	output, err := t.Init(req.ClusterName, req.Type, vars)
	if err != nil {
		return nil, err
	}
	resp.Output = output
	resp.Success = true
	return resp, nil
}

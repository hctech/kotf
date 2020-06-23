package server

import (
	"github.com/kotf/api"
	"golang.org/x/net/context"
)

type kotf struct {
}

func NewKotf() *kotf {
	return &kotf{}
}

func (k kotf) Init(ctx context.Context, req *api.TerraformInitRequest) (*api.Result, error) {
	t := NewTerraform()
	err := t.init(req.ClusterName, req.Type, req.Vars)
	if err != nil {
		return nil, err
	}
	resp := &api.Result{
		Success: true,
	}
	return resp, nil
}

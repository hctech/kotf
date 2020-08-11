package server

import (
	"encoding/json"
	"errors"
	"github.com/KubeOperator/kotf/api"
	"github.com/KubeOperator/kotf/pkg/terraform"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"golang.org/x/net/context"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Kotf struct {
}

func NewKotf() *Kotf {
	return &Kotf{}
}

func (k Kotf) Init(ctx context.Context, req *api.TerraformInitRequest) (*api.Result, error) {

	t := terraform.NewTerraform()
	resp := &api.Result{
		Success: false,
	}
	var provider map[string]interface{}
	if err := json.Unmarshal([]byte(req.Provider), &provider); err != nil {
		return resp, err
	}
	var cloudRegion map[string]interface{}
	if err := json.Unmarshal([]byte(req.CloudRegion), &cloudRegion); err != nil {
		return resp, err
	}
	var hosts []interface{}
	if err := json.Unmarshal([]byte(req.Hosts), &hosts); err != nil {
		return resp, err
	}
	vars := map[string]interface{}{
		"provider":    provider,
		"cloudRegion": cloudRegion,
		"hosts":       hosts,
	}

	output, err := t.Init(req.ClusterName, req.Type, vars)
	if err != nil {
		return resp, errors.New(output)
	}
	resp.Output = output
	resp.Success = true
	return resp, nil
}

func (k Kotf) Apply(ctx context.Context, req *api.TerraformApplyRequest) (*api.Result, error) {
	t := terraform.NewTerraform()
	resp := &api.Result{
		Success: false,
	}
	output, err := t.Apply(req.ClusterName)
	if err != nil {
		resp.Output = output
		return resp, errors.New(output)
	}
	resp.Output = output
	resp.Success = true
	return resp, nil
}

func (k Kotf) Destroy(ctx context.Context, req *api.TerraformDestroyRequest) (*api.Result, error) {
	t := terraform.NewTerraform()
	resp := &api.Result{
		Success: false,
	}
	output, err := t.Destroy(req.ClusterName)
	if err != nil {
		resp.Output = output
		detail, _ := ptypes.MarshalAny(resp)
		s := spb.Status{
			Code:    int32(codes.FailedPrecondition),
			Message: err.Error(),
			Details: []*any.Any{detail},
		}
		return resp, status.ErrorProto(&s)
	}
	resp.Output = output
	resp.Success = true
	return resp, nil
}

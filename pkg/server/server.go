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

func (k Kotf) Init(ctx context.Context, req *api.TerraformInitRequest) (*api.KotfResult, error) {

	t := terraform.NewTerraform()
	resp := &api.KotfResult{
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

func (k Kotf) Apply(ctx context.Context, req *api.TerraformApplyRequest) (*api.KotfResult, error) {
	t := terraform.NewTerraform()
	resp := &api.KotfResult{
		Success: false,
	}
	var cloudRegion map[string]interface{}
	if err := json.Unmarshal([]byte(req.CloudRegion), &cloudRegion); err != nil {
		return resp, err
	}
	var vars = make([]string, 0)
	for k, v := range cloudRegion {
		if !isVars(k) {
			continue
		}
		if value, ok := v.(string); ok {
			if value != "" {
				vars = append(vars, "-var")
				vars = append(vars, k+"="+value)
			}
		}
	}
	output, err := t.Apply(req.ClusterName, vars)
	if err != nil {
		resp.Output = output
		return resp, errors.New(output)
	}
	resp.Output = output
	resp.Success = true
	return resp, nil
}

func (k Kotf) Destroy(ctx context.Context, req *api.TerraformDestroyRequest) (*api.KotfResult, error) {
	t := terraform.NewTerraform()
	resp := &api.KotfResult{
		Success: false,
	}
	var cloudRegion map[string]interface{}
	if err := json.Unmarshal([]byte(req.CloudRegion), &cloudRegion); err != nil {
		return resp, err
	}
	var vars = make([]string, 0)
	for k, v := range cloudRegion {
		if !isVars(k) {
			continue
		}
		if value, ok := v.(string); ok {
			if value != "" {
				vars = append(vars, "-var")
				vars = append(vars, k+"="+value)
			}
		}
	}
	output, err := t.Destroy(req.ClusterName, vars)
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

func isVars(key string) bool {
	keys := [3]string{"username", "user", "password"}
	for _, v := range keys {
		if v == key {
			return true
		}
	}
	return false
}

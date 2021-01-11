package terraform

import (
	"bytes"
	"errors"
	"github.com/KubeOperator/kotf/pkg/constant"
	"github.com/KubeOperator/kotf/util"
	"os"
	"os/exec"
	"path"
)

type Terraform struct {
}

func NewTerraform() *Terraform {
	return &Terraform{}
}

func (t *Terraform) Init(cluster string, cloud string, vars map[string]interface{}) (string, error) {

	dir := path.Join(constant.ProjectDir, cluster)
	exist, err := util.PathExists(dir)
	if err != nil {
		return "", err
	}
	if !exist {
		err = util.CreatePath(dir)
		if err != nil {
			return "", err
		}
	}
	filePath := ""
	if cloud == constant.OpenStack {
		filePath = constant.OpenStackFilePath
	} else if cloud == constant.VSphere {
		filePath = constant.VSphereFilePath
	} else if cloud == constant.FusionCompute {
		filePath = constant.FusionComputeFilePath
	} else {
		return "", errors.New("not support")
	}
	err = util.MoveFileToPath(filePath, constant.TerraformFile, dir)
	if err != nil {
		return "", err
	}
	templateFilePath := path.Join(dir, constant.TerraformFile)
	err = util.CoverFileVars(templateFilePath, vars, path.Join(dir, constant.MainFile))
	if err != nil {
		return "", err
	}
	err = os.Remove(templateFilePath)
	if err != nil {
		return "", err
	}

	result, err := ExecCommand(dir, constant.TerraformCommand, constant.TerraformInit, constant.TerraformNoColor)
	if err != nil {
		return result, err
	}
	return result, err
}

func (t *Terraform) Apply(cluster string) (string, error) {

	dir := path.Join(constant.ProjectDir, cluster)
	result, err := ExecCommand(dir, constant.TerraformCommand, constant.TerraformApply, constant.TerraformApplyApprove, constant.TerraformNoColor)
	if err != nil {
		return result, err
	}
	return result, err
}

func (t *Terraform) Destroy(cluster string) (string, error) {
	dir := path.Join(constant.ProjectDir, cluster)
	exist, err := util.PathExists(dir)
	if err != nil {
		return "", err
	}
	if !exist {
		err = util.CreatePath(dir)
		if err != nil {
			return "", err
		}
	}
	result, err := ExecCommand(dir, constant.TerraformCommand, constant.TerraformDestroy, constant.TerraformApplyApprove, constant.TerraformNoColor)
	if err != nil {
		return result, err
	}
	return result, err
}

func ExecCommand(path string, name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	cmd.Dir = path
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	cmd.Stderr = cmd.Stdout
	if err = cmd.Start(); err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	for {
		out := make([]byte, 1024)
		_, err := stdout.Read(out)
		buffer.Write(out)
		if err != nil {
			break
		}
	}

	if err = cmd.Wait(); err != nil {
		return string(buffer.Bytes()), err
	}
	return string(buffer.Bytes()), nil
}

package terraform

import (
	"errors"
	"github.com/KubeOperator/kotf/pkg/constant"
	"github.com/KubeOperator/kotf/util"
	"io/ioutil"
	"os/exec"
	"path"
)

type Terraform struct {
}

func NewTerraform() *Terraform {
	return &Terraform{}
}

func (t *Terraform) Init(cluster string, cloud string, vars map[string]interface{}) (string, error) {

	dir := path.Join(constant.TerraformDir, cluster)
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
		filePath = "/Users/zk.wang/go/src/github.com/kotf/resource/vsphere"
	} else {
		return "", errors.New("not support")
	}
	err = util.MoveFileToPath(filePath, constant.TerraformFile, dir)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("terraform", "init")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		return "", err
	}
	if opBytes, err := ioutil.ReadAll(stdout); err != nil {
		return "", err
	} else {
		return string(opBytes), err
	}
	return "", err
}

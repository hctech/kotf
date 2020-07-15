package constant

import "path"

const (
	OpenStack = "OpenStack"
	VSphere   = "vSphere"
)

var (
	BaseDir               = "/var/kotf"
	TerraformFile         = "terraform.tf"
	MainFile              = "main.tf"
	DataDir               = path.Join(BaseDir, "data")
	WorkDir               = path.Join(BaseDir, "work")
	ProjectDir            = path.Join(DataDir, "project")
	ResourceDir           = path.Join(DataDir, "resource")
	OpenStackFilePath     = path.Join(ResourceDir, "openstack")
	VSphereFilePath       = path.Join(ResourceDir, "vsphere")
	TerraformCommand      = "terraform"
	TerraformInit         = "init"
	TerraformApply        = "apply"
	TerraformDestroy      = "destroy"
	TerraformApplyApprove = "-auto-approve"
	TerraformNoColor      = "-no-color"
)

package constant

import "path"

const (
	OpenStack     = "OpenStack"
	VSphere       = "vSphere"
	FusionCompute = "FusionCompute"
)

var (
	BaseDir               = "/var/kotf"
	TerraformFile         = "terraform.tf"
	MainFile              = "main.tf"
	DataDir               = path.Join(BaseDir, "data")
	ProjectDir            = path.Join(DataDir, "project")
	ResourceDir           = path.Join(BaseDir, "resource")
	OpenStackFilePath     = path.Join(ResourceDir, "openstack")
	VSphereFilePath       = path.Join(ResourceDir, "vsphere")
	FusionComputeFilePath = path.Join(ResourceDir, "fusioncompute")
	TerraformCommand      = "terraform"
	TerraformInit         = "init"
	TerraformApply        = "apply"
	TerraformDestroy      = "destroy"
	TerraformApplyApprove = "-auto-approve"
	TerraformNoColor      = "-no-color"
)

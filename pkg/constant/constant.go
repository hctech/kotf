package constant

import "path"

const (
	OpenStack = "openStack"
	VSphere   = "vSphere"
)

var (
	BaseDir           = "/var/kotf"
	ResourceDir       = "resource"
	TerraformFile     = "terraform.tf"
	DataDir           = path.Join(BaseDir, "data")
	TerraformDir      = path.Join(BaseDir, "terraform")
	WorkDir           = path.Join(BaseDir, "work")
	OpenStackFilePath = path.Join(ResourceDir, "openstack")
	VSphereFilePath   = path.Join(ResourceDir, "vsphere")
	ProjectDir        = path.Join(DataDir, "project")
)

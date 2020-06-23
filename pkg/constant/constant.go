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
	TerraformDir      = path.Join(BaseDir, "terraform")
	OpenStackFilePath = path.Join(ResourceDir, "openstack")
	VSphereFilePath   = path.Join(ResourceDir, "vsphere")
)

provider "vsphere" {
  user = "test"
  password = "123123"
  vsphere_server = "123123"

  # If you have a self-signed cert
  allow_unverified_ssl = true
}

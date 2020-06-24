provider "vsphere" {
  user = "{{ .Username }}"
  password = "{{ .Password }}"
  vsphere_server = "{{ .Host }}"

  # If you have a self-signed cert
  allow_unverified_ssl = true
}

{{ $provider := .provider }}
{{ $region := .cloudRegion }}
{{ $hosts := .hosts }}

provider "vsphere" {
  user = "{{ $provider.username }}"
  password = "{{ $provider.password }}"
  vsphere_server = "{{ $provider.host }}"

  # If you have a self-signed cert
  allow_unverified_ssl = true
}

data "vsphere_datacenter" "dc" {
  name = "{{ $region.datacenter }}"
}

{{ range $region.zones}}
data "vsphere_resource_pool" "{{ .key }}" {
  {{ if  eq .name "Resources" }}
   name  = "{{ .cluster }}/Resources"
  {{ else if ne .name "Resources" }}
   name  = "{{ .cluster }}/Resources/{{ .name }}"
  {{ end }}
   datacenter_id = data.vsphere_datacenter.dc.id
}

data "vsphere_network" "{{ .key }}" {
  name = "{{ .network }}"
  datacenter_id = data.vsphere_datacenter.dc.id
}

data "vsphere_datastore" "{{ .key }}" {
  name = "{{ .datastore }}"
  datacenter_id = data.vsphere_datacenter.dc.id
}

data "vsphere_virtual_machine" "{{ .key }}" {
  name = "{{ .imageName }}"
  datacenter_id = data.vsphere_datacenter.dc.id
}
{{ end }}


{{ range $hosts}}
resource "vsphere_virtual_machine" "{{.shortName}}" {
  name = "{{ .name }}"
  folder = "kubeoperator"
  resource_pool_id = data.vsphere_resource_pool.{{ .zone.key }}.id
  datastore_id = data.vsphere_datastore.{{ .zone.key }}.id
  num_cpus = {{ .cpu }}
  memory = {{ .memory }}
  guest_id = "{{ .zone.guestId }}"

  network_interface {
    network_id = data.vsphere_network.{{ .zone.key }}.id
  }

  disk {
    label            = "disk0"
    size             = data.vsphere_virtual_machine.{{ .zone.key }}.disks.0.size
    eagerly_scrub    = data.vsphere_virtual_machine.{{ .zone.key }}.disks.0.eagerly_scrub
    thin_provisioned = data.vsphere_virtual_machine.{{ .zone.key }}.disks.0.thin_provisioned
  }

  lifecycle {
    ignore_changes = [
      disk,
    ]
  }

  clone {
    template_uuid = data.vsphere_virtual_machine.{{ .zone.key }}.id
    timeout = 60
    customize {

      linux_options {
        host_name = "{{ .shortName }}"
        domain = "{{ .domain }}"
      }

      network_interface {
        ipv4_address = "{{ .ip }}"
        ipv4_netmask = "{{ .zone.netMask }}"
      }
      ipv4_gateway = "{{ .zone.gateway}}"
    }
  }
}
{{ end }}
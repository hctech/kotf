{{ $provider := .provider }}
{{ $region := .cloudRegion }}
{{ $hosts := .hosts }}

provider "fusioncompute" {
  user =  "{{ $provider.user }}"
  password = "{{ $provider.password }}"
  server = "{{ $provider.server }}"
}

data "fusioncompute_site" "site" {
  name = "{{ $region.datacenter }}"
}

data "fusioncompute_cluster" "cluster" {
  name = "ManagementCluster"
  site_uri = data.fusioncompute_site.site.id
}

{{ range $region.zones}}
data "fusioncompute_cluster" "{{ .key }}" {
  name  = "{{ .cluster }}"
  site_uri = data.fusioncompute_site.site.id
}

data "fusioncompute_dvswitch" "{{ .key }}" {
  name = "{{ .switch }}"
  site_uri = data.fusioncompute_site.site.id
}
data "fusioncompute_portgroup" "{{ .key }}" {
  name = "{{ .portgroup }}"
  dvswitch_uri = data.fusioncompute_dvswitch.{{ .key }}.id
  site_uri = data.fusioncompute_site.site.id
}

data "fusioncompute_datastore" "{{ .key }}" {
  name = "{{ .datastore }}"
  site_uri = data.fusioncompute_site.site.id
}

data "fusioncompute_vm" "{{ .key }}" {
  name = "{{ .template }}"
  site_uri = data.fusioncompute_site.site.id
}
{{ end }}

{{ range $hosts}}
resource "fusioncompute_vm" "{{.shortName}}" {
  name = "{{ .name }}"
  timeout = 30
  num_cpus = {{ .cpu }}
  memory = {{ .memory }}
  site_uri = data.fusioncompute_site.site.id
  cluster_uri = data.fusioncompute_cluster.{{ .zone.key }}.id
  datastore_uri = data.fusioncompute_datastore.{{ .zone.key }}.id
  template_uri = data.fusioncompute_vm.{{ .zone.key }}.id
  network_interface {
    portgroup_uri = data.fusioncompute_portgroup.{{ .zone.key }}.id
  }
  disk {
    size = 50
    thin = true
  }
  customize {
    host_name = "{{ .shortName }}"
    network_interface {
      ipv4_address = "{{ .ip }}"
      ipv4_netmask = "{{ .zone.netMask }}"
    }
    ipv4_gateway = "{{ .zone.gateway}}"
    set_dns = "{{ .zone.dns1 }}"
    add_dns = "{{ .zone.dns2 }}"
  }
}
{{ end }}
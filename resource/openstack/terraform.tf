{{ $provider := .provider }}
{{ $region := .cloudRegion }}
{{ $hosts := .hosts }}



provider "openstack" {
  user_name   = "{{ $provider.username }}"
  tenant_id   = "{{ $provider.projectId }}"
  password    = "{{ $provider.password }}"
  auth_url    = "{{ $provider.identity }}"
  region      = "{{ $region.datacenter }}"
  user_domain_name = "{{ $provider.domainName }}"
}



{{ range $hosts}}

    {{ if  eq .zone.ipType "private" }}
        resource "openstack_networking_port_v2" "{{.shortName}}" {
          name               = "{{.shortName}}"
          network_id         = "{{.zone.network}}"
          admin_state_up     = "true"
          fixed_ip {
            subnet_id  = "{{.zone.subnet}}"
            ip_address = "{{.ip}}"
          }
        }

        resource "openstack_compute_instance_v2" "{{.shortName}}" {
          name            = "{{.name}}"
          image_name      = "{{.zone.imageName}}"
          flavor_name     = "{{.model}}"

          admin_pass = "KubeOperator@2019"
          user_data = "#!/bin/sh \necho 'root:KubeOperator@2019' | chpasswd"
          security_groups = ["{{.zone.securityGroup}}"]
          network {
            port = "${openstack_networking_port_v2.{{.shortName}}.id}"
          }
        }
    {{ else if eq .zone.ipType "floating" }}
          resource "openstack_networking_port_v2" "{{.shortName}}" {
             name            = "{{.shortName}}"
             admin_state_up  = "true"
             network_id      = "{{.zone.network}}"
           }
            resource "openstack_compute_instance_v2" "{{.shortName}}" {
              name            = "{{.name}}"
              image_name      = "{{.zone.imageName}}"
              flavor_name     = "{{.model}}"
              security_groups = ["{{.zone.securityGroup}}"]
              network {
                port = "${openstack_networking_port_v2.{{.shortName}}.id}"
              }
            }

            resource "openstack_networking_floatingip_v2" "{{.shortName}}" {
              pool = "{{.zone.floatingNetwork}}"
              address = "{{.ip}}"
            }

            resource "openstack_compute_floatingip_associate_v2" "{{.shortName}}" {
              floating_ip = "${openstack_networking_floatingip_v2.{{.shortName}}.address}"
              instance_id = "${openstack_compute_instance_v2.{{.shortName}}.id}"
              fixed_ip    = "${openstack_compute_instance_v2.{{.shortName}}.network.0.fixed_ip_v4}"
            }
    {{ end }}

{{ end }}

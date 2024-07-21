// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceCcFabricL3HandoffIPTransit(t *testing.T) {
	if os.Getenv("SDA") == "" {
		t.Skip("skipping test, set environment variable SDA")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_l3_handoff_ip_transit.test", "network_device_id", "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_l3_handoff_ip_transit.test", "interface_name", "TenGigabitEthernet1/0/2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_l3_handoff_ip_transit.test", "virtual_network_name", "SDA_VN1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_l3_handoff_ip_transit.test", "vlan_id", "205"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_l3_handoff_ip_transit.test", "tcp_mss_adjustment", "1400"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_l3_handoff_ip_transit.test", "local_ip_address", "10.0.0.1/24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_l3_handoff_ip_transit.test", "remote_ip_address", "10.0.0.2/24"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcFabricL3HandoffIPTransitPrerequisitesConfig + testAccDataSourceCcFabricL3HandoffIPTransitConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcFabricL3HandoffIPTransitPrerequisitesConfig = `
resource "catalystcenter_area" "test" {
  name        = "Area1"
  parent_name = "Global"
}
resource "catalystcenter_fabric_site" "test" {
  site_id                     = catalystcenter_area.test.id
  pub_sub_enabled             = false
  authentication_profile_name = "No Authentication"
  depends_on = [catalystcenter_area.test]
}
resource "catalystcenter_transit_peer_network" "test" {
  transit_peer_network_name = "TRANSIT_1"
  transit_peer_network_type = "ip_transit"
  routing_protocol_name     = "BGP"
  autonomous_system_number  = "65010"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcFabricL3HandoffIPTransitConfig() string {
	config := `resource "catalystcenter_fabric_l3_handoff_ip_transit" "test" {` + "\n"
	config += `	network_device_id = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"` + "\n"
	config += `	fabric_id = catalystcenter_fabric_site.test.id` + "\n"
	config += `	transit_network_id = catalystcenter_transit_peer_network.test.id` + "\n"
	config += `	interface_name = "TenGigabitEthernet1/0/2"` + "\n"
	config += `	virtual_network_name = "SDA_VN1"` + "\n"
	config += `	vlan_id = 205` + "\n"
	config += `	tcp_mss_adjustment = 1400` + "\n"
	config += `	local_ip_address = "10.0.0.1/24"` + "\n"
	config += `	remote_ip_address = "10.0.0.2/24"` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_fabric_l3_handoff_ip_transit" "test" {
			id = catalystcenter_fabric_l3_handoff_ip_transit.test.id
			network_device_id = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"
			fabric_id = catalystcenter_fabric_site.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig

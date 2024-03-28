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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceCcWirelessEnterpriseSSID(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_enterprise_ssid.test", "name", "mySSID1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_enterprise_ssid.test", "security_level", "wpa3_enterprise"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_enterprise_ssid.test", "enable_fast_lane", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_enterprise_ssid.test", "enable_mac_filtering", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_enterprise_ssid.test", "traffic_type", "data"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_enterprise_ssid.test", "enable_broadcast_ssid", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_enterprise_ssid.test", "fast_transition", "Adaptive"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_enterprise_ssid.test", "enable_session_time_out", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_enterprise_ssid.test", "session_time_out", "1800"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_enterprise_ssid.test", "enable_client_exclusion", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_enterprise_ssid.test", "client_exclusion_timeout", "180"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_enterprise_ssid.test", "enable_basic_service_set_max_idle", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_enterprise_ssid.test", "basic_service_set_client_idle_timeout", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_enterprise_ssid.test", "enable_directed_multicast_service", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_enterprise_ssid.test", "enable_neighbor_list", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_enterprise_ssid.test", "mfp_client_protection", "data"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_enterprise_ssid.test", "protected_management_frame", "Disabled"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcWirelessEnterpriseSSIDConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcWirelessEnterpriseSSIDConfig() string {
	config := `resource "catalystcenter_wireless_enterprise_ssid" "test" {` + "\n"
	config += `	name = "mySSID1"` + "\n"
	config += `	security_level = "wpa3_enterprise"` + "\n"
	config += `	passphrase = "Cisco123"` + "\n"
	config += `	enable_fast_lane = false` + "\n"
	config += `	enable_mac_filtering = false` + "\n"
	config += `	traffic_type = "data"` + "\n"
	config += `	radio_policy = "Triple band operation(2.4GHz, 5GHz and 6GHz)"` + "\n"
	config += `	enable_broadcast_ssid = true` + "\n"
	config += `	fast_transition = "Adaptive"` + "\n"
	config += `	enable_session_time_out = true` + "\n"
	config += `	session_time_out = 1800` + "\n"
	config += `	enable_client_exclusion = true` + "\n"
	config += `	client_exclusion_timeout = 180` + "\n"
	config += `	enable_basic_service_set_max_idle = true` + "\n"
	config += `	basic_service_set_client_idle_timeout = 300` + "\n"
	config += `	enable_directed_multicast_service = true` + "\n"
	config += `	enable_neighbor_list = true` + "\n"
	config += `	mfp_client_protection = "data"` + "\n"
	config += `	protected_management_frame = "Disabled"` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_wireless_enterprise_ssid" "test" {
			id = catalystcenter_wireless_enterprise_ssid.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig

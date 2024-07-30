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
func TestAccDataSourceCcFabricProvisionDevice(t *testing.T) {
	if os.Getenv("SDA") == "" {
		t.Skip("skipping test, set environment variable SDA")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_provision_device.test", "network_device_id", "4cb565d3-1944-42be-be9f-a87cff79e831"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcFabricProvisionDevicePrerequisitesConfig + testAccDataSourceCcFabricProvisionDeviceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcFabricProvisionDevicePrerequisitesConfig = `
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
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcFabricProvisionDeviceConfig() string {
	config := `resource "catalystcenter_fabric_provision_device" "test" {` + "\n"
	config += `	site_id = catalystcenter_area.test.id` + "\n"
	config += `	network_device_id = "4cb565d3-1944-42be-be9f-a87cff79e831"` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_fabric_provision_device" "test" {
			id = catalystcenter_fabric_provision_device.test.id
			site_id = catalystcenter_area.test.id
			network_device_id = "4cb565d3-1944-42be-be9f-a87cff79e831"
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig

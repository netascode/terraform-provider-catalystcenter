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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccCcImageDistribution(t *testing.T) {
	if os.Getenv("IMAGE_DISTRIBUTION") == "" {
		t.Skip("skipping test, set environment variable IMAGE_DISTRIBUTION")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_image_distribution.test", "device_uuid", "138b3181-f9c5-4271-9292-cf3152ab4d3e"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_image_distribution.test", "image_uuid", "faa9c5f7-d093-459a-8164-cc555bbf3b80"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcImageDistributionConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcImageDistributionConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcImageDistributionConfig_minimum() string {
	config := `resource "catalystcenter_image_distribution" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcImageDistributionConfig_all() string {
	config := `resource "catalystcenter_image_distribution" "test" {` + "\n"
	config += `	device_uuid = "138b3181-f9c5-4271-9292-cf3152ab4d3e"` + "\n"
	config += `	image_uuid = "faa9c5f7-d093-459a-8164-cc555bbf3b80"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

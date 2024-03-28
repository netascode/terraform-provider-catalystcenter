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
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &DeviceDetailDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceDetailDataSource{}
)

func NewDeviceDetailDataSource() datasource.DataSource {
	return &DeviceDetailDataSource{}
}

type DeviceDetailDataSource struct {
	client *cc.Client
}

func (d *DeviceDetailDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_detail"
}

func (d *DeviceDetailDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source fetches device's details, including the assigned physical location.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"policy_tag_name": schema.StringAttribute{
				MarkdownDescription: "Policy tag name.",
				Computed:            true,
			},
			"nw_device_role": schema.StringAttribute{
				MarkdownDescription: "Network device role, such as CORE.",
				Computed:            true,
			},
			"serial_number": schema.StringAttribute{
				MarkdownDescription: "Serial number of the device.",
				Computed:            true,
			},
			"nw_device_name": schema.StringAttribute{
				MarkdownDescription: "Name of the device.",
				Computed:            true,
			},
			"device_group_hierarchy_id": schema.StringAttribute{
				MarkdownDescription: "Identifier of the group hierarchy where the device is assigned, such as \"/360b1804-969f-4eab-a4ba-9832ea3f1731/26d4bbe6-f908-4b83-86f1-49bfbb1d05f6/\".",
				Computed:            true,
			},
			"cpu": schema.StringAttribute{
				MarkdownDescription: "CPU of the device, such as \"1.25\".",
				Computed:            true,
			},
			"nw_device_id": schema.StringAttribute{
				MarkdownDescription: "As of now, same as `id`.",
				Computed:            true,
			},
			"site_hierarchy_graph_id": schema.StringAttribute{
				MarkdownDescription: "Identifier of the site hierarchy graph where the device is assigned, such as \"/2b0a78ee-482e-4b4d-ae85-df289873cbbb/76b11b6a-d94a-431b-8bab-fd16b09f5d40/adad4a8a-17ec-4be3-a4b5-b6549b840afe/e8ce8788-9b13-46ec-86c8-740f7ea443c1/\".",
				Computed:            true,
			},
			"nw_device_family": schema.StringAttribute{
				MarkdownDescription: "Family of the network device, such as \"Switches and Hubs\".",
				Computed:            true,
			},
			"mac_address": schema.StringAttribute{
				MarkdownDescription: "MAC address.",
				Computed:            true,
			},
			"device_series": schema.StringAttribute{
				MarkdownDescription: "Device series, such as \"Cisco Catalyst 9400 Series Switches\".",
				Computed:            true,
			},
			"collection_status": schema.StringAttribute{
				MarkdownDescription: "Collection status in Catalyst Center, such as \"SUCCESS\".",
				Computed:            true,
			},
			"maintenance_mode": schema.BoolAttribute{
				MarkdownDescription: "Whether the device is in maintenance mode.",
				Computed:            true,
			},
			"software_version": schema.StringAttribute{
				MarkdownDescription: "Software version on the device.",
				Computed:            true,
			},
			"tag_id_list": schema.SetAttribute{
				MarkdownDescription: "Tags assigned to the device.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"overall_health": schema.Int64Attribute{
				MarkdownDescription: "Overall numerical health score.",
				Computed:            true,
			},
			"management_ip_address": schema.StringAttribute{
				MarkdownDescription: "IP address for managing the device.",
				Computed:            true,
			},
			"memory": schema.StringAttribute{
				MarkdownDescription: "Memory in gigabytes, e.g. \"39.186811767835785\".",
				Computed:            true,
			},
			"communication_state": schema.StringAttribute{
				MarkdownDescription: "Communication state of the device, such as \"REACHABLE\".",
				Computed:            true,
			},
			"nw_device_type": schema.StringAttribute{
				MarkdownDescription: "Type of network device.",
				Computed:            true,
			},
			"platform_id": schema.StringAttribute{
				MarkdownDescription: "Platform identifier, such as \"C9KV-UADP-8P\".",
				Computed:            true,
			},
			"location": schema.StringAttribute{
				MarkdownDescription: "Location of the device as a slash-separated path, such as \"Global/USA/New York/DC01\".",
				Computed:            true,
			},
			"ha_status": schema.StringAttribute{
				MarkdownDescription: "High availability status.",
				Computed:            true,
			},
			"os_type": schema.StringAttribute{
				MarkdownDescription: "Corresponds to the `catalystcenter_network_devices.devices.*.software_type`.",
				Computed:            true,
			},
		},
	}
}

func (d *DeviceDetailDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *DeviceDetailDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DeviceDetail

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?searchBy=" + url.QueryEscape(config.Id.ValueString())
	params += "&identifier=uuid"
	res, err := d.client.Get(config.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

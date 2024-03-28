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

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &PnPDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &PnPDeviceDataSource{}
)

func NewPnPDeviceDataSource() datasource.DataSource {
	return &PnPDeviceDataSource{}
}

type PnPDeviceDataSource struct {
	client *cc.Client
}

func (d *PnPDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pnp_device"
}

func (d *PnPDeviceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the PnP Device.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Optional:            true,
				Computed:            true,
			},
			"serial_number": schema.StringAttribute{
				MarkdownDescription: "Device serial number",
				Optional:            true,
				Computed:            true,
			},
			"stack": schema.BoolAttribute{
				MarkdownDescription: "Device is a stacked switch",
				Computed:            true,
			},
			"pid": schema.StringAttribute{
				MarkdownDescription: "Device product ID",
				Computed:            true,
			},
			"hostname": schema.StringAttribute{
				MarkdownDescription: "Device hostname",
				Computed:            true,
			},
		},
	}
}
func (d *PnPDeviceDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("serial_number"),
		),
	}
}

func (d *PnPDeviceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *PnPDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config PnPDevice

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))
	if config.Id.IsNull() && !config.SerialNumber.IsNull() {
		res, err := d.client.Get(config.getPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
			return
		}
		if value := res; len(value.Array()) > 0 {
			value.ForEach(func(k, v gjson.Result) bool {
				if config.SerialNumber.ValueString() == v.Get("deviceInfo.serialNumber").String() {
					config.Id = types.StringValue(v.Get("id").String())
					tflog.Debug(ctx, fmt.Sprintf("%s: Found object with serialNumber '%v', id: %v", config.Id.String(), config.SerialNumber.ValueString(), config.Id.String()))
					return false
				}
				return true
			})
		}

		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with serialNumber: %s", config.SerialNumber.ValueString()))
			return
		}
	}

	params := ""
	params += "/" + url.QueryEscape(config.Id.ValueString())
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

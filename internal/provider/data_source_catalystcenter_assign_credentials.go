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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &AssignCredentialsDataSource{}
	_ datasource.DataSourceWithConfigure = &AssignCredentialsDataSource{}
)

func NewAssignCredentialsDataSource() datasource.DataSource {
	return &AssignCredentialsDataSource{}
}

type AssignCredentialsDataSource struct {
	client *cc.Client
}

func (d *AssignCredentialsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_assign_credentials"
}

func (d *AssignCredentialsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the assigned credentials of a site.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"site_id": schema.StringAttribute{
				MarkdownDescription: "The site ID",
				Required:            true,
			},
			"cli_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the CLI credentials",
				Computed:            true,
			},
			"snmp_v2_read_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the SNMPv2 read credentials",
				Computed:            true,
			},
			"snmp_v2_write_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the SNMPv2 write credentials",
				Computed:            true,
			},
			"snmp_v3_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the SNMPv3 credentials",
				Computed:            true,
			},
			"https_read_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the HTTPS read credentials",
				Computed:            true,
			},
			"https_write_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the HTTPS write credentials",
				Computed:            true,
			},
		},
	}
}

func (d *AssignCredentialsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *AssignCredentialsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config AssignCredentials

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "/" + url.QueryEscape(config.Id.ValueString())
	res, err := d.client.Get("/api/v1/commonsetting/global" + params)
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

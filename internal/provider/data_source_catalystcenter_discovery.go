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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

//template:begin imports
import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

//template:end imports

//template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &DiscoveryDataSource{}
	_ datasource.DataSourceWithConfigure = &DiscoveryDataSource{}
)

func NewDiscoveryDataSource() datasource.DataSource {
	return &DiscoveryDataSource{}
}

type DiscoveryDataSource struct {
	client *cc.Client
}

func (d *DiscoveryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_discovery"
}

func (d *DiscoveryDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Discovery.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"cdp_level": schema.Int64Attribute{
				MarkdownDescription: "CDP level is the number of hops between neighbor devices.",
				Computed:            true,
			},
			"discovery_type": schema.StringAttribute{
				MarkdownDescription: "Type of Discovery.",
				Computed:            true,
			},
			"enable_password_list": schema.ListAttribute{
				MarkdownDescription: "",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"global_credential_id_list": schema.ListAttribute{
				MarkdownDescription: "A list of IDs, which must include SNMP credential and CLI credential.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"http_read_credential": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"http_write_credential": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ip_address_list": schema.StringAttribute{
				MarkdownDescription: "A string of IP address ranges to discover.  E.g.: '172.30.0.1' for discovery_type Single, CDP and LLDP; '172.30.0.1-172.30.0.4' for Range; '72.30.0.1-172.30.0.4,172.31.0.1-172.31.0.4' for Multi Range; '172.30.0.1/20' for CIDR.",
				Computed:            true,
			},
			"ip_filter_list": schema.ListAttribute{
				MarkdownDescription: "A list of IP address ranges to exclude from the discovery.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"lldp_level": schema.Int64Attribute{
				MarkdownDescription: "LLDP level to which neighbor devices to be discovered.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "A name of the discovery.",
				Computed:            true,
			},
			"netconf_port": schema.StringAttribute{
				MarkdownDescription: "Port number for netconf as a string. It requires SSH protocol to work.",
				Computed:            true,
			},
			"password_list": schema.ListAttribute{
				MarkdownDescription: "",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"preferred_ip_method": schema.StringAttribute{
				MarkdownDescription: "Preferred method for selecting management IP address.",
				Computed:            true,
			},
			"protocol_order": schema.StringAttribute{
				MarkdownDescription: "A string of comma-separated protocols (SSH/Telnet), in the same order in which the connections to each device are attempted. E.g.: 'Telnet': only telnet; 'SSH,Telnet': ssh first, with telnet fallback.",
				Computed:            true,
			},
			"retry": schema.Int64Attribute{
				MarkdownDescription: "Number of times to try establishing SSH/Telnet connection to a device.",
				Computed:            true,
			},
			"snmp_auth_passphrase": schema.StringAttribute{
				MarkdownDescription: "Auth passphrase for SNMP.",
				Computed:            true,
			},
			"snmp_auth_protocol": schema.StringAttribute{
				MarkdownDescription: "SNMP auth protocol.",
				Computed:            true,
			},
			"snmp_mode": schema.StringAttribute{
				MarkdownDescription: "Mode of SNMP. The `snmp_auth_protocol` and `snmp_auth_passphrase` are required for \"AuthNoPriv\" mode. Additionally, `snmp_priv_protocol` and `snmp_priv_passphrase` are required for \"AuthPriv\" mode.",
				Computed:            true,
			},
			"snmp_priv_passphrase": schema.StringAttribute{
				MarkdownDescription: "Passphrase for SNMP privacy.",
				Computed:            true,
			},
			"snmp_priv_protocol": schema.StringAttribute{
				MarkdownDescription: "SNMP privacy protocol.",
				Computed:            true,
			},
			"snmp_ro_community": schema.StringAttribute{
				MarkdownDescription: "SNMP RO community of the devices to be discovered.",
				Computed:            true,
			},
			"snmp_ro_community_desc": schema.StringAttribute{
				MarkdownDescription: "Description for snmp_ro_community.",
				Computed:            true,
			},
			"snmp_rw_community": schema.StringAttribute{
				MarkdownDescription: "SNMP RW community of the devices to be discovered.",
				Computed:            true,
			},
			"snmp_rw_community_desc": schema.StringAttribute{
				MarkdownDescription: "Description for snmp_rw_community",
				Computed:            true,
			},
			"snmp_user_name": schema.StringAttribute{
				MarkdownDescription: "SNMP username of the devices to be discovered.",
				Computed:            true,
			},
			"snmp_version": schema.StringAttribute{
				MarkdownDescription: "SNMP version",
				Computed:            true,
			},
			"timeout_seconds": schema.Int64Attribute{
				MarkdownDescription: "Number of seconds to wait for each SSH/Telnet connection to a device.",
				Computed:            true,
			},
			"user_name_list": schema.ListAttribute{
				MarkdownDescription: "",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}
}

func (d *DiscoveryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

//template:end model

//template:begin read
func (d *DiscoveryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config Discovery

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	res, err := d.client.Get("/dna/intent/api/v1/discovery/1/500" + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	res = res.Get("response.#(id==\"" + config.Id.ValueString() + "\")")

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

//template:end read

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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type VirtualNetworkIPPool struct {
	Id                   types.String `tfsdk:"id"`
	VirtualNetworkName   types.String `tfsdk:"virtual_network_name"`
	SiteNameHierarchy    types.String `tfsdk:"site_name_hierarchy"`
	Layer2Only           types.Bool   `tfsdk:"layer2_only"`
	IpPoolName           types.String `tfsdk:"ip_pool_name"`
	VlanId               types.String `tfsdk:"vlan_id"`
	VlanName             types.String `tfsdk:"vlan_name"`
	AutoGenerateVlanName types.Bool   `tfsdk:"auto_generate_vlan_name"`
	TrafficType          types.String `tfsdk:"traffic_type"`
	ScalableGroupName    types.String `tfsdk:"scalable_group_name"`
	L2FloodingEnabled    types.Bool   `tfsdk:"l2_flooding_enabled"`
	CriticalPool         types.Bool   `tfsdk:"critical_pool"`
	WirelessPool         types.Bool   `tfsdk:"wireless_pool"`
	IpDirectedBroadcast  types.Bool   `tfsdk:"ip_directed_broadcast"`
	CommonPool           types.Bool   `tfsdk:"common_pool"`
	BridgeModeVm         types.Bool   `tfsdk:"bridge_mode_vm"`
	PoolType             types.String `tfsdk:"pool_type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data VirtualNetworkIPPool) getPath() string {
	return "/dna/intent/api/v1/business/sda/virtualnetwork/ippool"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data VirtualNetworkIPPool) toBody(ctx context.Context, state VirtualNetworkIPPool) string {
	body := ""
	if !data.VirtualNetworkName.IsNull() {
		body, _ = sjson.Set(body, "virtualNetworkName", data.VirtualNetworkName.ValueString())
	}
	if !data.SiteNameHierarchy.IsNull() {
		body, _ = sjson.Set(body, "siteNameHierarchy", data.SiteNameHierarchy.ValueString())
	}
	if !data.Layer2Only.IsNull() {
		body, _ = sjson.Set(body, "isLayer2Only", data.Layer2Only.ValueBool())
	}
	if !data.IpPoolName.IsNull() {
		body, _ = sjson.Set(body, "ipPoolName", data.IpPoolName.ValueString())
	}
	if !data.VlanId.IsNull() {
		body, _ = sjson.Set(body, "vlanId", data.VlanId.ValueString())
	}
	if !data.VlanName.IsNull() {
		body, _ = sjson.Set(body, "vlanName", data.VlanName.ValueString())
	}
	if !data.AutoGenerateVlanName.IsNull() {
		body, _ = sjson.Set(body, "autoGenerateVlanName", data.AutoGenerateVlanName.ValueBool())
	}
	if !data.TrafficType.IsNull() {
		body, _ = sjson.Set(body, "trafficType", data.TrafficType.ValueString())
	}
	if !data.ScalableGroupName.IsNull() {
		body, _ = sjson.Set(body, "scalableGroupName", data.ScalableGroupName.ValueString())
	}
	if !data.L2FloodingEnabled.IsNull() {
		body, _ = sjson.Set(body, "isL2FloodingEnabled", data.L2FloodingEnabled.ValueBool())
	}
	if !data.CriticalPool.IsNull() {
		body, _ = sjson.Set(body, "isThisCriticalPool", data.CriticalPool.ValueBool())
	}
	if !data.WirelessPool.IsNull() {
		body, _ = sjson.Set(body, "isWirelessPool", data.WirelessPool.ValueBool())
	}
	if !data.IpDirectedBroadcast.IsNull() {
		body, _ = sjson.Set(body, "isIpDirectedBroadcast", data.IpDirectedBroadcast.ValueBool())
	}
	if !data.CommonPool.IsNull() {
		body, _ = sjson.Set(body, "isCommonPool", data.CommonPool.ValueBool())
	}
	if !data.BridgeModeVm.IsNull() {
		body, _ = sjson.Set(body, "isBridgeModeVm", data.BridgeModeVm.ValueBool())
	}
	if !data.PoolType.IsNull() {
		body, _ = sjson.Set(body, "poolType", data.PoolType.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *VirtualNetworkIPPool) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("virtualNetworkName"); value.Exists() {
		data.VirtualNetworkName = types.StringValue(value.String())
	} else {
		data.VirtualNetworkName = types.StringNull()
	}
	if value := res.Get("siteNameHierarchy"); value.Exists() {
		data.SiteNameHierarchy = types.StringValue(value.String())
	} else {
		data.SiteNameHierarchy = types.StringNull()
	}
	if value := res.Get("isLayer2Only"); value.Exists() {
		data.Layer2Only = types.BoolValue(value.Bool())
	} else {
		data.Layer2Only = types.BoolValue(false)
	}
	if value := res.Get("ipPoolName"); value.Exists() {
		data.IpPoolName = types.StringValue(value.String())
	} else {
		data.IpPoolName = types.StringNull()
	}
	if value := res.Get("vlanId"); value.Exists() {
		data.VlanId = types.StringValue(value.String())
	} else {
		data.VlanId = types.StringNull()
	}
	if value := res.Get("vlanName"); value.Exists() {
		data.VlanName = types.StringValue(value.String())
	} else {
		data.VlanName = types.StringNull()
	}
	if value := res.Get("autoGenerateVlanName"); value.Exists() {
		data.AutoGenerateVlanName = types.BoolValue(value.Bool())
	} else {
		data.AutoGenerateVlanName = types.BoolNull()
	}
	if value := res.Get("trafficType"); value.Exists() {
		data.TrafficType = types.StringValue(value.String())
	} else {
		data.TrafficType = types.StringNull()
	}
	if value := res.Get("scalableGroupName"); value.Exists() {
		data.ScalableGroupName = types.StringValue(value.String())
	} else {
		data.ScalableGroupName = types.StringNull()
	}
	if value := res.Get("isSelectiveFloodingEnabled"); value.Exists() {
		data.L2FloodingEnabled = types.BoolValue(value.Bool())
	} else {
		data.L2FloodingEnabled = types.BoolNull()
	}
	if value := res.Get("isWirelessPool"); value.Exists() {
		data.WirelessPool = types.BoolValue(value.Bool())
	} else {
		data.WirelessPool = types.BoolNull()
	}
	if value := res.Get("isIpDirectedBroadcast"); value.Exists() {
		data.IpDirectedBroadcast = types.BoolValue(value.Bool())
	} else {
		data.IpDirectedBroadcast = types.BoolNull()
	}
	if value := res.Get("isCommonPool"); value.Exists() {
		data.CommonPool = types.BoolValue(value.Bool())
	} else {
		data.CommonPool = types.BoolNull()
	}
	if value := res.Get("isBridgeModeVm"); value.Exists() {
		data.BridgeModeVm = types.BoolValue(value.Bool())
	} else {
		data.BridgeModeVm = types.BoolNull()
	}
	if value := res.Get("poolType"); value.Exists() {
		data.PoolType = types.StringValue(value.String())
	} else {
		data.PoolType = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *VirtualNetworkIPPool) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("virtualNetworkName"); value.Exists() && !data.VirtualNetworkName.IsNull() {
		data.VirtualNetworkName = types.StringValue(value.String())
	} else {
		data.VirtualNetworkName = types.StringNull()
	}
	if value := res.Get("siteNameHierarchy"); value.Exists() && !data.SiteNameHierarchy.IsNull() {
		data.SiteNameHierarchy = types.StringValue(value.String())
	} else {
		data.SiteNameHierarchy = types.StringNull()
	}
	if value := res.Get("isLayer2Only"); value.Exists() && !data.Layer2Only.IsNull() {
		data.Layer2Only = types.BoolValue(value.Bool())
	} else if data.Layer2Only.ValueBool() != false {
		data.Layer2Only = types.BoolNull()
	}
	if value := res.Get("ipPoolName"); value.Exists() && !data.IpPoolName.IsNull() {
		data.IpPoolName = types.StringValue(value.String())
	} else {
		data.IpPoolName = types.StringNull()
	}
	if value := res.Get("vlanId"); value.Exists() && !data.VlanId.IsNull() {
		data.VlanId = types.StringValue(value.String())
	} else {
		data.VlanId = types.StringNull()
	}
	if value := res.Get("vlanName"); value.Exists() && !data.VlanName.IsNull() {
		data.VlanName = types.StringValue(value.String())
	} else {
		data.VlanName = types.StringNull()
	}
	if value := res.Get("autoGenerateVlanName"); value.Exists() && !data.AutoGenerateVlanName.IsNull() {
		data.AutoGenerateVlanName = types.BoolValue(value.Bool())
	} else {
		data.AutoGenerateVlanName = types.BoolNull()
	}
	if value := res.Get("trafficType"); value.Exists() && !data.TrafficType.IsNull() {
		data.TrafficType = types.StringValue(value.String())
	} else {
		data.TrafficType = types.StringNull()
	}
	if value := res.Get("scalableGroupName"); value.Exists() && !data.ScalableGroupName.IsNull() {
		data.ScalableGroupName = types.StringValue(value.String())
	} else {
		data.ScalableGroupName = types.StringNull()
	}
	if value := res.Get("isSelectiveFloodingEnabled"); value.Exists() && !data.L2FloodingEnabled.IsNull() {
		data.L2FloodingEnabled = types.BoolValue(value.Bool())
	} else {
		data.L2FloodingEnabled = types.BoolNull()
	}
	if value := res.Get("isWirelessPool"); value.Exists() && !data.WirelessPool.IsNull() {
		data.WirelessPool = types.BoolValue(value.Bool())
	} else {
		data.WirelessPool = types.BoolNull()
	}
	if value := res.Get("isIpDirectedBroadcast"); value.Exists() && !data.IpDirectedBroadcast.IsNull() {
		data.IpDirectedBroadcast = types.BoolValue(value.Bool())
	} else {
		data.IpDirectedBroadcast = types.BoolNull()
	}
	if value := res.Get("isCommonPool"); value.Exists() && !data.CommonPool.IsNull() {
		data.CommonPool = types.BoolValue(value.Bool())
	} else {
		data.CommonPool = types.BoolNull()
	}
	if value := res.Get("isBridgeModeVm"); value.Exists() && !data.BridgeModeVm.IsNull() {
		data.BridgeModeVm = types.BoolValue(value.Bool())
	} else {
		data.BridgeModeVm = types.BoolNull()
	}
	if value := res.Get("poolType"); value.Exists() && !data.PoolType.IsNull() {
		data.PoolType = types.StringValue(value.String())
	} else {
		data.PoolType = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *VirtualNetworkIPPool) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.VirtualNetworkName.IsNull() {
		return false
	}
	if !data.SiteNameHierarchy.IsNull() {
		return false
	}
	if !data.Layer2Only.IsNull() {
		return false
	}
	if !data.IpPoolName.IsNull() {
		return false
	}
	if !data.VlanId.IsNull() {
		return false
	}
	if !data.VlanName.IsNull() {
		return false
	}
	if !data.AutoGenerateVlanName.IsNull() {
		return false
	}
	if !data.TrafficType.IsNull() {
		return false
	}
	if !data.ScalableGroupName.IsNull() {
		return false
	}
	if !data.L2FloodingEnabled.IsNull() {
		return false
	}
	if !data.CriticalPool.IsNull() {
		return false
	}
	if !data.WirelessPool.IsNull() {
		return false
	}
	if !data.IpDirectedBroadcast.IsNull() {
		return false
	}
	if !data.CommonPool.IsNull() {
		return false
	}
	if !data.BridgeModeVm.IsNull() {
		return false
	}
	if !data.PoolType.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull

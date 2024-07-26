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

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type Tag struct {
	Id           types.String      `tfsdk:"id"`
	Name         types.String      `tfsdk:"name"`
	Description  types.String      `tfsdk:"description"`
	SystemTag    types.Bool        `tfsdk:"system_tag"`
	DynamicRules []TagDynamicRules `tfsdk:"dynamic_rules"`
}

type TagDynamicRules struct {
	MemberType    types.String               `tfsdk:"member_type"`
	RuleValues    types.List                 `tfsdk:"rule_values"`
	RuleItems     []TagDynamicRulesRuleItems `tfsdk:"rule_items"`
	RuleOperation types.String               `tfsdk:"rule_operation"`
	RuleName      types.String               `tfsdk:"rule_name"`
	RuleValue     types.String               `tfsdk:"rule_value"`
}

type TagDynamicRulesRuleItems struct {
	Operation types.String `tfsdk:"operation"`
	Name      types.String `tfsdk:"name"`
	Value     types.String `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data Tag) getPath() string {
	return "/dna/intent/api/v1/tag"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data Tag) toBody(ctx context.Context, state Tag) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
		body, _ = sjson.Set(body, "id", state.Id.ValueString())
	}
	_ = put
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.SystemTag.IsNull() {
		body, _ = sjson.Set(body, "systemTag", data.SystemTag.ValueBool())
	}
	if len(data.DynamicRules) > 0 {
		body, _ = sjson.Set(body, "dynamicRules", []interface{}{})
		for _, item := range data.DynamicRules {
			itemBody := ""
			if !item.MemberType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "memberType", item.MemberType.ValueString())
			}
			if !item.RuleValues.IsNull() {
				var values []string
				item.RuleValues.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "rules.values", values)
			}
			if len(item.RuleItems) > 0 {
				itemBody, _ = sjson.Set(itemBody, "rules.items", []interface{}{})
				for _, childItem := range item.RuleItems {
					itemChildBody := ""
					if !childItem.Operation.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "operation", childItem.Operation.ValueString())
					}
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "rules.items.-1", itemChildBody)
				}
			}
			if !item.RuleOperation.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "rules.operation", item.RuleOperation.ValueString())
			}
			if !item.RuleName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "rules.name", item.RuleName.ValueString())
			}
			if !item.RuleValue.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "rules.value", item.RuleValue.ValueString())
			}
			body, _ = sjson.SetRaw(body, "dynamicRules.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *Tag) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get("response.0.id"); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("response.0.name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.0.description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("response.0.systemTag"); value.Exists() {
		data.SystemTag = types.BoolValue(value.Bool())
	} else {
		data.SystemTag = types.BoolNull()
	}
	if value := res.Get("response.0.dynamicRules"); value.Exists() && len(value.Array()) > 0 {
		data.DynamicRules = make([]TagDynamicRules, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TagDynamicRules{}
			if cValue := v.Get("memberType"); cValue.Exists() {
				item.MemberType = types.StringValue(cValue.String())
			} else {
				item.MemberType = types.StringNull()
			}
			if cValue := v.Get("rules.values"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.RuleValues = helpers.GetStringList(cValue.Array())
			} else {
				item.RuleValues = types.ListNull(types.StringType)
			}
			if cValue := v.Get("rules.items"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.RuleItems = make([]TagDynamicRulesRuleItems, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := TagDynamicRulesRuleItems{}
					if ccValue := cv.Get("operation"); ccValue.Exists() {
						cItem.Operation = types.StringValue(ccValue.String())
					} else {
						cItem.Operation = types.StringNull()
					}
					if ccValue := cv.Get("name"); ccValue.Exists() {
						cItem.Name = types.StringValue(ccValue.String())
					} else {
						cItem.Name = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() {
						cItem.Value = types.StringValue(ccValue.String())
					} else {
						cItem.Value = types.StringNull()
					}
					item.RuleItems = append(item.RuleItems, cItem)
					return true
				})
			}
			if cValue := v.Get("rules.operation"); cValue.Exists() {
				item.RuleOperation = types.StringValue(cValue.String())
			} else {
				item.RuleOperation = types.StringNull()
			}
			if cValue := v.Get("rules.name"); cValue.Exists() {
				item.RuleName = types.StringValue(cValue.String())
			} else {
				item.RuleName = types.StringNull()
			}
			if cValue := v.Get("rules.value"); cValue.Exists() {
				item.RuleValue = types.StringValue(cValue.String())
			} else {
				item.RuleValue = types.StringNull()
			}
			data.DynamicRules = append(data.DynamicRules, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *Tag) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.0.description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("response.0.systemTag"); value.Exists() && !data.SystemTag.IsNull() {
		data.SystemTag = types.BoolValue(value.Bool())
	} else {
		data.SystemTag = types.BoolNull()
	}
	for i := range data.DynamicRules {
		keys := [...]string{"memberType", "rules.operation", "rules.name", "rules.value"}
		keyValues := [...]string{data.DynamicRules[i].MemberType.ValueString(), data.DynamicRules[i].RuleOperation.ValueString(), data.DynamicRules[i].RuleName.ValueString(), data.DynamicRules[i].RuleValue.ValueString()}

		var r gjson.Result
		res.Get("response.0.dynamicRules").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("memberType"); value.Exists() && !data.DynamicRules[i].MemberType.IsNull() {
			data.DynamicRules[i].MemberType = types.StringValue(value.String())
		} else {
			data.DynamicRules[i].MemberType = types.StringNull()
		}
		if value := r.Get("rules.values"); value.Exists() && !data.DynamicRules[i].RuleValues.IsNull() {
			data.DynamicRules[i].RuleValues = helpers.GetStringList(value.Array())
		} else {
			data.DynamicRules[i].RuleValues = types.ListNull(types.StringType)
		}
		for ci := range data.DynamicRules[i].RuleItems {
			keys := [...]string{"operation", "name", "value"}
			keyValues := [...]string{data.DynamicRules[i].RuleItems[ci].Operation.ValueString(), data.DynamicRules[i].RuleItems[ci].Name.ValueString(), data.DynamicRules[i].RuleItems[ci].Value.ValueString()}

			var cr gjson.Result
			r.Get("rules.items").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() == keyValues[ik] {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			if value := cr.Get("operation"); value.Exists() && !data.DynamicRules[i].RuleItems[ci].Operation.IsNull() {
				data.DynamicRules[i].RuleItems[ci].Operation = types.StringValue(value.String())
			} else {
				data.DynamicRules[i].RuleItems[ci].Operation = types.StringNull()
			}
			if value := cr.Get("name"); value.Exists() && !data.DynamicRules[i].RuleItems[ci].Name.IsNull() {
				data.DynamicRules[i].RuleItems[ci].Name = types.StringValue(value.String())
			} else {
				data.DynamicRules[i].RuleItems[ci].Name = types.StringNull()
			}
			if value := cr.Get("value"); value.Exists() && !data.DynamicRules[i].RuleItems[ci].Value.IsNull() {
				data.DynamicRules[i].RuleItems[ci].Value = types.StringValue(value.String())
			} else {
				data.DynamicRules[i].RuleItems[ci].Value = types.StringNull()
			}
		}
		if value := r.Get("rules.operation"); value.Exists() && !data.DynamicRules[i].RuleOperation.IsNull() {
			data.DynamicRules[i].RuleOperation = types.StringValue(value.String())
		} else {
			data.DynamicRules[i].RuleOperation = types.StringNull()
		}
		if value := r.Get("rules.name"); value.Exists() && !data.DynamicRules[i].RuleName.IsNull() {
			data.DynamicRules[i].RuleName = types.StringValue(value.String())
		} else {
			data.DynamicRules[i].RuleName = types.StringNull()
		}
		if value := r.Get("rules.value"); value.Exists() && !data.DynamicRules[i].RuleValue.IsNull() {
			data.DynamicRules[i].RuleValue = types.StringValue(value.String())
		} else {
			data.DynamicRules[i].RuleValue = types.StringNull()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *Tag) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Description.IsNull() {
		return false
	}
	if !data.SystemTag.IsNull() {
		return false
	}
	if len(data.DynamicRules) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull

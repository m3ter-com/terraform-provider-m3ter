// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package permission_policy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type PermissionPolicyDataSourceModel struct {
	ID               types.String                                                                  `tfsdk:"id" path:"id,required"`
	OrgID            types.String                                                                  `tfsdk:"org_id" path:"orgId,optional"`
	ManagedPolicy    types.Bool                                                                    `tfsdk:"managed_policy" json:"managedPolicy,computed"`
	Name             types.String                                                                  `tfsdk:"name" json:"name,computed"`
	Version          types.Int64                                                                   `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	PermissionPolicy customfield.NestedObjectList[PermissionPolicyPermissionPolicyDataSourceModel] `tfsdk:"permission_policy" json:"permissionPolicy,computed"`
}

func (m *PermissionPolicyDataSourceModel) toReadParams(_ context.Context) (params m3ter.PermissionPolicyGetParams, diags diag.Diagnostics) {
	params = m3ter.PermissionPolicyGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type PermissionPolicyPermissionPolicyDataSourceModel struct {
	Action   customfield.List[types.String] `tfsdk:"action" json:"action,computed"`
	Effect   types.String                   `tfsdk:"effect" json:"effect,computed"`
	Resource customfield.List[types.String] `tfsdk:"resource" json:"resource,computed"`
}

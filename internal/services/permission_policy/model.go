// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package permission_policy

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type PermissionPolicyModel struct {
	ID               types.String                              `tfsdk:"id" json:"id,computed"`
	OrgID            types.String                              `tfsdk:"org_id" path:"orgId,optional"`
	Name             types.String                              `tfsdk:"name" json:"name,required"`
	PermissionPolicy *[]*PermissionPolicyPermissionPolicyModel `tfsdk:"permission_policy" json:"permissionPolicy,required"`
	ManagedPolicy    types.Bool                                `tfsdk:"managed_policy" json:"managedPolicy,computed"`
	Version          types.Int64                               `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m PermissionPolicyModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m PermissionPolicyModel) MarshalJSONForUpdate(state PermissionPolicyModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type PermissionPolicyPermissionPolicyModel struct {
	Action   *[]types.String `tfsdk:"action" json:"action,required"`
	Effect   types.String    `tfsdk:"effect" json:"effect,required"`
	Resource *[]types.String `tfsdk:"resource" json:"resource,required"`
}

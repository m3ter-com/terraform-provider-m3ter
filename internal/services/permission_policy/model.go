// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package permission_policy

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type PermissionPolicyModel struct {
	ID               types.String                              `tfsdk:"id" json:"id,computed"`
	OrgID            types.String                              `tfsdk:"org_id" path:"orgId,optional"`
	Name             types.String                              `tfsdk:"name" json:"name,required"`
	PermissionPolicy *[]*PermissionPolicyPermissionPolicyModel `tfsdk:"permission_policy" json:"permissionPolicy,required"`
	CreatedBy        types.String                              `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated        timetypes.RFC3339                         `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified   timetypes.RFC3339                         `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy   types.String                              `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	ManagedPolicy    types.Bool                                `tfsdk:"managed_policy" json:"managedPolicy,computed"`
	Version          types.Int64                               `tfsdk:"version" json:"version,computed"`
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

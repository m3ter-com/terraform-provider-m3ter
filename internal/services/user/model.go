// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type UserModel struct {
ID types.String `tfsdk:"id" path:"id,required"`
OrgID types.String `tfsdk:"org_id" path:"orgId,required"`
DtEndAccess timetypes.RFC3339 `tfsdk:"dt_end_access" json:"dtEndAccess,optional" format:"date-time"`
Version types.Int64 `tfsdk:"version" json:"version,optional"`
PermissionPolicy customfield.NestedObjectList[UserPermissionPolicyModel] `tfsdk:"permission_policy" json:"permissionPolicy,computed_optional"`
ContactNumber types.String `tfsdk:"contact_number" json:"contactNumber,computed"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
Email types.String `tfsdk:"email" json:"email,computed"`
FirstAcceptedTermsAndConditions timetypes.RFC3339 `tfsdk:"first_accepted_terms_and_conditions" json:"firstAcceptedTermsAndConditions,computed" format:"date-time"`
FirstName types.String `tfsdk:"first_name" json:"firstName,computed"`
LastAcceptedTermsAndConditions timetypes.RFC3339 `tfsdk:"last_accepted_terms_and_conditions" json:"lastAcceptedTermsAndConditions,computed" format:"date-time"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
LastName types.String `tfsdk:"last_name" json:"lastName,computed"`
SupportUser types.Bool `tfsdk:"support_user" json:"supportUser,computed"`
Organizations customfield.List[types.String] `tfsdk:"organizations" json:"organizations,computed"`
}

func (m UserModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m UserModel) MarshalJSONForUpdate(state UserModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}

type UserPermissionPolicyModel struct {
Action *[]types.String `tfsdk:"action" json:"action,required"`
Effect types.String `tfsdk:"effect" json:"effect,required"`
Resource *[]types.String `tfsdk:"resource" json:"resource,required"`
}

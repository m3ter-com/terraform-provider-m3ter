// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type UsersDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[UsersItemsDataSourceModel] `json:"data,computed"`
}

type UsersDataSourceModel struct {
	OrgID    types.String                                            `tfsdk:"org_id" path:"orgId,required"`
	IDs      *[]types.String                                         `tfsdk:"ids" query:"ids,optional"`
	MaxItems types.Int64                                             `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[UsersItemsDataSourceModel] `tfsdk:"items"`
}

func (m *UsersDataSourceModel) toListParams(_ context.Context) (params m3ter.UserListParams, diags diag.Diagnostics) {
	mIDs := []string{}
	for _, item := range *m.IDs {
		mIDs = append(mIDs, item.ValueString())
	}

	params = m3ter.UserListParams{
		OrgID: m3ter.F(m.OrgID.ValueString()),
		IDs:   m3ter.F(mIDs),
	}

	return
}

type UsersItemsDataSourceModel struct {
	ID                              types.String                                                       `tfsdk:"id" json:"id,computed"`
	ContactNumber                   types.String                                                       `tfsdk:"contact_number" json:"contactNumber,computed"`
	CreatedBy                       types.String                                                       `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated                       timetypes.RFC3339                                                  `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtEndAccess                     timetypes.RFC3339                                                  `tfsdk:"dt_end_access" json:"dtEndAccess,computed" format:"date-time"`
	DtLastModified                  timetypes.RFC3339                                                  `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	Email                           types.String                                                       `tfsdk:"email" json:"email,computed"`
	FirstAcceptedTermsAndConditions timetypes.RFC3339                                                  `tfsdk:"first_accepted_terms_and_conditions" json:"firstAcceptedTermsAndConditions,computed" format:"date-time"`
	FirstName                       types.String                                                       `tfsdk:"first_name" json:"firstName,computed"`
	LastAcceptedTermsAndConditions  timetypes.RFC3339                                                  `tfsdk:"last_accepted_terms_and_conditions" json:"lastAcceptedTermsAndConditions,computed" format:"date-time"`
	LastModifiedBy                  types.String                                                       `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	LastName                        types.String                                                       `tfsdk:"last_name" json:"lastName,computed"`
	Organizations                   customfield.List[types.String]                                     `tfsdk:"organizations" json:"organizations,computed"`
	PermissionPolicy                customfield.NestedObjectList[UsersPermissionPolicyDataSourceModel] `tfsdk:"permission_policy" json:"permissionPolicy,computed"`
	SupportUser                     types.Bool                                                         `tfsdk:"support_user" json:"supportUser,computed"`
	Version                         types.Int64                                                        `tfsdk:"version" json:"version,computed"`
}

type UsersPermissionPolicyDataSourceModel struct {
	Action   customfield.List[types.String] `tfsdk:"action" json:"action,computed"`
	Effect   types.String                   `tfsdk:"effect" json:"effect,computed"`
	Resource customfield.List[types.String] `tfsdk:"resource" json:"resource,computed"`
}

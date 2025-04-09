// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_invitation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type UserInvitationsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[UserInvitationsItemsDataSourceModel] `json:"data,computed"`
}

type UserInvitationsDataSourceModel struct {
	OrgID    types.String                                                      `tfsdk:"org_id" path:"orgId,optional"`
	MaxItems types.Int64                                                       `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[UserInvitationsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *UserInvitationsDataSourceModel) toListParams(_ context.Context) (params m3ter.UserInvitationListParams, diags diag.Diagnostics) {
	params = m3ter.UserInvitationListParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type UserInvitationsItemsDataSourceModel struct {
	ID                  types.String                   `tfsdk:"id" json:"id,computed"`
	Accepted            types.Bool                     `tfsdk:"accepted" json:"accepted,computed"`
	DtEndAccess         timetypes.RFC3339              `tfsdk:"dt_end_access" json:"dtEndAccess,computed" format:"date-time"`
	DtExpiry            timetypes.RFC3339              `tfsdk:"dt_expiry" json:"dtExpiry,computed" format:"date-time"`
	Email               types.String                   `tfsdk:"email" json:"email,computed"`
	FirstName           types.String                   `tfsdk:"first_name" json:"firstName,computed"`
	InvitingPrincipalID types.String                   `tfsdk:"inviting_principal_id" json:"invitingPrincipalId,computed"`
	LastName            types.String                   `tfsdk:"last_name" json:"lastName,computed"`
	PermissionPolicyIDs customfield.List[types.String] `tfsdk:"permission_policy_ids" json:"permissionPolicyIds,computed"`
	Version             types.Int64                    `tfsdk:"version" json:"version,computed"`
	CreatedBy           types.String                   `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated           timetypes.RFC3339              `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified      timetypes.RFC3339              `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy      types.String                   `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
}

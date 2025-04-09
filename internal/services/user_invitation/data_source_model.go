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

type UserInvitationDataSourceModel struct {
	ID                  types.String                   `tfsdk:"id" path:"id,required"`
	OrgID               types.String                   `tfsdk:"org_id" path:"orgId,required"`
	Accepted            types.Bool                     `tfsdk:"accepted" json:"accepted,computed"`
	CreatedBy           types.String                   `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated           timetypes.RFC3339              `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtEndAccess         timetypes.RFC3339              `tfsdk:"dt_end_access" json:"dtEndAccess,computed" format:"date-time"`
	DtExpiry            timetypes.RFC3339              `tfsdk:"dt_expiry" json:"dtExpiry,computed" format:"date-time"`
	DtLastModified      timetypes.RFC3339              `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	Email               types.String                   `tfsdk:"email" json:"email,computed"`
	FirstName           types.String                   `tfsdk:"first_name" json:"firstName,computed"`
	InvitingPrincipalID types.String                   `tfsdk:"inviting_principal_id" json:"invitingPrincipalId,computed"`
	LastModifiedBy      types.String                   `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	LastName            types.String                   `tfsdk:"last_name" json:"lastName,computed"`
	Version             types.Int64                    `tfsdk:"version" json:"version,computed"`
	PermissionPolicyIDs customfield.List[types.String] `tfsdk:"permission_policy_ids" json:"permissionPolicyIds,computed"`
}

func (m *UserInvitationDataSourceModel) toReadParams(_ context.Context) (params m3ter.UserInvitationGetParams, diags diag.Diagnostics) {
	params = m3ter.UserInvitationGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

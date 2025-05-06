// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_invitation

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type UserInvitationModel struct {
	ID                  types.String      `tfsdk:"id" json:"id,computed"`
	OrgID               types.String      `tfsdk:"org_id" path:"orgId,optional"`
	Email               types.String      `tfsdk:"email" json:"email,required"`
	FirstName           types.String      `tfsdk:"first_name" json:"firstName,required"`
	LastName            types.String      `tfsdk:"last_name" json:"lastName,required"`
	ContactNumber       types.String      `tfsdk:"contact_number" json:"contactNumber,optional,no_refresh"`
	DtEndAccess         timetypes.RFC3339 `tfsdk:"dt_end_access" json:"dtEndAccess,optional" format:"date-time"`
	DtExpiry            timetypes.RFC3339 `tfsdk:"dt_expiry" json:"dtExpiry,optional" format:"date-time"`
	M3terUser           types.Bool        `tfsdk:"m3ter_user" json:"m3terUser,optional,no_refresh"`
	Version             types.Int64       `tfsdk:"version" json:"version,optional"`
	PermissionPolicyIDs *[]types.String   `tfsdk:"permission_policy_ids" json:"permissionPolicyIds,optional"`
	Accepted            types.Bool        `tfsdk:"accepted" json:"accepted,computed"`
	CreatedBy           types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated           timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified      timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	InvitingPrincipalID types.String      `tfsdk:"inviting_principal_id" json:"invitingPrincipalId,computed"`
	LastModifiedBy      types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
}

func (m UserInvitationModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m UserInvitationModel) MarshalJSONForUpdate(state UserInvitationModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

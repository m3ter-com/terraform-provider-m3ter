// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package debit_reason

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type DebitReasonModel struct {
	ID             types.String      `tfsdk:"id" json:"id,computed"`
	OrgID          types.String      `tfsdk:"org_id" path:"orgId,optional"`
	Name           types.String      `tfsdk:"name" json:"name,required"`
	Archived       types.Bool        `tfsdk:"archived" json:"archived,optional"`
	Code           types.String      `tfsdk:"code" json:"code,optional"`
	CreatedBy      types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated      timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Version        types.Int64       `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m DebitReasonModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m DebitReasonModel) MarshalJSONForUpdate(state DebitReasonModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

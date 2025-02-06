// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_config

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BillConfigDataSourceModel struct {
	OrgID          types.String      `tfsdk:"org_id" path:"orgId,required"`
	BillLockDate   timetypes.RFC3339 `tfsdk:"bill_lock_date" json:"billLockDate,computed" format:"date"`
	CreatedBy      types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated      timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	ID             types.String      `tfsdk:"id" json:"id,computed"`
	LastModifiedBy types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Version        types.Int64       `tfsdk:"version" json:"version,computed"`
}

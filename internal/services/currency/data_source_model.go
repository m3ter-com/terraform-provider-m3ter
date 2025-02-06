// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package currency

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CurrencyDataSourceModel struct {
	OrgID            types.String      `tfsdk:"org_id" path:"orgId,required"`
	ID               types.String      `tfsdk:"id" path:"id,computed"`
	Archived         types.Bool        `tfsdk:"archived" json:"archived,computed"`
	Code             types.String      `tfsdk:"code" json:"code,computed"`
	CreatedBy        types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated        timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified   timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy   types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	MaxDecimalPlaces types.Int64       `tfsdk:"max_decimal_places" json:"maxDecimalPlaces,computed"`
	Name             types.String      `tfsdk:"name" json:"name,computed"`
	RoundingMode     types.String      `tfsdk:"rounding_mode" json:"roundingMode,computed"`
	Version          types.Int64       `tfsdk:"version" json:"version,computed"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package aggregation

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type AggregationDataSourceModel struct {
	OrgID           types.String                                    `tfsdk:"org_id" path:"orgId,required"`
	ID              types.String                                    `tfsdk:"id" path:"id,computed"`
	Aggregation     types.String                                    `tfsdk:"aggregation" json:"aggregation,computed"`
	Code            types.String                                    `tfsdk:"code" json:"code,computed"`
	CreatedBy       types.String                                    `tfsdk:"created_by" json:"createdBy,computed"`
	DefaultValue    types.Float64                                   `tfsdk:"default_value" json:"defaultValue,computed"`
	DtCreated       timetypes.RFC3339                               `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified  timetypes.RFC3339                               `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy  types.String                                    `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	MeterID         types.String                                    `tfsdk:"meter_id" json:"meterId,computed"`
	Name            types.String                                    `tfsdk:"name" json:"name,computed"`
	QuantityPerUnit types.Float64                                   `tfsdk:"quantity_per_unit" json:"quantityPerUnit,computed"`
	Rounding        types.String                                    `tfsdk:"rounding" json:"rounding,computed"`
	TargetField     types.String                                    `tfsdk:"target_field" json:"targetField,computed"`
	Unit            types.String                                    `tfsdk:"unit" json:"unit,computed"`
	Version         types.Int64                                     `tfsdk:"version" json:"version,computed"`
	CustomFields    customfield.Map[jsontypes.Normalized]           `tfsdk:"custom_fields" json:"customFields,computed"`
	SegmentedFields customfield.List[types.String]                  `tfsdk:"segmented_fields" json:"segmentedFields,computed"`
	Segments        customfield.List[customfield.Map[types.String]] `tfsdk:"segments" json:"segments,computed"`
}

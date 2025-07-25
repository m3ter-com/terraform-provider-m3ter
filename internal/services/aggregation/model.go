// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package aggregation

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type AggregationModel struct {
	ID                  types.String                `tfsdk:"id" json:"id,computed"`
	OrgID               types.String                `tfsdk:"org_id" path:"orgId,optional"`
	Aggregation         types.String                `tfsdk:"aggregation" json:"aggregation,required"`
	MeterID             types.String                `tfsdk:"meter_id" json:"meterId,required"`
	Name                types.String                `tfsdk:"name" json:"name,required"`
	QuantityPerUnit     types.Float64               `tfsdk:"quantity_per_unit" json:"quantityPerUnit,required"`
	Rounding            types.String                `tfsdk:"rounding" json:"rounding,required"`
	TargetField         types.String                `tfsdk:"target_field" json:"targetField,required"`
	Unit                types.String                `tfsdk:"unit" json:"unit,required"`
	AccountingProductID types.String                `tfsdk:"accounting_product_id" json:"accountingProductId,optional"`
	Code                types.String                `tfsdk:"code" json:"code,optional"`
	CustomSql           types.String                `tfsdk:"custom_sql" json:"customSql,optional"`
	DefaultValue        types.Float64               `tfsdk:"default_value" json:"defaultValue,optional"`
	SegmentedFields     *[]types.String             `tfsdk:"segmented_fields" json:"segmentedFields,optional"`
	Segments            *[]*map[string]types.String `tfsdk:"segments" json:"segments,optional"`
	CustomFields        types.Dynamic               `tfsdk:"custom_fields" json:"customFields,optional"`
	CreatedBy           types.String                `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated           timetypes.RFC3339           `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified      timetypes.RFC3339           `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy      types.String                `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Version             types.Int64                 `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m AggregationModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AggregationModel) MarshalJSONForUpdate(state AggregationModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

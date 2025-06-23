// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compound_aggregation

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type CompoundAggregationModel struct {
	ID                       types.String                                    `tfsdk:"id" json:"id,computed"`
	OrgID                    types.String                                    `tfsdk:"org_id" path:"orgId,optional"`
	Calculation              types.String                                    `tfsdk:"calculation" json:"calculation,required"`
	Name                     types.String                                    `tfsdk:"name" json:"name,required"`
	QuantityPerUnit          types.Float64                                   `tfsdk:"quantity_per_unit" json:"quantityPerUnit,required"`
	Rounding                 types.String                                    `tfsdk:"rounding" json:"rounding,required"`
	Unit                     types.String                                    `tfsdk:"unit" json:"unit,required"`
	AccountingProductID      types.String                                    `tfsdk:"accounting_product_id" json:"accountingProductId,optional"`
	Code                     types.String                                    `tfsdk:"code" json:"code,optional"`
	EvaluateNullAggregations types.Bool                                      `tfsdk:"evaluate_null_aggregations" json:"evaluateNullAggregations,optional"`
	ProductID                types.String                                    `tfsdk:"product_id" json:"productId,optional"`
	CustomFields             types.Dynamic                                   `tfsdk:"custom_fields" json:"customFields,optional"`
	Aggregation              types.String                                    `tfsdk:"aggregation" json:"aggregation,computed,no_refresh"`
	CreatedBy                types.String                                    `tfsdk:"created_by" json:"createdBy,computed"`
	CustomSql                types.String                                    `tfsdk:"custom_sql" json:"customSql,computed,no_refresh"`
	DefaultValue             types.Float64                                   `tfsdk:"default_value" json:"defaultValue,computed,no_refresh"`
	DtCreated                timetypes.RFC3339                               `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified           timetypes.RFC3339                               `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy           types.String                                    `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	MeterID                  types.String                                    `tfsdk:"meter_id" json:"meterId,computed,no_refresh"`
	TargetField              types.String                                    `tfsdk:"target_field" json:"targetField,computed,no_refresh"`
	Version                  types.Int64                                     `tfsdk:"version" json:"version,computed"`
	SegmentedFields          customfield.List[types.String]                  `tfsdk:"segmented_fields" json:"segmentedFields,computed,no_refresh"`
	Segments                 customfield.List[customfield.Map[types.String]] `tfsdk:"segments" json:"segments,computed"`
}

func (m CompoundAggregationModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CompoundAggregationModel) MarshalJSONForUpdate(state CompoundAggregationModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

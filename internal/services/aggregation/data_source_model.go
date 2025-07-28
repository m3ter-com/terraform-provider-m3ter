// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package aggregation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type AggregationDataSourceModel struct {
	ID                  types.String                                    `tfsdk:"id" path:"id,required"`
	OrgID               types.String                                    `tfsdk:"org_id" path:"orgId,required"`
	AccountingProductID types.String                                    `tfsdk:"accounting_product_id" json:"accountingProductId,computed"`
	Aggregation         types.String                                    `tfsdk:"aggregation" json:"aggregation,computed"`
	Code                types.String                                    `tfsdk:"code" json:"code,computed"`
	CustomSql           types.String                                    `tfsdk:"custom_sql" json:"customSql,computed"`
	DefaultValue        types.Float64                                   `tfsdk:"default_value" json:"defaultValue,computed"`
	MeterID             types.String                                    `tfsdk:"meter_id" json:"meterId,computed"`
	Name                types.String                                    `tfsdk:"name" json:"name,computed"`
	QuantityPerUnit     types.Float64                                   `tfsdk:"quantity_per_unit" json:"quantityPerUnit,computed"`
	Rounding            types.String                                    `tfsdk:"rounding" json:"rounding,computed"`
	TargetField         types.String                                    `tfsdk:"target_field" json:"targetField,computed"`
	Unit                types.String                                    `tfsdk:"unit" json:"unit,computed"`
	Version             types.Int64                                     `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	SegmentedFields     customfield.List[types.String]                  `tfsdk:"segmented_fields" json:"segmentedFields,computed"`
	Segments            customfield.List[customfield.Map[types.String]] `tfsdk:"segments" json:"segments,computed"`
	CustomFields        types.Dynamic                                   `tfsdk:"custom_fields" json:"customFields,computed"`
}

func (m *AggregationDataSourceModel) toReadParams(_ context.Context) (params m3ter.AggregationGetParams, diags diag.Diagnostics) {
	params = m3ter.AggregationGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

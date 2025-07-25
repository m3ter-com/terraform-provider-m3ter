// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compound_aggregation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type CompoundAggregationDataSourceModel struct {
	ID                       types.String                                    `tfsdk:"id" path:"id,required"`
	OrgID                    types.String                                    `tfsdk:"org_id" path:"orgId,required"`
	AccountingProductID      types.String                                    `tfsdk:"accounting_product_id" json:"accountingProductId,computed"`
	Calculation              types.String                                    `tfsdk:"calculation" json:"calculation,computed"`
	Code                     types.String                                    `tfsdk:"code" json:"code,computed"`
	CreatedBy                types.String                                    `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated                timetypes.RFC3339                               `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified           timetypes.RFC3339                               `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	EvaluateNullAggregations types.Bool                                      `tfsdk:"evaluate_null_aggregations" json:"evaluateNullAggregations,computed"`
	LastModifiedBy           types.String                                    `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Name                     types.String                                    `tfsdk:"name" json:"name,computed"`
	ProductID                types.String                                    `tfsdk:"product_id" json:"productId,computed"`
	QuantityPerUnit          types.Float64                                   `tfsdk:"quantity_per_unit" json:"quantityPerUnit,computed"`
	Rounding                 types.String                                    `tfsdk:"rounding" json:"rounding,computed"`
	Unit                     types.String                                    `tfsdk:"unit" json:"unit,computed"`
	Version                  types.Int64                                     `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	Segments                 customfield.List[customfield.Map[types.String]] `tfsdk:"segments" json:"segments,computed"`
	CustomFields             types.Dynamic                                   `tfsdk:"custom_fields" json:"customFields,computed"`
}

func (m *CompoundAggregationDataSourceModel) toReadParams(_ context.Context) (params m3ter.CompoundAggregationGetParams, diags diag.Diagnostics) {
	params = m3ter.CompoundAggregationGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

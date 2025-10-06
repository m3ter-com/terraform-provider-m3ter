// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compound_aggregation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type CompoundAggregationDataSourceModel struct {
	ID                       types.String                                    `tfsdk:"id" path:"id,computed_optional"`
	OrgID                    types.String                                    `tfsdk:"org_id" path:"orgId,required"`
	AccountingProductID      types.String                                    `tfsdk:"accounting_product_id" json:"accountingProductId,computed"`
	Calculation              types.String                                    `tfsdk:"calculation" json:"calculation,computed"`
	Code                     types.String                                    `tfsdk:"code" json:"code,computed"`
	EvaluateNullAggregations types.Bool                                      `tfsdk:"evaluate_null_aggregations" json:"evaluateNullAggregations,computed"`
	Name                     types.String                                    `tfsdk:"name" json:"name,computed"`
	ProductID                types.String                                    `tfsdk:"product_id" json:"productId,computed"`
	QuantityPerUnit          types.Float64                                   `tfsdk:"quantity_per_unit" json:"quantityPerUnit,computed"`
	Rounding                 types.String                                    `tfsdk:"rounding" json:"rounding,computed"`
	Unit                     types.String                                    `tfsdk:"unit" json:"unit,computed"`
	Version                  types.Int64                                     `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	Segments                 customfield.List[customfield.Map[types.String]] `tfsdk:"segments" json:"segments,computed"`
	CustomFields             customfield.NormalizedDynamicValue              `tfsdk:"custom_fields" json:"customFields,computed"`
	FindOneBy                *CompoundAggregationFindOneByDataSourceModel    `tfsdk:"find_one_by"`
}

func (m *CompoundAggregationDataSourceModel) toReadParams(_ context.Context) (params m3ter.CompoundAggregationGetParams, diags diag.Diagnostics) {
	params = m3ter.CompoundAggregationGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

func (m *CompoundAggregationDataSourceModel) toListParams(_ context.Context) (params m3ter.CompoundAggregationListParams, diags diag.Diagnostics) {
	mFindOneByCodes := []string{}
	if m.FindOneBy.Codes != nil {
		for _, item := range *m.FindOneBy.Codes {
			mFindOneByCodes = append(mFindOneByCodes, item.ValueString())
		}
	}
	mFindOneByIDs := []string{}
	if m.FindOneBy.IDs != nil {
		for _, item := range *m.FindOneBy.IDs {
			mFindOneByIDs = append(mFindOneByIDs, item.ValueString())
		}
	}
	mFindOneByProductID := []string{}
	if m.FindOneBy.ProductID != nil {
		for _, item := range *m.FindOneBy.ProductID {
			mFindOneByProductID = append(mFindOneByProductID, item.ValueString())
		}
	}

	params = m3ter.CompoundAggregationListParams{
		Codes:     m3ter.F(mFindOneByCodes),
		IDs:       m3ter.F(mFindOneByIDs),
		ProductID: m3ter.F(mFindOneByProductID),
	}

	return
}

type CompoundAggregationFindOneByDataSourceModel struct {
	Codes     *[]types.String `tfsdk:"codes" query:"codes,optional"`
	IDs       *[]types.String `tfsdk:"ids" query:"ids,optional"`
	ProductID *[]types.String `tfsdk:"product_id" query:"productId,optional"`
}

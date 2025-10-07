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
	ID                  types.String                                    `tfsdk:"id" path:"id,computed_optional"`
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
	CustomFields        customfield.NormalizedDynamicValue              `tfsdk:"custom_fields" json:"customFields,computed"`
	FindOneBy           *AggregationFindOneByDataSourceModel            `tfsdk:"find_one_by"`
}

func (m *AggregationDataSourceModel) toReadParams(_ context.Context) (params m3ter.AggregationGetParams, diags diag.Diagnostics) {
	params = m3ter.AggregationGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

func (m *AggregationDataSourceModel) toListParams(_ context.Context) (params m3ter.AggregationListParams, diags diag.Diagnostics) {
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

	params = m3ter.AggregationListParams{
		Codes:     m3ter.F(mFindOneByCodes),
		IDs:       m3ter.F(mFindOneByIDs),
		ProductID: m3ter.F(mFindOneByProductID),
	}

	return
}

type AggregationFindOneByDataSourceModel struct {
	Codes     *[]types.String `tfsdk:"codes" query:"codes,optional"`
	IDs       *[]types.String `tfsdk:"ids" query:"ids,optional"`
	ProductID *[]types.String `tfsdk:"product_id" query:"productId,optional"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compound_aggregation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type CompoundAggregationDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[CompoundAggregationDataSourceModel] `json:"data,computed"`
}

type CompoundAggregationDataSourceModel struct {
	ID                       types.String                                    `tfsdk:"id" path:"id,computed_optional"`
	OrgID                    types.String                                    `tfsdk:"org_id" path:"orgId,required"`
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
	Version                  types.Int64                                     `tfsdk:"version" json:"version,computed"`
	CustomFields             customfield.Map[jsontypes.Normalized]           `tfsdk:"custom_fields" json:"customFields,computed"`
	Segments                 customfield.List[customfield.Map[types.String]] `tfsdk:"segments" json:"segments,computed"`
	FindOneBy                *CompoundAggregationFindOneByDataSourceModel    `tfsdk:"find_one_by"`
}

func (m *CompoundAggregationDataSourceModel) toListParams(_ context.Context) (params m3ter.CompoundAggregationListParams, diags diag.Diagnostics) {
	mFindOneByCodes := []CompoundAggregationListParamsCode{}
	for _, item := range *m.FindOneBy.Codes {
		mFindOneByCodes = append(mFindOneByCodes, item.ValueString())
	}
	mFindOneByIDs := []CompoundAggregationListParamsID{}
	for _, item := range *m.FindOneBy.IDs {
		mFindOneByIDs = append(mFindOneByIDs, item.ValueString())
	}
	mFindOneByProductID := []CompoundAggregationListParamsProductID{}
	for _, item := range *m.FindOneBy.ProductID {
		mFindOneByProductID = append(mFindOneByProductID, item.ValueString())
	}

	params = m3ter.CompoundAggregationListParams{
		Codes:     m3ter.F(mFindOneByCodes),
		IDs:       m3ter.F(mFindOneByIDs),
		ProductID: m3ter.F(mFindOneByProductID),
	}

	return
}

type CompoundAggregationFindOneByDataSourceModel struct {
	Codes     *[]types.String         `tfsdk:"codes" query:"codes,optional"`
	IDs       *[]types.String         `tfsdk:"ids" query:"ids,optional"`
	ProductID *[]jsontypes.Normalized `tfsdk:"product_id" query:"productId,optional"`
}

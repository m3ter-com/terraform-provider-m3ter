// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package aggregation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type AggregationsDataListDataSourceEnvelope struct {
	Data customfield.NormalizedDynamicValue `json:"data,computed"`
}

type AggregationsDataSourceModel struct {
	OrgID     types.String                       `tfsdk:"org_id" path:"orgId,optional"`
	Codes     *[]types.String                    `tfsdk:"codes" query:"codes,optional"`
	IDs       *[]types.String                    `tfsdk:"ids" query:"ids,optional"`
	ProductID *[]types.String                    `tfsdk:"product_id" query:"productId,optional"`
	MaxItems  types.Int64                        `tfsdk:"max_items"`
	Items     customfield.NormalizedDynamicValue `tfsdk:"items"`
}

func (m *AggregationsDataSourceModel) toListParams(_ context.Context) (params m3ter.AggregationListParams, diags diag.Diagnostics) {
	mCodes := []string{}
	if m.Codes != nil {
		for _, item := range *m.Codes {
			mCodes = append(mCodes, item.ValueString())
		}
	}
	mIDs := []string{}
	if m.IDs != nil {
		for _, item := range *m.IDs {
			mIDs = append(mIDs, item.ValueString())
		}
	}
	mProductID := []string{}
	if m.ProductID != nil {
		for _, item := range *m.ProductID {
			mProductID = append(mProductID, item.ValueString())
		}
	}

	params = m3ter.AggregationListParams{
		Codes:     m3ter.F(mCodes),
		IDs:       m3ter.F(mIDs),
		ProductID: m3ter.F(mProductID),
	}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

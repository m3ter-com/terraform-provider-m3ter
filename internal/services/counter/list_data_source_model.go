// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type CountersDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[CountersItemsDataSourceModel] `json:"data,computed"`
}

type CountersDataSourceModel struct {
	OrgID     types.String                                               `tfsdk:"org_id" path:"orgId,optional"`
	Codes     *[]types.String                                            `tfsdk:"codes" query:"codes,optional"`
	IDs       *[]types.String                                            `tfsdk:"ids" query:"ids,optional"`
	ProductID *[]types.String                                            `tfsdk:"product_id" query:"productId,optional"`
	MaxItems  types.Int64                                                `tfsdk:"max_items"`
	Items     customfield.NestedObjectList[CountersItemsDataSourceModel] `tfsdk:"items"`
}

func (m *CountersDataSourceModel) toListParams(_ context.Context) (params m3ter.CounterListParams, diags diag.Diagnostics) {
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

	params = m3ter.CounterListParams{
		Codes:     m3ter.F(mCodes),
		IDs:       m3ter.F(mIDs),
		ProductID: m3ter.F(mProductID),
	}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type CountersItemsDataSourceModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	Code      types.String `tfsdk:"code" json:"code,computed"`
	Name      types.String `tfsdk:"name" json:"name,computed"`
	ProductID types.String `tfsdk:"product_id" json:"productId,computed"`
	Unit      types.String `tfsdk:"unit" json:"unit,computed"`
	Version   types.Int64  `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

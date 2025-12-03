// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
)

type CounterDataSourceModel struct {
	ID        types.String                     `tfsdk:"id" path:"id,computed_optional"`
	OrgID     types.String                     `tfsdk:"org_id" path:"orgId,optional"`
	Code      types.String                     `tfsdk:"code" json:"code,computed"`
	Name      types.String                     `tfsdk:"name" json:"name,computed"`
	ProductID types.String                     `tfsdk:"product_id" json:"productId,computed"`
	Unit      types.String                     `tfsdk:"unit" json:"unit,computed"`
	Version   types.Int64                      `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	FindOneBy *CounterFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *CounterDataSourceModel) toReadParams(_ context.Context) (params m3ter.CounterGetParams, diags diag.Diagnostics) {
	params = m3ter.CounterGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

func (m *CounterDataSourceModel) toListParams(_ context.Context) (params m3ter.CounterListParams, diags diag.Diagnostics) {
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

	params = m3ter.CounterListParams{
		Codes:     m3ter.F(mFindOneByCodes),
		IDs:       m3ter.F(mFindOneByIDs),
		ProductID: m3ter.F(mFindOneByProductID),
	}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type CounterFindOneByDataSourceModel struct {
	Codes     *[]types.String `tfsdk:"codes" query:"codes,optional"`
	IDs       *[]types.String `tfsdk:"ids" query:"ids,optional"`
	ProductID *[]types.String `tfsdk:"product_id" query:"productId,optional"`
}

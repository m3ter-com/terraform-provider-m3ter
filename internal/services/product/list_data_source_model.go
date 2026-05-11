// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type ProductsDataListDataSourceEnvelope struct {
	Data customfield.NormalizedDynamicValue `json:"data,computed"`
}

type ProductsDataSourceModel struct {
	OrgID    types.String                       `tfsdk:"org_id" path:"orgId,optional"`
	IDs      *[]types.String                    `tfsdk:"ids" query:"ids,optional"`
	MaxItems types.Int64                        `tfsdk:"max_items"`
	Items    customfield.NormalizedDynamicValue `tfsdk:"items"`
}

func (m *ProductsDataSourceModel) toListParams(_ context.Context) (params m3ter.ProductListParams, diags diag.Diagnostics) {
	mIDs := []string{}
	if m.IDs != nil {
		for _, item := range *m.IDs {
			mIDs = append(mIDs, item.ValueString())
		}
	}

	params = m3ter.ProductListParams{
		IDs: m3ter.F(mIDs),
	}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

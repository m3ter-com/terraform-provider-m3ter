// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type ProductDataSourceModel struct {
	ID           types.String                       `tfsdk:"id" path:"id,computed_optional"`
	OrgID        types.String                       `tfsdk:"org_id" path:"orgId,optional"`
	Code         types.String                       `tfsdk:"code" json:"code,computed"`
	Name         types.String                       `tfsdk:"name" json:"name,computed"`
	Version      types.Int64                        `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	CustomFields customfield.NormalizedDynamicValue `tfsdk:"custom_fields" json:"customFields,computed"`
	FindOneBy    *ProductFindOneByDataSourceModel   `tfsdk:"find_one_by"`
}

func (m *ProductDataSourceModel) toReadParams(_ context.Context) (params m3ter.ProductGetParams, diags diag.Diagnostics) {
	params = m3ter.ProductGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

func (m *ProductDataSourceModel) toListParams(_ context.Context) (params m3ter.ProductListParams, diags diag.Diagnostics) {
	mFindOneByIDs := []string{}
	if m.FindOneBy.IDs != nil {
		for _, item := range *m.FindOneBy.IDs {
			mFindOneByIDs = append(mFindOneByIDs, item.ValueString())
		}
	}

	params = m3ter.ProductListParams{
		IDs: m3ter.F(mFindOneByIDs),
	}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type ProductFindOneByDataSourceModel struct {
	IDs *[]types.String `tfsdk:"ids" query:"ids,optional"`
}

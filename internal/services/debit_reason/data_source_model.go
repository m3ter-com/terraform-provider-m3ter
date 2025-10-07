// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package debit_reason

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
)

type DebitReasonDataSourceModel struct {
	ID        types.String                         `tfsdk:"id" path:"id,computed_optional"`
	OrgID     types.String                         `tfsdk:"org_id" path:"orgId,required"`
	Archived  types.Bool                           `tfsdk:"archived" json:"archived,computed"`
	Code      types.String                         `tfsdk:"code" json:"code,computed"`
	Name      types.String                         `tfsdk:"name" json:"name,computed"`
	Version   types.Int64                          `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	FindOneBy *DebitReasonFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *DebitReasonDataSourceModel) toReadParams(_ context.Context) (params m3ter.DebitReasonGetParams, diags diag.Diagnostics) {
	params = m3ter.DebitReasonGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

func (m *DebitReasonDataSourceModel) toListParams(_ context.Context) (params m3ter.DebitReasonListParams, diags diag.Diagnostics) {
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

	params = m3ter.DebitReasonListParams{
		Codes: m3ter.F(mFindOneByCodes),
		IDs:   m3ter.F(mFindOneByIDs),
	}

	if !m.FindOneBy.Archived.IsNull() {
		params.Archived = m3ter.F(m.FindOneBy.Archived.ValueBool())
	}

	return
}

type DebitReasonFindOneByDataSourceModel struct {
	Archived types.Bool      `tfsdk:"archived" query:"archived,optional"`
	Codes    *[]types.String `tfsdk:"codes" query:"codes,optional"`
	IDs      *[]types.String `tfsdk:"ids" query:"ids,optional"`
}

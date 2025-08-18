// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package debit_reason

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type DebitReasonsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[DebitReasonsItemsDataSourceModel] `json:"data,computed"`
}

type DebitReasonsDataSourceModel struct {
	OrgID    types.String                                                   `tfsdk:"org_id" path:"orgId,optional"`
	Archived types.Bool                                                     `tfsdk:"archived" query:"archived,optional"`
	Codes    *[]types.String                                                `tfsdk:"codes" query:"codes,optional"`
	IDs      *[]types.String                                                `tfsdk:"ids" query:"ids,optional"`
	MaxItems types.Int64                                                    `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[DebitReasonsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *DebitReasonsDataSourceModel) toListParams(_ context.Context) (params m3ter.DebitReasonListParams, diags diag.Diagnostics) {
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

	params = m3ter.DebitReasonListParams{
		Codes: m3ter.F(mCodes),
		IDs:   m3ter.F(mIDs),
	}

	if !m.Archived.IsNull() {
		params.Archived = m3ter.F(m.Archived.ValueBool())
	}
	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type DebitReasonsItemsDataSourceModel struct {
	ID       types.String `tfsdk:"id" json:"id,computed"`
	Archived types.Bool   `tfsdk:"archived" json:"archived,computed"`
	Code     types.String `tfsdk:"code" json:"code,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	Version  types.Int64  `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

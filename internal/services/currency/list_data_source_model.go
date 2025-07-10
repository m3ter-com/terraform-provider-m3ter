// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package currency

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type CurrenciesDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[CurrenciesItemsDataSourceModel] `json:"data,computed"`
}

type CurrenciesDataSourceModel struct {
	OrgID    types.String                                                 `tfsdk:"org_id" path:"orgId,optional"`
	Archived types.Bool                                                   `tfsdk:"archived" query:"archived,optional"`
	Codes    *[]types.String                                              `tfsdk:"codes" query:"codes,optional"`
	IDs      *[]types.String                                              `tfsdk:"ids" query:"ids,optional"`
	MaxItems types.Int64                                                  `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[CurrenciesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *CurrenciesDataSourceModel) toListParams(_ context.Context) (params m3ter.CurrencyListParams, diags diag.Diagnostics) {
	mCodes := []string{}
	for _, item := range *m.Codes {
		mCodes = append(mCodes, item.ValueString())
	}
	mIDs := []string{}
	for _, item := range *m.IDs {
		mIDs = append(mIDs, item.ValueString())
	}

	params = m3ter.CurrencyListParams{
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

type CurrenciesItemsDataSourceModel struct {
	ID               types.String      `tfsdk:"id" json:"id,computed"`
	Archived         types.Bool        `tfsdk:"archived" json:"archived,computed"`
	Code             types.String      `tfsdk:"code" json:"code,computed"`
	CreatedBy        types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated        timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified   timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy   types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	MaxDecimalPlaces types.Int64       `tfsdk:"max_decimal_places" json:"maxDecimalPlaces,computed"`
	Name             types.String      `tfsdk:"name" json:"name,computed"`
	RoundingMode     types.String      `tfsdk:"rounding_mode" json:"roundingMode,computed"`
	Version          types.Int64       `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

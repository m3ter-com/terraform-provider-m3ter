// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package currency

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type CurrencyModel struct {
	ID               types.String `tfsdk:"id" json:"id,computed"`
	OrgID            types.String `tfsdk:"org_id" path:"orgId,optional"`
	Name             types.String `tfsdk:"name" json:"name,required"`
	Archived         types.Bool   `tfsdk:"archived" json:"archived,optional"`
	Code             types.String `tfsdk:"code" json:"code,optional"`
	MaxDecimalPlaces types.Int64  `tfsdk:"max_decimal_places" json:"maxDecimalPlaces,optional"`
	RoundingMode     types.String `tfsdk:"rounding_mode" json:"roundingMode,optional"`
	Version          types.Int64  `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m CurrencyModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CurrencyModel) MarshalJSONForUpdate(state CurrencyModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

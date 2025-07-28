// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type CounterModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	OrgID     types.String `tfsdk:"org_id" path:"orgId,optional"`
	Name      types.String `tfsdk:"name" json:"name,required"`
	Unit      types.String `tfsdk:"unit" json:"unit,required"`
	Code      types.String `tfsdk:"code" json:"code,optional"`
	ProductID types.String `tfsdk:"product_id" json:"productId,optional"`
	Version   types.Int64  `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m CounterModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CounterModel) MarshalJSONForUpdate(state CounterModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
